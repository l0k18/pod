package config

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	
	"github.com/l0k18/pod/app/apputil"
	"github.com/l0k18/pod/cmd/node"
	blockchain "github.com/l0k18/pod/pkg/chain"
	"github.com/l0k18/pod/pkg/chain/forkhash"
	"github.com/l0k18/pod/pkg/comm/peer/connmgr"
	"github.com/l0k18/pod/pkg/util"
	"github.com/l0k18/pod/pkg/util/interrupt"
	"github.com/l0k18/pod/pkg/util/normalize"
	"github.com/l0k18/pod/pkg/util/routeable"
	"github.com/l0k18/pod/pkg/wallet"
	
	"github.com/btcsuite/go-socks/socks"
	"github.com/urfave/cli"
	
	"github.com/l0k18/pod/app/appdata"
	"github.com/l0k18/pod/app/conte"
	"github.com/l0k18/pod/cmd/node/state"
	"github.com/l0k18/pod/pkg/chain/config/netparams"
	"github.com/l0k18/pod/pkg/chain/fork"
	"github.com/l0k18/pod/pkg/pod"
	"github.com/l0k18/pod/pkg/util/logi"
)

const (
	appName           = "pod"
	confExt           = ".json"
	appLanguage       = "en"
	podConfigFilename = appName + confExt
	PARSER            = "json"
)

var funcName = "loadConfig"

func initDictionary(cfg *pod.Config) {
	if cfg.Language == nil || *cfg.Language == "" {
		*cfg.Language = Lang("en")
	}
	Trace("lang set to", *cfg.Language)
}

func initDataDir(cfg *pod.Config) {
	if cfg.DataDir == nil || *cfg.DataDir == "" {
		Debug("setting default data dir")
		*cfg.DataDir = appdata.Dir("pod", false)
	}
	Trace("datadir set to", *cfg.DataDir)
}

func initWalletFile(cx *conte.Xt) {
	if cx.Config.WalletFile == nil || *cx.Config.WalletFile == "" {
		*cx.Config.WalletFile = *cx.Config.DataDir + string(os.PathSeparator) +
			cx.ActiveNet.Name + string(os.PathSeparator) + wallet.DbName
	}
	Trace("wallet file set to", *cx.Config.WalletFile, *cx.Config.Network)
}

func initConfigFile(cfg *pod.Config) {
	if *cfg.ConfigFile == "" {
		*cfg.ConfigFile =
			*cfg.DataDir + string(os.PathSeparator) + podConfigFilename
	}
	Trace("using config file:", *cfg.ConfigFile)
}

func initLogDir(cfg *pod.Config) {
	if *cfg.LogDir != "" {
		logi.L.SetLogPaths(*cfg.LogDir, "pod")
		interrupt.AddHandler(func() {
			Debug("initLogDir interrupt")
			_ = logi.L.LogFileHandle.Close()
		})
	}
}

func initParams(cx *conte.Xt) {
	network := "mainnet"
	if cx.Config.Network != nil {
		network = *cx.Config.Network
	}
	switch network {
	case "testnet", "testnet3", "t":
		Trace("on testnet")
		cx.ActiveNet = &netparams.TestNet3Params
		fork.IsTestnet = true
	case "regtestnet", "regressiontest", "r":
		Trace("on regression testnet")
		cx.ActiveNet = &netparams.RegressionTestParams
	case "simnet", "s":
		Trace("on simnet")
		cx.ActiveNet = &netparams.SimNetParams
	default:
		if network != "mainnet" && network != "m" {
			Warn("using mainnet for node")
		}
		Trace("on mainnet")
		cx.ActiveNet = &netparams.MainNetParams
	}
}

func validatePort(port string) bool {
	var err error
	var p int64
	if p, err = strconv.ParseInt(port, 10, 32); Check(err) {
		return false
	}
	if p < 1024 || p > 65535 {
		return false
	}
	return true
}

