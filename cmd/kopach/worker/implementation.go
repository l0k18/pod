package worker

import (
	"crypto/cipher"
	"errors"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
	
	qu "github.com/l0k18/pod/pkg/util/quit"
	
	"github.com/l0k18/pod/cmd/kopach/control/hashrate"
	"github.com/l0k18/pod/cmd/kopach/control/sol"
	blockchain "github.com/l0k18/pod/pkg/chain"
	"github.com/l0k18/pod/pkg/chain/fork"
	
	"github.com/VividCortex/ewma"
	"go.uber.org/atomic"
	
	"github.com/l0k18/pod/cmd/kopach/control"
	"github.com/l0k18/pod/cmd/kopach/control/job"
	chainhash "github.com/l0k18/pod/pkg/chain/hash"
	"github.com/l0k18/pod/pkg/chain/wire"
	"github.com/l0k18/pod/pkg/comm/stdconn"
	"github.com/l0k18/pod/pkg/comm/transport"
	"github.com/l0k18/pod/pkg/data/ring"
	"github.com/l0k18/pod/pkg/util"
	"github.com/l0k18/pod/pkg/util/interrupt"
)

const RoundsPerAlgo = 69

type Worker struct {
	mx            sync.Mutex
	id            string
	pipeConn      *stdconn.StdConn
	dispatchConn  *transport.Channel
	dispatchReady atomic.Bool
	ciph          cipher.AEAD
	quit          qu.C
	block         atomic.Value
	senderPort    atomic.Uint32
	msgBlock      atomic.Value // *wire.MsgBlock
	bitses        atomic.Value
	hashes        atomic.Value
	lastMerkle    *chainhash.Hash
	roller        *Counter
	startNonce    uint32
	startChan     qu.C
	stopChan      qu.C
	running       atomic.Bool
	hashCount     atomic.Uint64
	hashSampleBuf *ring.BufferUint64
}

type Counter struct {
	rpa           int32
	C             atomic.Int32
	Algos         atomic.Value // []int32
	RoundsPerAlgo atomic.Int32
}

// NewCounter returns an initialized algorithm rolling counter that ensures each miner does equal amounts of every
// algorithm
func NewCounter(roundsPerAlgo int32) (c *Counter) {
	// these will be populated when work arrives
	var algos []int32
	// Start the counter at a random position
	rand.Seed(time.Now().UnixNano())
	c = &Counter{}
	c.C.Store(int32(rand.Intn(int(roundsPerAlgo)+1) + 1))
	c.Algos.Store(algos)
	c.RoundsPerAlgo.Store(roundsPerAlgo)
	c.rpa = roundsPerAlgo
	return
}

// GetAlgoVer returns the next algo version based on the current configuration
func (c *Counter) GetAlgoVer() (ver int32) {
	// the formula below rolls through versions with blocks roundsPerAlgo long for each algorithm by its index
	algs := c.Algos.Load().([]int32)
	// Debug(algs)
	if c.RoundsPerAlgo.Load() < 1 {
		Debug("RoundsPerAlgo is", c.RoundsPerAlgo.Load(), len(algs))
		return 0
	}
	if len(algs) > 0 {
		ver = algs[(c.C.Load()/
			c.RoundsPerAlgo.Load())%
			int32(len(algs))]
		c.C.Add(1)
	}
	return
}

func (w *Worker) hashReport() {
	w.hashSampleBuf.Add(w.hashCount.Load())
	av := ewma.NewMovingAverage(15)
	var i int
	var prev uint64
	if err := w.hashSampleBuf.ForEach(
		func(v uint64) error {
			if i < 1 {
				prev = v
			} else {
				interval := v - prev
				av.Add(float64(interval))
				prev = v
			}
			i++
			return nil
		},
	); Check(err) {
	}
	// Info("kopach",w.hashSampleBuf.Cursor, w.hashSampleBuf.Buf)
	Tracef("average hashrate %.2f", av.Value())
}

