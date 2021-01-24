package netsync

import (
	"github.com/l0k18/pod/cmd/node/mempool"
	blockchain "github.com/l0k18/pod/pkg/chain"
	"github.com/l0k18/pod/pkg/chain/config/netparams"
	chainhash "github.com/l0k18/pod/pkg/chain/hash"
	"github.com/l0k18/pod/pkg/chain/wire"
	"github.com/l0k18/pod/pkg/comm/peer"
	"github.com/l0k18/pod/pkg/util"
)

// PeerNotifier exposes methods to notify peers of status changes to transactions, blocks, etc. Currently server (in the
// main package) implements this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)
	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)
	RelayInventory(invVect *wire.InvVect, data interface{})
	TransactionConfirmed(tx *util.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier       PeerNotifier
	Chain              *blockchain.BlockChain
	TxMemPool          *mempool.TxPool
	ChainParams        *netparams.Params
	DisableCheckpoints bool
	MaxPeers           int
	FeeEstimator       *mempool.FeeEstimator
}