func initListeners(cx *conte.Xt, commandName string, initial bool) {
	cfg := cx.Config
	var e error
	var fP int
	if fP, e = GetFreePort(); Check(e) {
	}
	// var routeableAddresses []net.Addr
	_, addresses := routeable.GetInterface()
	// routeableAddresses, e = routeableInterface[0].Addrs()
	// Debug(routeableAddresses)
	routeableAddress := addresses[0]
	// Debug("********************", routeableAddress)
	*cfg.Controller = net.JoinHostPort(routeableAddress, fmt.Sprint(fP))
	if len(*cfg.Listeners) < 1 && !*cfg.DisableListen && len(*cfg.ConnectPeers) < 1 {
		cfg.Listeners = &cli.StringSlice{fmt.Sprintf(routeableAddress + ":" + cx.ActiveNet.DefaultPort)}
		cx.StateCfg.Save = true
		Debug("Listeners")
	}
	if len(*cfg.RPCListeners) < 1 {
		address := fmt.Sprintf("%s:%s", routeableAddress, cx.ActiveNet.RPCClientPort)
		*cfg.RPCListeners = cli.StringSlice{address}
		*cfg.RPCConnect = address
		Debug("setting save flag because rpc listeners is empty and rpc is not disabled")
		cx.StateCfg.Save = true
		Debug("RPCListeners")
	}
	if len(*cfg.WalletRPCListeners) < 1 && !*cfg.DisableRPC {
		address := fmt.Sprintf(routeableAddress + ":" + cx.ActiveNet.WalletRPCServerPort)
		*cfg.WalletRPCListeners = cli.StringSlice{address}
		*cfg.WalletServer = address
		Debug("setting save flag because wallet rpc listeners is empty and" +
			" rpc is not disabled")
		cx.StateCfg.Save = true
		Debug("WalletRPCListeners")
	}
	// msgBase := pause.GetPauseContainer(cx)
	// mC := job.Get(cx, util.NewBlock(tpl.Block), msgBase)
	// listenHost := msgBase.GetIPs()[0].String() + ":0"
	// switch commandName {
	// // only the wallet listener is important with shell as it proxies for
	// // node, the rest better they are automatic
	// case "shell":
	// 	*cfg.Listeners = cli.StringSlice{listenHost}
	// 	*cfg.RPCListeners = cli.StringSlice{listenHost}
	// }
	if *cx.Config.AutoPorts || !initial {
		if fP, e = GetFreePort(); Check(e) {
		}
		*cfg.Listeners = cli.StringSlice{routeableAddress + ":" + fmt.Sprint(fP)}
		if fP, e = GetFreePort(); Check(e) {
		}
		*cfg.RPCListeners = cli.StringSlice{routeableAddress + ":" + fmt.Sprint(fP)}
		if fP, e = GetFreePort(); Check(e) {
		}
		*cfg.WalletRPCListeners = cli.StringSlice{routeableAddress + ":" + fmt.Sprint(fP)}
		cx.StateCfg.Save = true
		Debug("autoports")
	} else {
		// sanitize user input and set auto on any that fail
		l := cfg.Listeners
		r := cfg.RPCListeners
		w := cfg.WalletRPCListeners
		for i := range *l {
			if _, p, e := net.SplitHostPort((*l)[i]); !Check(e) {
				if !validatePort(p) {
					if fP, e = GetFreePort(); Check(e) {
					}
					(*l)[i] = routeableAddress + ":" + fmt.Sprint(fP)
					cx.StateCfg.Save = true
					Debug("port not validate Listeners")
				}
			}
		}
		for i := range *r {
			if _, p, e := net.SplitHostPort((*r)[i]); !Check(e) {
				if !validatePort(p) {
					if fP, e = GetFreePort(); Check(e) {
					}
					(*r)[i] = routeableAddress + ":" + fmt.Sprint(fP)
					cx.StateCfg.Save = true
					Debug("port not validate RPCListeners")
				}
			}
		}
		// if *cfg.RPCConnect == "" {
		// 	*cfg.RPCConnect = routeableAddress + ":" + fmt.Sprint(fP)
		// 	Debug("setting save flag because rpcconnect was not configured")
		// 	cx.StateCfg.Save = true
		// }
		for i := range *w {
			if _, p, e := net.SplitHostPort((*w)[i]); !Check(e) {
				if !validatePort(p) {
					if fP, e = GetFreePort(); Check(e) {
					}
					(*w)[i] = routeableAddress + ":" + fmt.Sprint(fP)
					cx.StateCfg.Save = true
					Debug("port not validate WalletRPCListeners")
				}
			}
		}
	}
	// all of these can be autodiscovered/set but to do that and know what they are we have to reserve them
	//
	// listeners := []*cli.StringSlice{
	//	cfg.WalletRPCListeners,
	//	cfg.Listeners,
	//	cfg.RPCListeners,
	// }
	// for i := range listeners {
	//	if h, p, err := net.SplitHostPort((*listeners[i])[0]); p == "0" {
	//		if err != nil {
	//			Error(err)
	//		} else {
	//			fP, err := GetFreePort()
	//			if err != nil {
	//				Error(err)
	//			}
	//			*listeners[i] = cli.
	//				StringSlice{net.JoinHostPort(h, fmt.Sprint(fP))}
	//		}
	//	}
	// }
	// (*cfg.WalletRPCListeners)[0] = (*listeners[0])[0]
	// (*cfg.Listeners)[0] = (*listeners[1])[0]
	// (*cfg.RPCListeners)[0] = (*listeners[2])[0]
	//
	// if lan mode is set, remove the peers.json so no unwanted nodes are connected to
	if *cfg.LAN && cx.ActiveNet.Name != "mainnet" {
		peersFile := filepath.Join(filepath.Join(*cfg.DataDir, cx.ActiveNet.Name), "peers.json")
		var err error
		if err = os.Remove(peersFile); err != nil {
			Trace("nothing to remove?", err)
		}
		Trace("removed", peersFile)
	}
	// *cfg.RPCConnect = (*cfg.RPCListeners)[0]
	// h, p, _ := net.SplitHostPort(*cfg.RPCConnect)
	// if h == "" {
	// 	*cfg.RPCConnect = net.JoinHostPort(routeableAddress, p)
	// }
	if len(*cfg.WalletRPCListeners) > 0 {
		splitted := strings.Split((*cfg.WalletRPCListeners)[0], ":")
		*cfg.WalletServer = routeableAddress + ":" + splitted[1]
	}
	// save.Pod(cfg)
}