// NewWithConnAndSemaphore is exposed to enable use an actual network connection while retaining the same RPC API to
// allow a worker to be configured to run on a bare metal system with a different launcher main
func NewWithConnAndSemaphore(id string, conn *stdconn.StdConn, quit qu.C) *Worker {
	Debug("creating new worker")
	msgBlock := wire.MsgBlock{Header: wire.BlockHeader{}}
	w := &Worker{
		id:            id,
		pipeConn:      conn,
		quit:          quit,
		roller:        NewCounter(RoundsPerAlgo),
		startChan:     qu.T(),
		stopChan:      qu.T(),
		hashSampleBuf: ring.NewBufferUint64(1000),
	}
	w.msgBlock.Store(msgBlock)
	w.block.Store(util.NewBlock(&msgBlock))
	w.dispatchReady.Store(false)
	// with this we can report cumulative hash counts as well as using it to distribute algorithms evenly
	w.startNonce = uint32(w.roller.C.Load())
	interrupt.AddHandler(
		func() {
			Debug("worker quitting")
			w.stopChan <- struct{}{}
			// _ = w.pipeConn.Close()
			w.dispatchReady.Store(false)
		},
	)
	go worker(w)
	return w
}

func worker(w *Worker) {
	Debug("main work loop starting")
	sampleTicker := time.NewTicker(time.Second)
out:
	for {
		// Pause state
		Trace("worker pausing")
	pausing:
		for {
			select {
			case <-sampleTicker.C:
				w.hashReport()
				break
			case <-w.stopChan:
				Trace("received pause signal while paused")
				// drain stop channel in pause
				break
			case <-w.startChan:
				Trace("received start signal")
				break pausing
			case <-w.quit:
				Trace("quitting")
				break out
			}
		}
		// Run state
		Trace("worker running")
	running:
		for {
			select {
			case <-sampleTicker.C:
				// Debug(interrupt.GoroutineDump())
				w.hashReport()
				break
			case <-w.startChan:
				Trace("received start signal while running")
				// drain start channel in run mode
				break
			case <-w.stopChan:
				Trace("received pause signal while running")
				break running
			case <-w.quit:
				Trace("worker stopping while running")
				break out
			default:
				if w.block.Load() == nil || w.bitses.Load() == nil ||
					w.hashes.Load() == nil || !w.dispatchReady.Load() {
					// Debug("not ready to work")
				} else {
					// Debug("working")
					// work
					nH := w.block.Load().(*util.Block).Height()
					hv := w.roller.GetAlgoVer()
					h := w.hashes.Load().(map[int32]*chainhash.Hash)
					mmb := w.msgBlock.Load().(wire.MsgBlock)
					mb := &mmb
					mb.Header.Version = hv
					if h != nil {
						mr, ok := h[hv]
						if !ok {
							continue
						}
						mb.Header.MerkleRoot = *mr
					} else {
						continue
					}
					b := w.bitses.Load().(blockchain.TargetBits)
					if bb, ok := b[mb.Header.Version]; ok {
						mb.Header.Bits = bb
					} else {
						continue
					}
					mb.Header.Timestamp = time.Now()
					var nextAlgo int32
					if w.roller.C.Load()%w.roller.RoundsPerAlgo.Load() == 0 {
						select {
						case <-w.quit:
							Trace("worker stopping on pausing message")
							break out
						default:
						}
						// send out broadcast containing worker nonce and algorithm and count of blocks
						w.hashCount.Store(w.hashCount.Load() + uint64(w.roller.RoundsPerAlgo.Load()))
						nextAlgo = w.roller.C.Load() + 1
						hashReport := hashrate.Get(w.roller.RoundsPerAlgo.Load(), nextAlgo, nH, w.id)
						err := w.dispatchConn.SendMany(
							hashrate.Magic,
							transport.GetShards(hashReport.Data),
						)
						if err != nil {
							Error(err)
						}
					}
					hash := mb.Header.BlockHashWithAlgos(nH)
					bigHash := blockchain.HashToBig(&hash)
					if bigHash.Cmp(fork.CompactToBig(mb.Header.Bits)) <= 0 {
						Debug("found solution")
						srs := sol.GetSolContainer(w.senderPort.Load(), mb)
						err := w.dispatchConn.SendMany(
							sol.SolutionMagic,
							transport.GetShards(srs.Data),
						)
						if err != nil {
							Error(err)
						}
						Debug("sent solution")
						break running
					}
					mb.Header.Version = nextAlgo
					mb.Header.Bits = w.bitses.Load().(blockchain.TargetBits)[mb.Header.Version]
					mb.Header.Nonce++
					w.msgBlock.Store(*mb)
				}
			}
		}
	}
	Debug("worker finished")
	interrupt.Request()
}