// GetFreePort asks the kernel for free open ports that are ready to use.
func GetFreePort() (int, error) {
	var port int
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := l.Close(); Check(err) {
		}
	}()
	port = l.Addr().(*net.TCPAddr).Port
	return port, nil
}

func initTLSStuffs(cfg *pod.Config, st *state.Config) {
	isNew := false
	if *cfg.RPCCert == "" {
		*cfg.RPCCert =
			*cfg.DataDir + string(os.PathSeparator) + "rpc.cert"
		Debug("setting save flag because rpc cert path was not set")
		st.Save = true
		isNew = true
	}
	if *cfg.RPCKey == "" {
		*cfg.RPCKey =
			*cfg.DataDir + string(os.PathSeparator) + "rpc.key"
		Debug("setting save flag because rpc key path was not set")
		st.Save = true
		isNew = true
	}
	if *cfg.CAFile == "" {
		*cfg.CAFile =
			*cfg.DataDir + string(os.PathSeparator) + "ca.cert"
		Debug("setting save flag because CA cert path was not set")
		st.Save = true
		isNew = true
	}
	if isNew {
		// Now is the best time to make the certs
		Info("generating TLS certificates")
		// Create directories for cert and key files if they do not yet exist.
		Warn("rpc tls ", *cfg.RPCCert, " ", *cfg.RPCKey)
		certDir, _ := filepath.Split(*cfg.RPCCert)
		keyDir, _ := filepath.Split(*cfg.RPCKey)
		err := os.MkdirAll(certDir, 0700)
		if err != nil {
			Error(err)
			return
		}
		err = os.MkdirAll(keyDir, 0700)
		if err != nil {
			Error(err)
			return
		}
		// Generate cert pair.
		org := "pod/wallet autogenerated cert"
		validUntil := time.Now().Add(time.Hour * 24 * 365 * 10)
		cert, key, err := util.NewTLSCertPair(org, validUntil, nil)
		if err != nil {
			Error(err)
			return
		}
		_, err = tls.X509KeyPair(cert, key)
		if err != nil {
			Error(err)
			return
		}
		// Write cert and (potentially) the key files.
		err = ioutil.WriteFile(*cfg.RPCCert, cert, 0600)
		if err != nil {
			rmErr := os.Remove(*cfg.RPCCert)
			if rmErr != nil {
				Warn("cannot remove written certificates:", rmErr)
			}
			return
		}
		err = ioutil.WriteFile(*cfg.CAFile, cert, 0600)
		if err != nil {
			rmErr := os.Remove(*cfg.RPCCert)
			if rmErr != nil {
				Warn("cannot remove written certificates:", rmErr)
			}
			return
		}
		err = ioutil.WriteFile(*cfg.RPCKey, key, 0600)
		if err != nil {
			Error(err)
			rmErr := os.Remove(*cfg.RPCCert)
			if rmErr != nil {
				Warn("cannot remove written certificates:", rmErr)
			}
			rmErr = os.Remove(*cfg.CAFile)
			if rmErr != nil {
				Warn("cannot remove written certificates:", rmErr)
			}
			return
		}
		Info("done generating TLS certificates")
		return
	}
}

func initLogLevel(cfg *pod.Config) {
	loglevel := *cfg.LogLevel
	switch loglevel {
	case "trace", "debug", "info", "warn", "error", "fatal", "off":
		Debug("log level", loglevel)
	default:
		Error("unrecognised loglevel", loglevel, "setting default info")
		*cfg.LogLevel = "info"
	}
	color := true
	if runtime.GOOS == "windows" {
		color = false
	}
	logi.L.SetLevel(*cfg.LogLevel, color, "pod")
}

func normalizeAddresses(cfg *pod.Config) {
	Trace("normalising addresses")
	port := node.DefaultPort
	nrm := normalize.StringSliceAddresses
	nrm(cfg.AddPeers, port)
	nrm(cfg.ConnectPeers, port)
	// nrm(cfg.Listeners, port)
	nrm(cfg.Whitelists, port)
	// nrm(cfg.RPCListeners, port)
}

func setRelayReject(cfg *pod.Config) {
	relayNonStd := *cfg.RelayNonStd
	switch {
	case *cfg.RelayNonStd && *cfg.RejectNonStd:
		errf := "%s: rejectnonstd and relaynonstd cannot be used together" +
			" -- choose only one, leaving neither activated"
		Error(errf, funcName)
		// just leave both false
		*cfg.RelayNonStd = false
		*cfg.RejectNonStd = false
	case *cfg.RejectNonStd:
		relayNonStd = false
	case *cfg.RelayNonStd:
		relayNonStd = true
	}
	*cfg.RelayNonStd = relayNonStd
}

func validateDBtype(cfg *pod.Config) {
	// Validate database type.
	Trace("validating database type")
	if !node.ValidDbType(*cfg.DbType) {
		str := "%s: The specified database type [%v] is invalid -- " +
			"supported types %v"
		err := fmt.Errorf(str, funcName, *cfg.DbType, node.KnownDbTypes)
		Error(funcName, err)
		// set to default
		*cfg.DbType = node.KnownDbTypes[0]
	}
}

func validateProfilePort(cfg *pod.Config) {
	// Validate profile port number
	Trace("validating profile port number")
	if *cfg.Profile != "" {
		profilePort, err := strconv.Atoi(*cfg.Profile)
		if err != nil || profilePort < 1024 || profilePort > 65535 {
			str := "%s: The profile port must be between 1024 and 65535"
			err := fmt.Errorf(str, funcName)
			Error(funcName, err)
			*cfg.Profile = ""
		}
	}
}
func validateBanDuration(cfg *pod.Config) {
	// Don't allow ban durations that are too short.
	Trace("validating ban duration")
	if *cfg.BanDuration < time.Second {
		err := fmt.Errorf("%s: The banduration option may not be less than 1s -- parsed [%v]",
			funcName, *cfg.BanDuration)
		Info(funcName, err)
		*cfg.BanDuration = node.DefaultBanDuration
	}
}