// New initialises the state for a worker, loading the work function handler that runs a round of processing between
// checking quit signal and work semaphore
func New(id string, quit qu.C) (w *Worker, conn net.Conn) {
	// log.L.SetLevel("trace", true)
	sc := stdconn.New(os.Stdin, os.Stdout, quit)
	
	return NewWithConnAndSemaphore(id, sc, quit), sc
}

// NewJob is a delivery of a new job for the worker, this makes the miner start mining from pause or pause, prepare the
// work and restart
func (w *Worker) NewJob(job *job.Container, reply *bool) (err error) {
	Debug("received new job")
	if !w.dispatchReady.Load() { // || !w.running.Load() {
		Debug("dispatch not ready")
		*reply = true
		return
	}
	j := job.Struct()
	w.bitses.Store(j.Bitses)
	w.hashes.Store(j.Hashes)
	if j.Hashes[5].IsEqual(w.lastMerkle) {
		Debug("not a new job")
		*reply = true
		return
	}
	var algos []int32
	for i := range j.Bitses {
		// we don't need to know net params if version numbers come with jobs
		algos = append(algos, i)
	}
	w.lastMerkle = j.Hashes[5]
	*reply = true
	// halting current work
	Debug("halting current work")
	w.stopChan <- struct{}{}
	newHeight := job.GetNewHeight()
	
	if len(algos) > 0 {
		// if we didn't get them in the job don't update the old
		w.roller.Algos.Store(algos)
	}
	mbb := w.msgBlock.Load().(wire.MsgBlock)
	mb := &mbb
	mb.Header.PrevBlock = *job.GetPrevBlockHash()
	// TODO: ensure worker time sync - ntp? time wrapper with skew adjustment
	hv := w.roller.GetAlgoVer()
	mb.Header.Version = hv
	var ok bool
	mb.Header.Bits, ok = j.Bitses[mb.Header.Version]
	if !ok {
		return errors.New("bits are empty")
	}
	rand.Seed(time.Now().UnixNano())
	mb.Header.Nonce = rand.Uint32()
	if j.Hashes == nil {
		return errors.New("failed to decode merkle roots")
	} else {
		hh, ok := j.Hashes[hv]
		if !ok {
			return errors.New("could not get merkle root from job")
		}
		mb.Header.MerkleRoot = *hh
	}
	mb.Header.Timestamp = time.Now()
	// make the work select block start running
	bb := util.NewBlock(mb)
	bb.SetHeight(newHeight)
	w.block.Store(bb)
	w.msgBlock.Store(*mb)
	w.senderPort.Store(uint32(job.GetControllerListenerPort()))
	// halting current work
	Debug("switching to new job")
	// w.stopChan <- struct{}{}
	w.startChan <- struct{}{}
	return
}

// Pause signals the worker to stop working, releases its semaphore and the worker is then idle
func (w *Worker) Pause(_ int, reply *bool) (err error) {
	Trace("pausing from IPC")
	w.running.Store(false)
	w.stopChan <- struct{}{}
	*reply = true
	return
}

// Stop signals the worker to quit
func (w *Worker) Stop(_ int, reply *bool) (err error) {
	Debug("stopping from IPC")
	w.stopChan <- struct{}{}
	defer w.quit.Q()
	*reply = true
	// time.Sleep(time.Second * 3)
	// os.Exit(0)
	return
}

// SendPass gives the encryption key configured in the kopach controller ( pod) configuration to allow workers to
// dispatch their solutions
func (w *Worker) SendPass(pass string, reply *bool) (err error) {
	Debug("receiving dispatch password", pass)
	rand.Seed(time.Now().UnixNano())
	// sp := fmt.Sprint(rand.Intn(32767) + 1025)
	// rp := fmt.Sprint(rand.Intn(32767) + 1025)
	var conn *transport.Channel
	conn, err = transport.NewBroadcastChannel(
		"kopachworker",
		w,
		pass,
		transport.DefaultPort,
		control.MaxDatagramSize,
		transport.Handlers{},
		w.quit,
	)
	if err != nil {
		Error(err)
	}
	w.dispatchConn = conn
	w.dispatchReady.Store(true)
	*reply = true
	return
}