func validateWhitelists(cfg *pod.Config, st *state.Config) {
	// Validate any given whitelisted IP addresses and networks.
	Trace("validating whitelists")
	if len(*cfg.Whitelists) > 0 {
		var ip net.IP
		st.ActiveWhitelists = make([]*net.IPNet, 0, len(*cfg.Whitelists))
		for _, addr := range *cfg.Whitelists {
			_, ipnet, err := net.ParseCIDR(addr)
			if err != nil {
				Error(err)
				ip = net.ParseIP(addr)
				if ip == nil {
					str := err.Error() + " %s: The whitelist value of '%s' is invalid"
					err = fmt.Errorf(str, funcName, addr)
					Error(err)
					_, _ = fmt.Fprintln(os.Stderr, err)
					interrupt.Request()
					// os.Exit(1)
				} else {
					var bits int
					if ip.To4() == nil {
						// IPv6
						bits = 128
					} else {
						bits = 32
					}
					ipnet = &net.IPNet{
						IP:   ip,
						Mask: net.CIDRMask(bits, bits),
					}
				}
			}
			st.ActiveWhitelists = append(st.ActiveWhitelists, ipnet)
		}
	}
}

func validatePeerLists(cfg *pod.Config) {
	Trace("checking addpeer and connectpeer lists")
	if len(*cfg.AddPeers) > 0 && len(*cfg.ConnectPeers) > 0 {
		err := fmt.Errorf(
			"%s: the --addpeer and --connect options can not be mixed",
			funcName)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
}
func configListener(cfg *pod.Config, params *netparams.Params) {
	// --proxy or --connect without --listen disables listening.
	Trace("checking proxy/connect for disabling listening")
	if (*cfg.Proxy != "" ||
		len(*cfg.ConnectPeers) > 0) &&
		len(*cfg.Listeners) == 0 {
		*cfg.DisableListen = true
	}
	// Add the default listener if none were specified. The default listener is all addresses on the listen port for the
	// network we are to connect to.
	Trace("checking if listener was set")
	if len(*cfg.Listeners) == 0 {
		*cfg.Listeners = []string{":" + params.DefaultPort}
	}
}

func validateUsers(cfg *pod.Config) {
	// Check to make sure limited and admin users don't have the same username
	Trace("checking admin and limited username is different")
	if *cfg.Username != "" &&
		*cfg.Username == *cfg.LimitUser {
		str := "%s: --username and --limituser must not specify the same username"
		err := fmt.Errorf(str, funcName)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	// Check to make sure limited and admin users don't have the same password
	Trace("checking limited and admin passwords are not the same")
	if *cfg.Password != "" &&
		*cfg.Password == *cfg.LimitPass {
		str := "%s: --password and --limitpass must not specify the same password"
		err := fmt.Errorf(str, funcName)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
}

func configRPC(cfg *pod.Config, params *netparams.Params) {
	// The RPC server is disabled if no username or password is provided.
	Trace("checking rpc server has a login enabled")
	if (*cfg.Username == "" || *cfg.Password == "") &&
		(*cfg.LimitUser == "" || *cfg.LimitPass == "") {
		*cfg.DisableRPC = true
	}
	if *cfg.DisableRPC {
		Trace("RPC service is disabled")
	}
	Trace("checking rpc server has listeners set")
	if !*cfg.DisableRPC && len(*cfg.RPCListeners) == 0 {
		Debug("looking up default listener")
		addrs, err := net.LookupHost(node.DefaultRPCListener)
		if err != nil {
			Error(err)
			// os.Exit(1)
		}
		*cfg.RPCListeners = make([]string, 0, len(addrs))
		Debug("setting listeners")
		for _, addr := range addrs {
			*cfg.RPCListeners = append(*cfg.RPCListeners, addr)
			addr = net.JoinHostPort(addr, params.RPCClientPort)
		}
	}
	Trace("checking rpc max concurrent requests")
	if *cfg.RPCMaxConcurrentReqs < 0 {
		str := "%s: The rpcmaxwebsocketconcurrentrequests option may not be" +
			" less than 0 -- parsed [%d]"
		err := fmt.Errorf(str, funcName, *cfg.RPCMaxConcurrentReqs)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	Trace("checking rpc listener addresses")
	nrms := normalize.Addresses
	// Add default port to all rpc listener addresses if needed and remove duplicate addresses.
	//
	// *cfg.RPCListeners = nrms(*cfg.RPCListeners, cx.ActiveNet.RPCClientPort)
	//
	// Add default port to all listener addresses if needed and remove duplicate addresses.
	//
	// *cfg.Listeners = nrms(*cfg.Listeners, cx.ActiveNet.DefaultPort)
	//
	// Add default port to all added peer addresses if needed and remove duplicate addresses.
	*cfg.AddPeers = nrms(*cfg.AddPeers, params.DefaultPort)
	*cfg.ConnectPeers = nrms(*cfg.ConnectPeers, params.DefaultPort)
}

func validatePolicies(cfg *pod.Config, stateConfig *state.Config) {
	var err error
	
	// Validate the the minrelaytxfee.
	Trace("checking min relay tx fee")
	stateConfig.ActiveMinRelayTxFee, err = util.NewAmount(*cfg.MinRelayTxFee)
	if err != nil {
		Error(err)
		str := "%s: invalid minrelaytxfee: %v"
		err := fmt.Errorf(str, funcName, err)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	// Limit the max block size to a sane value.
	Trace("checking max block size")
	if *cfg.BlockMaxSize < node.BlockMaxSizeMin ||
		*cfg.BlockMaxSize > node.BlockMaxSizeMax {
		str := "%s: The blockmaxsize option must be in between %d and %d -- parsed [%d]"
		err := fmt.Errorf(str, funcName, node.BlockMaxSizeMin,
			node.BlockMaxSizeMax, *cfg.BlockMaxSize)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	// Limit the max block weight to a sane value.
	Trace("checking max block weight")
	if *cfg.BlockMaxWeight < node.BlockMaxWeightMin ||
		*cfg.BlockMaxWeight > node.BlockMaxWeightMax {
		str := "%s: The blockmaxweight option must be in between %d and %d -- parsed [%d]"
		err := fmt.Errorf(str, funcName, node.BlockMaxWeightMin,
			node.BlockMaxWeightMax, *cfg.BlockMaxWeight)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	// Limit the max orphan count to a sane vlue.
	Trace("checking max orphan limit")
	if *cfg.MaxOrphanTxs < 0 {
		str := "%s: The maxorphantx option may not be less than 0 -- parsed [%d]"
		err := fmt.Errorf(str, funcName, *cfg.MaxOrphanTxs)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	// Limit the block priority and minimum block sizes to max block size.
	Trace("validating block priority and minimum size/weight")
	*cfg.BlockPrioritySize = int(apputil.MinUint32(
		uint32(*cfg.BlockPrioritySize),
		uint32(*cfg.BlockMaxSize)))
	*cfg.BlockMinSize = int(apputil.MinUint32(
		uint32(*cfg.BlockMinSize),
		uint32(*cfg.BlockMaxSize)))
	*cfg.BlockMinWeight = int(apputil.MinUint32(
		uint32(*cfg.BlockMinWeight),
		uint32(*cfg.BlockMaxWeight)))
	switch {
	// If the max block size isn't set, but the max weight is, then we'll set the limit for the max block size to a safe
	// limit so weight takes precedence.
	case *cfg.BlockMaxSize == node.DefaultBlockMaxSize &&
		*cfg.BlockMaxWeight != node.DefaultBlockMaxWeight:
		*cfg.BlockMaxSize = blockchain.MaxBlockBaseSize - 1000
		// If the max block weight isn't set, but the block size is, then we'll scale the set weight accordingly based
		// on the max block size value.
	case *cfg.BlockMaxSize != node.DefaultBlockMaxSize &&
		*cfg.BlockMaxWeight == node.DefaultBlockMaxWeight:
		*cfg.BlockMaxWeight = *cfg.BlockMaxSize * blockchain.WitnessScaleFactor
	}
	// Look for illegal characters in the user agent comments.
	Trace("checking user agent comments", cfg.UserAgentComments)
	for _, uaComment := range *cfg.UserAgentComments {
		if strings.ContainsAny(uaComment, "/:()") {
			err := fmt.Errorf("%s: The following characters must not "+
				"appear in user agent comments: '/', ':', '(', ')'",
				funcName)
			_, _ = fmt.Fprintln(os.Stderr, err)
			// os.Exit(1)
		}
	}
	// Check the checkpoints for syntax errors.
	Trace("checking the checkpoints")
	stateConfig.AddedCheckpoints, err = node.ParseCheckpoints(*cfg.
		AddCheckpoints)
	if err != nil {
		Error(err)
		str := "%s: Error parsing checkpoints: %v"
		err := fmt.Errorf(str, funcName, err)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
}
func validateOnions(cfg *pod.Config) {
	// --onionproxy and not --onion are contradictory
	// TODO: this is kinda stupid hm? switch *and* toggle by presence of flag value, one should be enough
	if *cfg.Onion && *cfg.OnionProxy != "" {
		Error("onion enabled but no onionproxy has been configured")
		Fatal("halting to avoid exposing IP address")
		// os.Exit(1)
	}
	// Tor stream isolation requires either proxy or onion proxy to be set.
	if *cfg.TorIsolation &&
		*cfg.Proxy == "" &&
		*cfg.OnionProxy == "" {
		str := "%s: Tor stream isolation requires either proxy or onionproxy to be set"
		err := fmt.Errorf(str, funcName)
		_, _ = fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	if !*cfg.Onion {
		*cfg.OnionProxy = ""
	}
	
}

func validateMiningStuff(cfg *pod.Config, state *state.Config,
	params *netparams.Params) {
	// Check mining addresses are valid and saved parsed versions.
	Trace("checking mining addresses")
	state.ActiveMiningAddrs = make([]util.Address, 0, len(*cfg.MiningAddrs))
	for _, strAddr := range *cfg.MiningAddrs {
		addr, err := util.DecodeAddress(strAddr, params)
		if err != nil {
			Error(err)
			str := "%s: mining address '%s' failed to decode: %v"
			err := fmt.Errorf(str, funcName, strAddr, err)
			_, _ = fmt.Fprintln(os.Stderr, err)
			// os.Exit(1)
			continue
		}
		if !addr.IsForNet(params) {
			str := "%s: mining address '%s' is on the wrong network"
			err := fmt.Errorf(str, funcName, strAddr)
			_, _ = fmt.Fprintln(os.Stderr, err)
			// os.Exit(1)
			continue
		}
		state.ActiveMiningAddrs = append(state.ActiveMiningAddrs, addr)
	}
	// Ensure there is at least one mining address when the generate flag is set.
	if (*cfg.Generate) && len(state.ActiveMiningAddrs) == 0 {
		Error("the generate flag is set, " +
			"but there are no mining addresses specified ")
		// Traces(cfg)
		*cfg.Generate = false
		// os.Exit(1)
	}
	if *cfg.MinerPass != "" {
		state.ActiveMinerKey = forkhash.Argon2i([]byte(*cfg.MinerPass))
	}
}

func setDiallers(cfg *pod.Config, stateConfig *state.Config) {
	// Setup dial and DNS resolution (lookup) functions depending on the specified options. The default is to use the
	// standard net.DialTimeout function as well as the system DNS resolver. When a proxy is specified, the dial
	// function is set to the proxy specific dial function and the lookup is set to use tor (unless --noonion is
	// specified in which case the system DNS resolver is used).
	Trace("setting network dialer and lookup")
	stateConfig.Dial = net.DialTimeout
	stateConfig.Lookup = net.LookupIP
	if *cfg.Proxy != "" {
		Trace("we are loading a proxy!")
		_, _, err := net.SplitHostPort(*cfg.Proxy)
		if err != nil {
			Error(err)
			str := "%s: Proxy address '%s' is invalid: %v"
			err := fmt.Errorf(str, funcName, *cfg.Proxy, err)
			_, _ = fmt.Fprintln(os.Stderr, err)
			// os.Exit(1)
		}
		// Tor isolation flag means proxy credentials will be overridden unless there is also an onion proxy configured
		// in which case that one will be overridden.
		torIsolation := false
		if *cfg.TorIsolation &&
			*cfg.OnionProxy == "" &&
			(*cfg.ProxyUser != "" ||
				*cfg.ProxyPass != "") {
			torIsolation = true
			Warn("Tor isolation set -- overriding specified" +
				" proxy user credentials")
		}
		proxy := &socks.Proxy{
			Addr:         *cfg.Proxy,
			Username:     *cfg.ProxyUser,
			Password:     *cfg.ProxyPass,
			TorIsolation: torIsolation,
		}
		stateConfig.Dial = proxy.DialTimeout
		// Treat the proxy as tor and perform DNS resolution through it unless the --noonion flag is set or there is an
		// onion-specific proxy configured.
		if *cfg.Onion &&
			*cfg.OnionProxy == "" {
			stateConfig.Lookup = func(host string) ([]net.IP, error) {
				return connmgr.TorLookupIP(host, *cfg.Proxy)
			}
		}
	}
	// Setup onion address dial function depending on the specified options. The default is to use the same dial
	// function selected above. However, when an onion-specific proxy is specified, the onion address dial function is
	// set to use the onion-specific proxy while leaving the normal dial function as selected above. This allows .onion
	// address traffic to be routed through a different proxy than normal traffic.
	Trace("setting up tor proxy if enabled")
	if *cfg.OnionProxy != "" {
		_, _, err := net.SplitHostPort(*cfg.OnionProxy)
		if err != nil {
			Error(err)
			str := "%s: Onion proxy address '%s' is invalid: %v"
			err := fmt.Errorf(str, funcName, *cfg.OnionProxy, err)
			_, _ = fmt.Fprintln(os.Stderr, err)
			// os.Exit(1)
		}
		// Tor isolation flag means onion proxy credentials will be overridden.
		if *cfg.TorIsolation &&
			(*cfg.OnionProxyUser != "" || *cfg.OnionProxyPass != "") {
			Warn("Tor isolation set - overriding specified onionproxy user" +
				" credentials")
		}
	}
	Trace("setting onion dialer")
	stateConfig.Oniondial =
		func(network, addr string, timeout time.Duration) (net.Conn, error) {
			proxy := &socks.Proxy{
				Addr:         *cfg.OnionProxy,
				Username:     *cfg.OnionProxyUser,
				Password:     *cfg.OnionProxyPass,
				TorIsolation: *cfg.TorIsolation,
			}
			return proxy.DialTimeout(network, addr, timeout)
		}
	
	// When configured in bridge mode (both --onion and --proxy are configured), it means that the proxy configured by
	// --proxy is not a tor proxy, so override the DNS resolution to use the onion-specific proxy.
	Trace("setting proxy lookup")
	if *cfg.Proxy != "" {
		stateConfig.Lookup = func(host string) ([]net.IP, error) {
			return connmgr.TorLookupIP(host, *cfg.OnionProxy)
		}
	} else {
		stateConfig.Oniondial = stateConfig.Dial
	}
	// Specifying --noonion means the onion address dial function results in an error.
	if !*cfg.Onion {
		stateConfig.Oniondial = func(a, b string, t time.Duration) (net.Conn, error) {
			return nil, errors.New("tor has been disabled")
		}
	}
}
