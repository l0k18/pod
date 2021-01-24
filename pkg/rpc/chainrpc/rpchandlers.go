// generated; DO NOT EDIT
//go:generate go run genapi/genapi.go genapi/log.go

package chainrpc

import (
	"io"
	"net/rpc"
	"time"
	
	qu "github.com/l0k18/pod/pkg/util/quit"
	
	"github.com/l0k18/pod/pkg/rpc/btcjson"
)

// API stores the channel, parameters and result values from calls via
// the channel
type API struct {
	Ch     interface{}
	Params interface{}
	Result interface{}
}

// CAPI is the central structure for configuration and access to a 
// net/rpc API access endpoint for this RPC API
type CAPI struct {
	Timeout time.Duration
	quit    qu.C
}

// NewCAPI returns a new CAPI 
func NewCAPI(quit qu.C, timeout ...time.Duration) (c *CAPI) {
	c = &CAPI{quit: quit}
	if len(timeout) > 0 {
		c.Timeout = timeout[0]
	} else {
		c.Timeout = time.Second * 5
	}
	return
}

// Wrappers around RPC calls
type CAPIClient struct {
	*rpc.Client
}

// New creates a new client for a kopach_worker.
// Note that any kind of connection can be used here, other than the StdConn
func NewCAPIClient(conn io.ReadWriteCloser) *CAPIClient {
	return &CAPIClient{rpc.NewClient(conn)}
}

type (
	// None means no parameters it is not checked so it can be nil
	None struct{}
	// AddNodeRes is the result from a call to AddNode
	AddNodeRes struct {
		Res *None
		Err error
	}
	// CreateRawTransactionRes is the result from a call to CreateRawTransaction
	CreateRawTransactionRes struct {
		Res *string
		Err error
	}
	// DecodeRawTransactionRes is the result from a call to DecodeRawTransaction
	DecodeRawTransactionRes struct {
		Res *btcjson.TxRawDecodeResult
		Err error
	}
	// DecodeScriptRes is the result from a call to DecodeScript
	DecodeScriptRes struct {
		Res *btcjson.DecodeScriptResult
		Err error
	}
	// EstimateFeeRes is the result from a call to EstimateFee
	EstimateFeeRes struct {
		Res *float64
		Err error
	}
	// GenerateRes is the result from a call to Generate
	GenerateRes struct {
		Res *[]string
		Err error
	}
	// GetAddedNodeInfoRes is the result from a call to GetAddedNodeInfo
	GetAddedNodeInfoRes struct {
		Res *[]btcjson.GetAddedNodeInfoResultAddr
		Err error
	}
	// GetBestBlockRes is the result from a call to GetBestBlock
	GetBestBlockRes struct {
		Res *btcjson.GetBestBlockResult
		Err error
	}
	// GetBestBlockHashRes is the result from a call to GetBestBlockHash
	GetBestBlockHashRes struct {
		Res *string
		Err error
	}
	// GetBlockRes is the result from a call to GetBlock
	GetBlockRes struct {
		Res *btcjson.GetBlockVerboseResult
		Err error
	}
	// GetBlockChainInfoRes is the result from a call to GetBlockChainInfo
	GetBlockChainInfoRes struct {
		Res *btcjson.GetBlockChainInfoResult
		Err error
	}
	// GetBlockCountRes is the result from a call to GetBlockCount
	GetBlockCountRes struct {
		Res *int64
		Err error
	}
	// GetBlockHashRes is the result from a call to GetBlockHash
	GetBlockHashRes struct {
		Res *string
		Err error
	}
	// GetBlockHeaderRes is the result from a call to GetBlockHeader
	GetBlockHeaderRes struct {
		Res *btcjson.GetBlockHeaderVerboseResult
		Err error
	}
	// GetBlockTemplateRes is the result from a call to GetBlockTemplate
	GetBlockTemplateRes struct {
		Res *string
		Err error
	}
	// GetCFilterRes is the result from a call to GetCFilter
	GetCFilterRes struct {
		Res *string
		Err error
	}
	// GetCFilterHeaderRes is the result from a call to GetCFilterHeader
	GetCFilterHeaderRes struct {
		Res *string
		Err error
	}
	// GetConnectionCountRes is the result from a call to GetConnectionCount
	GetConnectionCountRes struct {
		Res *int32
		Err error
	}
	// GetCurrentNetRes is the result from a call to GetCurrentNet
	GetCurrentNetRes struct {
		Res *string
		Err error
	}
	// GetDifficultyRes is the result from a call to GetDifficulty
	GetDifficultyRes struct {
		Res *float64
		Err error
	}
	// GetGenerateRes is the result from a call to GetGenerate
	GetGenerateRes struct {
		Res *bool
		Err error
	}
	// GetHashesPerSecRes is the result from a call to GetHashesPerSec
	GetHashesPerSecRes struct {
		Res *float64
		Err error
	}
	// GetHeadersRes is the result from a call to GetHeaders
	GetHeadersRes struct {
		Res *[]string
		Err error
	}
	// GetInfoRes is the result from a call to GetInfo
	GetInfoRes struct {
		Res *btcjson.InfoChainResult0
		Err error
	}
	// GetMempoolInfoRes is the result from a call to GetMempoolInfo
	GetMempoolInfoRes struct {
		Res *btcjson.GetMempoolInfoResult
		Err error
	}
	// GetMiningInfoRes is the result from a call to GetMiningInfo
	GetMiningInfoRes struct {
		Res *btcjson.GetMiningInfoResult
		Err error
	}
	// GetNetTotalsRes is the result from a call to GetNetTotals
	GetNetTotalsRes struct {
		Res *btcjson.GetNetTotalsResult
		Err error
	}
	// GetNetworkHashPSRes is the result from a call to GetNetworkHashPS
	GetNetworkHashPSRes struct {
		Res *[]btcjson.GetPeerInfoResult
		Err error
	}
	// GetPeerInfoRes is the result from a call to GetPeerInfo
	GetPeerInfoRes struct {
		Res *[]btcjson.GetPeerInfoResult
		Err error
	}
	// GetRawMempoolRes is the result from a call to GetRawMempool
	GetRawMempoolRes struct {
		Res *[]string
		Err error
	}
	// GetRawTransactionRes is the result from a call to GetRawTransaction
	GetRawTransactionRes struct {
		Res *string
		Err error
	}
	// GetTxOutRes is the result from a call to GetTxOut
	GetTxOutRes struct {
		Res *string
		Err error
	}
	// HelpRes is the result from a call to Help
	HelpRes struct {
		Res *string
		Err error
	}
	// NodeRes is the result from a call to Node
	NodeRes struct {
		Res *None
		Err error
	}
	// PingRes is the result from a call to Ping
	PingRes struct {
		Res *None
		Err error
	}
	// ResetChainRes is the result from a call to ResetChain
	ResetChainRes struct {
		Res *None
		Err error
	}
	// RestartRes is the result from a call to Restart
	RestartRes struct {
		Res *None
		Err error
	}
	// SearchRawTransactionsRes is the result from a call to SearchRawTransactions
	SearchRawTransactionsRes struct {
		Res *[]btcjson.SearchRawTransactionsResult
		Err error
	}
	// SendRawTransactionRes is the result from a call to SendRawTransaction
	SendRawTransactionRes struct {
		Res *None
		Err error
	}
	// SetGenerateRes is the result from a call to SetGenerate
	SetGenerateRes struct {
		Res *None
		Err error
	}
	// StopRes is the result from a call to Stop
	StopRes struct {
		Res *None
		Err error
	}
	// SubmitBlockRes is the result from a call to SubmitBlock
	SubmitBlockRes struct {
		Res *string
		Err error
	}
	// UptimeRes is the result from a call to Uptime
	UptimeRes struct {
		Res *btcjson.GetMempoolInfoResult
		Err error
	}
	// ValidateAddressRes is the result from a call to ValidateAddress
	ValidateAddressRes struct {
		Res *btcjson.ValidateAddressChainResult
		Err error
	}
	// VerifyChainRes is the result from a call to VerifyChain
	VerifyChainRes struct {
		Res *bool
		Err error
	}
	// VerifyMessageRes is the result from a call to VerifyMessage
	VerifyMessageRes struct {
		Res *bool
		Err error
	}
	// VersionRes is the result from a call to Version
	VersionRes struct {
		Res *map[string]btcjson.VersionResult
		Err error
	}
)

// RPCHandlersBeforeInit are created first and are added to the main list 
// when the init runs.
//
// - Fn is the handler function
// 
// - Call is a channel carrying a struct containing parameters and error that is 
// listened to in RunAPI to dispatch the calls
// 
// - Result is a bundle of command parameters and a channel that the result will be sent 
// back on
//
// Get and save the Result function's return, and you can then call the call functions
// check, result and wait functions for asynchronous and synchronous calls to RPC functions
var RPCHandlersBeforeInit = map[string]CommandHandler{
	"addnode": {
		Fn: HandleAddNode, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan AddNodeRes)} }},
	"createrawtransaction": {
		Fn: HandleCreateRawTransaction, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan CreateRawTransactionRes)} }},
	"decoderawtransaction": {
		Fn: HandleDecodeRawTransaction, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan DecodeRawTransactionRes)} }},
	"decodescript": {
		Fn: HandleDecodeScript, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan DecodeScriptRes)} }},
	"estimatefee": {
		Fn: HandleEstimateFee, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan EstimateFeeRes)} }},
	"generate": {
		Fn: HandleGenerate, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GenerateRes)} }},
	"getaddednodeinfo": {
		Fn: HandleGetAddedNodeInfo, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetAddedNodeInfoRes)} }},
	"getbestblock": {
		Fn: HandleGetBestBlock, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetBestBlockRes)} }},
	"getbestblockhash": {
		Fn: HandleGetBestBlockHash, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetBestBlockHashRes)} }},
	"getblock": {
		Fn: HandleGetBlock, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetBlockRes)} }},
	"getblockchaininfo": {
		Fn: HandleGetBlockChainInfo, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetBlockChainInfoRes)} }},
	"getblockcount": {
		Fn: HandleGetBlockCount, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetBlockCountRes)} }},
	"getblockhash": {
		Fn: HandleGetBlockHash, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetBlockHashRes)} }},
	"getblockheader": {
		Fn: HandleGetBlockHeader, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetBlockHeaderRes)} }},
	"getblocktemplate": {
		Fn: HandleGetBlockTemplate, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetBlockTemplateRes)} }},
	"getcfilter": {
		Fn: HandleGetCFilter, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetCFilterRes)} }},
	"getcfilterheader": {
		Fn: HandleGetCFilterHeader, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetCFilterHeaderRes)} }},
	"getconnectioncount": {
		Fn: HandleGetConnectionCount, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetConnectionCountRes)} }},
	"getcurrentnet": {
		Fn: HandleGetCurrentNet, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetCurrentNetRes)} }},
	"getdifficulty": {
		Fn: HandleGetDifficulty, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetDifficultyRes)} }},
	"getgenerate": {
		Fn: HandleGetGenerate, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetGenerateRes)} }},
	"gethashespersec": {
		Fn: HandleGetHashesPerSec, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetHashesPerSecRes)} }},
	"getheaders": {
		Fn: HandleGetHeaders, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetHeadersRes)} }},
	"getinfo": {
		Fn: HandleGetInfo, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetInfoRes)} }},
	"getmempoolinfo": {
		Fn: HandleGetMempoolInfo, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetMempoolInfoRes)} }},
	"getmininginfo": {
		Fn: HandleGetMiningInfo, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetMiningInfoRes)} }},
	"getnettotals": {
		Fn: HandleGetNetTotals, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetNetTotalsRes)} }},
	"getnetworkhashps": {
		Fn: HandleGetNetworkHashPS, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetNetworkHashPSRes)} }},
	"getpeerinfo": {
		Fn: HandleGetPeerInfo, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetPeerInfoRes)} }},
	"getrawmempool": {
		Fn: HandleGetRawMempool, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetRawMempoolRes)} }},
	"getrawtransaction": {
		Fn: HandleGetRawTransaction, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetRawTransactionRes)} }},
	"gettxout": {
		Fn: HandleGetTxOut, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan GetTxOutRes)} }},
	"help": {
		Fn: HandleHelp, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan HelpRes)} }},
	"node": {
		Fn: HandleNode, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan NodeRes)} }},
	"ping": {
		Fn: HandlePing, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan PingRes)} }},
	"resetchain": {
		Fn: HandleResetChain, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan ResetChainRes)} }},
	"restart": {
		Fn: HandleRestart, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan RestartRes)} }},
	"searchrawtransactions": {
		Fn: HandleSearchRawTransactions, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan SearchRawTransactionsRes)} }},
	"sendrawtransaction": {
		Fn: HandleSendRawTransaction, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan SendRawTransactionRes)} }},
	"setgenerate": {
		Fn: HandleSetGenerate, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan SetGenerateRes)} }},
	"stop": {
		Fn: HandleStop, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan StopRes)} }},
	"submitblock": {
		Fn: HandleSubmitBlock, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan SubmitBlockRes)} }},
	"uptime": {
		Fn: HandleUptime, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan UptimeRes)} }},
	"validateaddress": {
		Fn: HandleValidateAddress, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan ValidateAddressRes)} }},
	"verifychain": {
		Fn: HandleVerifyChain, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan VerifyChainRes)} }},
	"verifymessage": {
		Fn: HandleVerifyMessage, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan VerifyMessageRes)} }},
	"version": {
		Fn: HandleVersion, Call: make(chan API, 32),
		Result: func() API { return API{Ch: make(chan VersionRes)} }},
}

// API functions
//
// The functions here provide access to the RPC through a convenient set of functions
// generated for each call in the RPC API to request, check for, access the results and
// wait on results

// AddNode calls the method with the given parameters
func (a API) AddNode(cmd *btcjson.AddNodeCmd) (err error) {
	RPCHandlers["addnode"].Call <- API{a.Ch, cmd, nil}
	return
}

// AddNodeCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) AddNodeCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan AddNodeRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// AddNodeGetRes returns a pointer to the value in the Result field
func (a API) AddNodeGetRes() (out *None, err error) {
	out, _ = a.Result.(*None)
	err, _ = a.Result.(error)
	return
}

// AddNodeWait calls the method and blocks until it returns or 5 seconds passes
func (a API) AddNodeWait(cmd *btcjson.AddNodeCmd) (out *None, err error) {
	RPCHandlers["addnode"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan AddNodeRes):
		out, err = o.Res, o.Err
	}
	return
}

// CreateRawTransaction calls the method with the given parameters
func (a API) CreateRawTransaction(cmd *btcjson.CreateRawTransactionCmd) (err error) {
	RPCHandlers["createrawtransaction"].Call <- API{a.Ch, cmd, nil}
	return
}

// CreateRawTransactionCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) CreateRawTransactionCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan CreateRawTransactionRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// CreateRawTransactionGetRes returns a pointer to the value in the Result field
func (a API) CreateRawTransactionGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// CreateRawTransactionWait calls the method and blocks until it returns or 5 seconds passes
func (a API) CreateRawTransactionWait(cmd *btcjson.CreateRawTransactionCmd) (out *string, err error) {
	RPCHandlers["createrawtransaction"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan CreateRawTransactionRes):
		out, err = o.Res, o.Err
	}
	return
}

// DecodeRawTransaction calls the method with the given parameters
func (a API) DecodeRawTransaction(cmd *btcjson.DecodeRawTransactionCmd) (err error) {
	RPCHandlers["decoderawtransaction"].Call <- API{a.Ch, cmd, nil}
	return
}

// DecodeRawTransactionCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) DecodeRawTransactionCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan DecodeRawTransactionRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// DecodeRawTransactionGetRes returns a pointer to the value in the Result field
func (a API) DecodeRawTransactionGetRes() (out *btcjson.TxRawDecodeResult, err error) {
	out, _ = a.Result.(*btcjson.TxRawDecodeResult)
	err, _ = a.Result.(error)
	return
}

// DecodeRawTransactionWait calls the method and blocks until it returns or 5 seconds passes
func (a API) DecodeRawTransactionWait(cmd *btcjson.DecodeRawTransactionCmd) (out *btcjson.TxRawDecodeResult, err error) {
	RPCHandlers["decoderawtransaction"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan DecodeRawTransactionRes):
		out, err = o.Res, o.Err
	}
	return
}

// DecodeScript calls the method with the given parameters
func (a API) DecodeScript(cmd *btcjson.DecodeScriptCmd) (err error) {
	RPCHandlers["decodescript"].Call <- API{a.Ch, cmd, nil}
	return
}

// DecodeScriptCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) DecodeScriptCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan DecodeScriptRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// DecodeScriptGetRes returns a pointer to the value in the Result field
func (a API) DecodeScriptGetRes() (out *btcjson.DecodeScriptResult, err error) {
	out, _ = a.Result.(*btcjson.DecodeScriptResult)
	err, _ = a.Result.(error)
	return
}

// DecodeScriptWait calls the method and blocks until it returns or 5 seconds passes
func (a API) DecodeScriptWait(cmd *btcjson.DecodeScriptCmd) (out *btcjson.DecodeScriptResult, err error) {
	RPCHandlers["decodescript"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan DecodeScriptRes):
		out, err = o.Res, o.Err
	}
	return
}

// EstimateFee calls the method with the given parameters
func (a API) EstimateFee(cmd *btcjson.EstimateFeeCmd) (err error) {
	RPCHandlers["estimatefee"].Call <- API{a.Ch, cmd, nil}
	return
}

// EstimateFeeCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) EstimateFeeCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan EstimateFeeRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// EstimateFeeGetRes returns a pointer to the value in the Result field
func (a API) EstimateFeeGetRes() (out *float64, err error) {
	out, _ = a.Result.(*float64)
	err, _ = a.Result.(error)
	return
}

// EstimateFeeWait calls the method and blocks until it returns or 5 seconds passes
func (a API) EstimateFeeWait(cmd *btcjson.EstimateFeeCmd) (out *float64, err error) {
	RPCHandlers["estimatefee"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan EstimateFeeRes):
		out, err = o.Res, o.Err
	}
	return
}

// Generate calls the method with the given parameters
func (a API) Generate(cmd *None) (err error) {
	RPCHandlers["generate"].Call <- API{a.Ch, cmd, nil}
	return
}

// GenerateCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GenerateCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GenerateRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GenerateGetRes returns a pointer to the value in the Result field
func (a API) GenerateGetRes() (out *[]string, err error) {
	out, _ = a.Result.(*[]string)
	err, _ = a.Result.(error)
	return
}

// GenerateWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GenerateWait(cmd *None) (out *[]string, err error) {
	RPCHandlers["generate"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GenerateRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetAddedNodeInfo calls the method with the given parameters
func (a API) GetAddedNodeInfo(cmd *btcjson.GetAddedNodeInfoCmd) (err error) {
	RPCHandlers["getaddednodeinfo"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetAddedNodeInfoCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetAddedNodeInfoCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetAddedNodeInfoRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetAddedNodeInfoGetRes returns a pointer to the value in the Result field
func (a API) GetAddedNodeInfoGetRes() (out *[]btcjson.GetAddedNodeInfoResultAddr, err error) {
	out, _ = a.Result.(*[]btcjson.GetAddedNodeInfoResultAddr)
	err, _ = a.Result.(error)
	return
}

// GetAddedNodeInfoWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetAddedNodeInfoWait(cmd *btcjson.GetAddedNodeInfoCmd) (out *[]btcjson.GetAddedNodeInfoResultAddr, err error) {
	RPCHandlers["getaddednodeinfo"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetAddedNodeInfoRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetBestBlock calls the method with the given parameters
func (a API) GetBestBlock(cmd *None) (err error) {
	RPCHandlers["getbestblock"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetBestBlockCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetBestBlockCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetBestBlockRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetBestBlockGetRes returns a pointer to the value in the Result field
func (a API) GetBestBlockGetRes() (out *btcjson.GetBestBlockResult, err error) {
	out, _ = a.Result.(*btcjson.GetBestBlockResult)
	err, _ = a.Result.(error)
	return
}

// GetBestBlockWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetBestBlockWait(cmd *None) (out *btcjson.GetBestBlockResult, err error) {
	RPCHandlers["getbestblock"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetBestBlockRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetBestBlockHash calls the method with the given parameters
func (a API) GetBestBlockHash(cmd *None) (err error) {
	RPCHandlers["getbestblockhash"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetBestBlockHashCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetBestBlockHashCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetBestBlockHashRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetBestBlockHashGetRes returns a pointer to the value in the Result field
func (a API) GetBestBlockHashGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// GetBestBlockHashWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetBestBlockHashWait(cmd *None) (out *string, err error) {
	RPCHandlers["getbestblockhash"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetBestBlockHashRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetBlock calls the method with the given parameters
func (a API) GetBlock(cmd *btcjson.GetBlockCmd) (err error) {
	RPCHandlers["getblock"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetBlockCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetBlockCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetBlockRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetBlockGetRes returns a pointer to the value in the Result field
func (a API) GetBlockGetRes() (out *btcjson.GetBlockVerboseResult, err error) {
	out, _ = a.Result.(*btcjson.GetBlockVerboseResult)
	err, _ = a.Result.(error)
	return
}

// GetBlockWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetBlockWait(cmd *btcjson.GetBlockCmd) (out *btcjson.GetBlockVerboseResult, err error) {
	RPCHandlers["getblock"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetBlockRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetBlockChainInfo calls the method with the given parameters
func (a API) GetBlockChainInfo(cmd *None) (err error) {
	RPCHandlers["getblockchaininfo"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetBlockChainInfoCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetBlockChainInfoCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetBlockChainInfoRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetBlockChainInfoGetRes returns a pointer to the value in the Result field
func (a API) GetBlockChainInfoGetRes() (out *btcjson.GetBlockChainInfoResult, err error) {
	out, _ = a.Result.(*btcjson.GetBlockChainInfoResult)
	err, _ = a.Result.(error)
	return
}

// GetBlockChainInfoWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetBlockChainInfoWait(cmd *None) (out *btcjson.GetBlockChainInfoResult, err error) {
	RPCHandlers["getblockchaininfo"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetBlockChainInfoRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetBlockCount calls the method with the given parameters
func (a API) GetBlockCount(cmd *None) (err error) {
	RPCHandlers["getblockcount"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetBlockCountCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetBlockCountCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetBlockCountRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetBlockCountGetRes returns a pointer to the value in the Result field
func (a API) GetBlockCountGetRes() (out *int64, err error) {
	out, _ = a.Result.(*int64)
	err, _ = a.Result.(error)
	return
}

// GetBlockCountWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetBlockCountWait(cmd *None) (out *int64, err error) {
	RPCHandlers["getblockcount"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetBlockCountRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetBlockHash calls the method with the given parameters
func (a API) GetBlockHash(cmd *btcjson.GetBlockHashCmd) (err error) {
	RPCHandlers["getblockhash"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetBlockHashCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetBlockHashCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetBlockHashRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetBlockHashGetRes returns a pointer to the value in the Result field
func (a API) GetBlockHashGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// GetBlockHashWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetBlockHashWait(cmd *btcjson.GetBlockHashCmd) (out *string, err error) {
	RPCHandlers["getblockhash"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetBlockHashRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetBlockHeader calls the method with the given parameters
func (a API) GetBlockHeader(cmd *btcjson.GetBlockHeaderCmd) (err error) {
	RPCHandlers["getblockheader"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetBlockHeaderCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetBlockHeaderCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetBlockHeaderRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetBlockHeaderGetRes returns a pointer to the value in the Result field
func (a API) GetBlockHeaderGetRes() (out *btcjson.GetBlockHeaderVerboseResult, err error) {
	out, _ = a.Result.(*btcjson.GetBlockHeaderVerboseResult)
	err, _ = a.Result.(error)
	return
}

// GetBlockHeaderWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetBlockHeaderWait(cmd *btcjson.GetBlockHeaderCmd) (out *btcjson.GetBlockHeaderVerboseResult, err error) {
	RPCHandlers["getblockheader"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetBlockHeaderRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetBlockTemplate calls the method with the given parameters
func (a API) GetBlockTemplate(cmd *btcjson.GetBlockTemplateCmd) (err error) {
	RPCHandlers["getblocktemplate"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetBlockTemplateCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetBlockTemplateCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetBlockTemplateRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetBlockTemplateGetRes returns a pointer to the value in the Result field
func (a API) GetBlockTemplateGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// GetBlockTemplateWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetBlockTemplateWait(cmd *btcjson.GetBlockTemplateCmd) (out *string, err error) {
	RPCHandlers["getblocktemplate"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetBlockTemplateRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetCFilter calls the method with the given parameters
func (a API) GetCFilter(cmd *btcjson.GetCFilterCmd) (err error) {
	RPCHandlers["getcfilter"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetCFilterCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetCFilterCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetCFilterRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetCFilterGetRes returns a pointer to the value in the Result field
func (a API) GetCFilterGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// GetCFilterWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetCFilterWait(cmd *btcjson.GetCFilterCmd) (out *string, err error) {
	RPCHandlers["getcfilter"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetCFilterRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetCFilterHeader calls the method with the given parameters
func (a API) GetCFilterHeader(cmd *btcjson.GetCFilterHeaderCmd) (err error) {
	RPCHandlers["getcfilterheader"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetCFilterHeaderCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetCFilterHeaderCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetCFilterHeaderRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetCFilterHeaderGetRes returns a pointer to the value in the Result field
func (a API) GetCFilterHeaderGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// GetCFilterHeaderWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetCFilterHeaderWait(cmd *btcjson.GetCFilterHeaderCmd) (out *string, err error) {
	RPCHandlers["getcfilterheader"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetCFilterHeaderRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetConnectionCount calls the method with the given parameters
func (a API) GetConnectionCount(cmd *None) (err error) {
	RPCHandlers["getconnectioncount"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetConnectionCountCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetConnectionCountCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetConnectionCountRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetConnectionCountGetRes returns a pointer to the value in the Result field
func (a API) GetConnectionCountGetRes() (out *int32, err error) {
	out, _ = a.Result.(*int32)
	err, _ = a.Result.(error)
	return
}

// GetConnectionCountWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetConnectionCountWait(cmd *None) (out *int32, err error) {
	RPCHandlers["getconnectioncount"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetConnectionCountRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetCurrentNet calls the method with the given parameters
func (a API) GetCurrentNet(cmd *None) (err error) {
	RPCHandlers["getcurrentnet"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetCurrentNetCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetCurrentNetCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetCurrentNetRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetCurrentNetGetRes returns a pointer to the value in the Result field
func (a API) GetCurrentNetGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// GetCurrentNetWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetCurrentNetWait(cmd *None) (out *string, err error) {
	RPCHandlers["getcurrentnet"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetCurrentNetRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetDifficulty calls the method with the given parameters
func (a API) GetDifficulty(cmd *btcjson.GetDifficultyCmd) (err error) {
	RPCHandlers["getdifficulty"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetDifficultyCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetDifficultyCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetDifficultyRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetDifficultyGetRes returns a pointer to the value in the Result field
func (a API) GetDifficultyGetRes() (out *float64, err error) {
	out, _ = a.Result.(*float64)
	err, _ = a.Result.(error)
	return
}

// GetDifficultyWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetDifficultyWait(cmd *btcjson.GetDifficultyCmd) (out *float64, err error) {
	RPCHandlers["getdifficulty"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetDifficultyRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetGenerate calls the method with the given parameters
func (a API) GetGenerate(cmd *btcjson.GetHeadersCmd) (err error) {
	RPCHandlers["getgenerate"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetGenerateCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetGenerateCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetGenerateRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetGenerateGetRes returns a pointer to the value in the Result field
func (a API) GetGenerateGetRes() (out *bool, err error) {
	out, _ = a.Result.(*bool)
	err, _ = a.Result.(error)
	return
}

// GetGenerateWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetGenerateWait(cmd *btcjson.GetHeadersCmd) (out *bool, err error) {
	RPCHandlers["getgenerate"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetGenerateRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetHashesPerSec calls the method with the given parameters
func (a API) GetHashesPerSec(cmd *None) (err error) {
	RPCHandlers["gethashespersec"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetHashesPerSecCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetHashesPerSecCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetHashesPerSecRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetHashesPerSecGetRes returns a pointer to the value in the Result field
func (a API) GetHashesPerSecGetRes() (out *float64, err error) {
	out, _ = a.Result.(*float64)
	err, _ = a.Result.(error)
	return
}

// GetHashesPerSecWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetHashesPerSecWait(cmd *None) (out *float64, err error) {
	RPCHandlers["gethashespersec"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetHashesPerSecRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetHeaders calls the method with the given parameters
func (a API) GetHeaders(cmd *btcjson.GetHeadersCmd) (err error) {
	RPCHandlers["getheaders"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetHeadersCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetHeadersCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetHeadersRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetHeadersGetRes returns a pointer to the value in the Result field
func (a API) GetHeadersGetRes() (out *[]string, err error) {
	out, _ = a.Result.(*[]string)
	err, _ = a.Result.(error)
	return
}

// GetHeadersWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetHeadersWait(cmd *btcjson.GetHeadersCmd) (out *[]string, err error) {
	RPCHandlers["getheaders"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetHeadersRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetInfo calls the method with the given parameters
func (a API) GetInfo(cmd *None) (err error) {
	RPCHandlers["getinfo"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetInfoCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetInfoCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetInfoRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetInfoGetRes returns a pointer to the value in the Result field
func (a API) GetInfoGetRes() (out *btcjson.InfoChainResult0, err error) {
	out, _ = a.Result.(*btcjson.InfoChainResult0)
	err, _ = a.Result.(error)
	return
}

// GetInfoWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetInfoWait(cmd *None) (out *btcjson.InfoChainResult0, err error) {
	RPCHandlers["getinfo"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetInfoRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetMempoolInfo calls the method with the given parameters
func (a API) GetMempoolInfo(cmd *None) (err error) {
	RPCHandlers["getmempoolinfo"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetMempoolInfoCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetMempoolInfoCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetMempoolInfoRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetMempoolInfoGetRes returns a pointer to the value in the Result field
func (a API) GetMempoolInfoGetRes() (out *btcjson.GetMempoolInfoResult, err error) {
	out, _ = a.Result.(*btcjson.GetMempoolInfoResult)
	err, _ = a.Result.(error)
	return
}

// GetMempoolInfoWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetMempoolInfoWait(cmd *None) (out *btcjson.GetMempoolInfoResult, err error) {
	RPCHandlers["getmempoolinfo"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetMempoolInfoRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetMiningInfo calls the method with the given parameters
func (a API) GetMiningInfo(cmd *None) (err error) {
	RPCHandlers["getmininginfo"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetMiningInfoCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetMiningInfoCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetMiningInfoRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetMiningInfoGetRes returns a pointer to the value in the Result field
func (a API) GetMiningInfoGetRes() (out *btcjson.GetMiningInfoResult, err error) {
	out, _ = a.Result.(*btcjson.GetMiningInfoResult)
	err, _ = a.Result.(error)
	return
}

// GetMiningInfoWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetMiningInfoWait(cmd *None) (out *btcjson.GetMiningInfoResult, err error) {
	RPCHandlers["getmininginfo"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetMiningInfoRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetNetTotals calls the method with the given parameters
func (a API) GetNetTotals(cmd *None) (err error) {
	RPCHandlers["getnettotals"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetNetTotalsCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetNetTotalsCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetNetTotalsRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetNetTotalsGetRes returns a pointer to the value in the Result field
func (a API) GetNetTotalsGetRes() (out *btcjson.GetNetTotalsResult, err error) {
	out, _ = a.Result.(*btcjson.GetNetTotalsResult)
	err, _ = a.Result.(error)
	return
}

// GetNetTotalsWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetNetTotalsWait(cmd *None) (out *btcjson.GetNetTotalsResult, err error) {
	RPCHandlers["getnettotals"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetNetTotalsRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetNetworkHashPS calls the method with the given parameters
func (a API) GetNetworkHashPS(cmd *btcjson.GetNetworkHashPSCmd) (err error) {
	RPCHandlers["getnetworkhashps"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetNetworkHashPSCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetNetworkHashPSCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetNetworkHashPSRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetNetworkHashPSGetRes returns a pointer to the value in the Result field
func (a API) GetNetworkHashPSGetRes() (out *[]btcjson.GetPeerInfoResult, err error) {
	out, _ = a.Result.(*[]btcjson.GetPeerInfoResult)
	err, _ = a.Result.(error)
	return
}

// GetNetworkHashPSWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetNetworkHashPSWait(cmd *btcjson.GetNetworkHashPSCmd) (out *[]btcjson.GetPeerInfoResult, err error) {
	RPCHandlers["getnetworkhashps"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetNetworkHashPSRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetPeerInfo calls the method with the given parameters
func (a API) GetPeerInfo(cmd *None) (err error) {
	RPCHandlers["getpeerinfo"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetPeerInfoCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetPeerInfoCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetPeerInfoRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetPeerInfoGetRes returns a pointer to the value in the Result field
func (a API) GetPeerInfoGetRes() (out *[]btcjson.GetPeerInfoResult, err error) {
	out, _ = a.Result.(*[]btcjson.GetPeerInfoResult)
	err, _ = a.Result.(error)
	return
}

// GetPeerInfoWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetPeerInfoWait(cmd *None) (out *[]btcjson.GetPeerInfoResult, err error) {
	RPCHandlers["getpeerinfo"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetPeerInfoRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetRawMempool calls the method with the given parameters
func (a API) GetRawMempool(cmd *btcjson.GetRawMempoolCmd) (err error) {
	RPCHandlers["getrawmempool"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetRawMempoolCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetRawMempoolCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetRawMempoolRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetRawMempoolGetRes returns a pointer to the value in the Result field
func (a API) GetRawMempoolGetRes() (out *[]string, err error) {
	out, _ = a.Result.(*[]string)
	err, _ = a.Result.(error)
	return
}

// GetRawMempoolWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetRawMempoolWait(cmd *btcjson.GetRawMempoolCmd) (out *[]string, err error) {
	RPCHandlers["getrawmempool"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetRawMempoolRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetRawTransaction calls the method with the given parameters
func (a API) GetRawTransaction(cmd *btcjson.GetRawTransactionCmd) (err error) {
	RPCHandlers["getrawtransaction"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetRawTransactionCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetRawTransactionCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetRawTransactionRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetRawTransactionGetRes returns a pointer to the value in the Result field
func (a API) GetRawTransactionGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// GetRawTransactionWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetRawTransactionWait(cmd *btcjson.GetRawTransactionCmd) (out *string, err error) {
	RPCHandlers["getrawtransaction"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetRawTransactionRes):
		out, err = o.Res, o.Err
	}
	return
}

// GetTxOut calls the method with the given parameters
func (a API) GetTxOut(cmd *btcjson.GetTxOutCmd) (err error) {
	RPCHandlers["gettxout"].Call <- API{a.Ch, cmd, nil}
	return
}

// GetTxOutCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) GetTxOutCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan GetTxOutRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// GetTxOutGetRes returns a pointer to the value in the Result field
func (a API) GetTxOutGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// GetTxOutWait calls the method and blocks until it returns or 5 seconds passes
func (a API) GetTxOutWait(cmd *btcjson.GetTxOutCmd) (out *string, err error) {
	RPCHandlers["gettxout"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan GetTxOutRes):
		out, err = o.Res, o.Err
	}
	return
}

// Help calls the method with the given parameters
func (a API) Help(cmd *btcjson.HelpCmd) (err error) {
	RPCHandlers["help"].Call <- API{a.Ch, cmd, nil}
	return
}

// HelpCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) HelpCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan HelpRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// HelpGetRes returns a pointer to the value in the Result field
func (a API) HelpGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// HelpWait calls the method and blocks until it returns or 5 seconds passes
func (a API) HelpWait(cmd *btcjson.HelpCmd) (out *string, err error) {
	RPCHandlers["help"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan HelpRes):
		out, err = o.Res, o.Err
	}
	return
}

// Node calls the method with the given parameters
func (a API) Node(cmd *btcjson.NodeCmd) (err error) {
	RPCHandlers["node"].Call <- API{a.Ch, cmd, nil}
	return
}

// NodeCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) NodeCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan NodeRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// NodeGetRes returns a pointer to the value in the Result field
func (a API) NodeGetRes() (out *None, err error) {
	out, _ = a.Result.(*None)
	err, _ = a.Result.(error)
	return
}

// NodeWait calls the method and blocks until it returns or 5 seconds passes
func (a API) NodeWait(cmd *btcjson.NodeCmd) (out *None, err error) {
	RPCHandlers["node"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan NodeRes):
		out, err = o.Res, o.Err
	}
	return
}

// Ping calls the method with the given parameters
func (a API) Ping(cmd *None) (err error) {
	RPCHandlers["ping"].Call <- API{a.Ch, cmd, nil}
	return
}

// PingCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) PingCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan PingRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// PingGetRes returns a pointer to the value in the Result field
func (a API) PingGetRes() (out *None, err error) {
	out, _ = a.Result.(*None)
	err, _ = a.Result.(error)
	return
}

// PingWait calls the method and blocks until it returns or 5 seconds passes
func (a API) PingWait(cmd *None) (out *None, err error) {
	RPCHandlers["ping"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan PingRes):
		out, err = o.Res, o.Err
	}
	return
}

// ResetChain calls the method with the given parameters
func (a API) ResetChain(cmd *None) (err error) {
	RPCHandlers["resetchain"].Call <- API{a.Ch, cmd, nil}
	return
}

// ResetChainCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) ResetChainCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan ResetChainRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// ResetChainGetRes returns a pointer to the value in the Result field
func (a API) ResetChainGetRes() (out *None, err error) {
	out, _ = a.Result.(*None)
	err, _ = a.Result.(error)
	return
}

// ResetChainWait calls the method and blocks until it returns or 5 seconds passes
func (a API) ResetChainWait(cmd *None) (out *None, err error) {
	RPCHandlers["resetchain"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan ResetChainRes):
		out, err = o.Res, o.Err
	}
	return
}

// Restart calls the method with the given parameters
func (a API) Restart(cmd *None) (err error) {
	RPCHandlers["restart"].Call <- API{a.Ch, cmd, nil}
	return
}

// RestartCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) RestartCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan RestartRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// RestartGetRes returns a pointer to the value in the Result field
func (a API) RestartGetRes() (out *None, err error) {
	out, _ = a.Result.(*None)
	err, _ = a.Result.(error)
	return
}

// RestartWait calls the method and blocks until it returns or 5 seconds passes
func (a API) RestartWait(cmd *None) (out *None, err error) {
	RPCHandlers["restart"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan RestartRes):
		out, err = o.Res, o.Err
	}
	return
}

// SearchRawTransactions calls the method with the given parameters
func (a API) SearchRawTransactions(cmd *btcjson.SearchRawTransactionsCmd) (err error) {
	RPCHandlers["searchrawtransactions"].Call <- API{a.Ch, cmd, nil}
	return
}

// SearchRawTransactionsCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) SearchRawTransactionsCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan SearchRawTransactionsRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// SearchRawTransactionsGetRes returns a pointer to the value in the Result field
func (a API) SearchRawTransactionsGetRes() (out *[]btcjson.SearchRawTransactionsResult, err error) {
	out, _ = a.Result.(*[]btcjson.SearchRawTransactionsResult)
	err, _ = a.Result.(error)
	return
}

// SearchRawTransactionsWait calls the method and blocks until it returns or 5 seconds passes
func (a API) SearchRawTransactionsWait(cmd *btcjson.SearchRawTransactionsCmd) (out *[]btcjson.SearchRawTransactionsResult, err error) {
	RPCHandlers["searchrawtransactions"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan SearchRawTransactionsRes):
		out, err = o.Res, o.Err
	}
	return
}

// SendRawTransaction calls the method with the given parameters
func (a API) SendRawTransaction(cmd *btcjson.SendRawTransactionCmd) (err error) {
	RPCHandlers["sendrawtransaction"].Call <- API{a.Ch, cmd, nil}
	return
}

// SendRawTransactionCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) SendRawTransactionCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan SendRawTransactionRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// SendRawTransactionGetRes returns a pointer to the value in the Result field
func (a API) SendRawTransactionGetRes() (out *None, err error) {
	out, _ = a.Result.(*None)
	err, _ = a.Result.(error)
	return
}

// SendRawTransactionWait calls the method and blocks until it returns or 5 seconds passes
func (a API) SendRawTransactionWait(cmd *btcjson.SendRawTransactionCmd) (out *None, err error) {
	RPCHandlers["sendrawtransaction"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan SendRawTransactionRes):
		out, err = o.Res, o.Err
	}
	return
}

// SetGenerate calls the method with the given parameters
func (a API) SetGenerate(cmd *btcjson.SetGenerateCmd) (err error) {
	RPCHandlers["setgenerate"].Call <- API{a.Ch, cmd, nil}
	return
}

// SetGenerateCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) SetGenerateCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan SetGenerateRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// SetGenerateGetRes returns a pointer to the value in the Result field
func (a API) SetGenerateGetRes() (out *None, err error) {
	out, _ = a.Result.(*None)
	err, _ = a.Result.(error)
	return
}

// SetGenerateWait calls the method and blocks until it returns or 5 seconds passes
func (a API) SetGenerateWait(cmd *btcjson.SetGenerateCmd) (out *None, err error) {
	RPCHandlers["setgenerate"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan SetGenerateRes):
		out, err = o.Res, o.Err
	}
	return
}

// Stop calls the method with the given parameters
func (a API) Stop(cmd *None) (err error) {
	RPCHandlers["stop"].Call <- API{a.Ch, cmd, nil}
	return
}

// StopCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) StopCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan StopRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// StopGetRes returns a pointer to the value in the Result field
func (a API) StopGetRes() (out *None, err error) {
	out, _ = a.Result.(*None)
	err, _ = a.Result.(error)
	return
}

// StopWait calls the method and blocks until it returns or 5 seconds passes
func (a API) StopWait(cmd *None) (out *None, err error) {
	RPCHandlers["stop"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan StopRes):
		out, err = o.Res, o.Err
	}
	return
}

// SubmitBlock calls the method with the given parameters
func (a API) SubmitBlock(cmd *btcjson.SubmitBlockCmd) (err error) {
	RPCHandlers["submitblock"].Call <- API{a.Ch, cmd, nil}
	return
}

// SubmitBlockCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) SubmitBlockCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan SubmitBlockRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// SubmitBlockGetRes returns a pointer to the value in the Result field
func (a API) SubmitBlockGetRes() (out *string, err error) {
	out, _ = a.Result.(*string)
	err, _ = a.Result.(error)
	return
}

// SubmitBlockWait calls the method and blocks until it returns or 5 seconds passes
func (a API) SubmitBlockWait(cmd *btcjson.SubmitBlockCmd) (out *string, err error) {
	RPCHandlers["submitblock"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan SubmitBlockRes):
		out, err = o.Res, o.Err
	}
	return
}

// Uptime calls the method with the given parameters
func (a API) Uptime(cmd *None) (err error) {
	RPCHandlers["uptime"].Call <- API{a.Ch, cmd, nil}
	return
}

// UptimeCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) UptimeCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan UptimeRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// UptimeGetRes returns a pointer to the value in the Result field
func (a API) UptimeGetRes() (out *btcjson.GetMempoolInfoResult, err error) {
	out, _ = a.Result.(*btcjson.GetMempoolInfoResult)
	err, _ = a.Result.(error)
	return
}

// UptimeWait calls the method and blocks until it returns or 5 seconds passes
func (a API) UptimeWait(cmd *None) (out *btcjson.GetMempoolInfoResult, err error) {
	RPCHandlers["uptime"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan UptimeRes):
		out, err = o.Res, o.Err
	}
	return
}

// ValidateAddress calls the method with the given parameters
func (a API) ValidateAddress(cmd *btcjson.ValidateAddressCmd) (err error) {
	RPCHandlers["validateaddress"].Call <- API{a.Ch, cmd, nil}
	return
}

// ValidateAddressCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) ValidateAddressCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan ValidateAddressRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// ValidateAddressGetRes returns a pointer to the value in the Result field
func (a API) ValidateAddressGetRes() (out *btcjson.ValidateAddressChainResult, err error) {
	out, _ = a.Result.(*btcjson.ValidateAddressChainResult)
	err, _ = a.Result.(error)
	return
}

// ValidateAddressWait calls the method and blocks until it returns or 5 seconds passes
func (a API) ValidateAddressWait(cmd *btcjson.ValidateAddressCmd) (out *btcjson.ValidateAddressChainResult, err error) {
	RPCHandlers["validateaddress"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan ValidateAddressRes):
		out, err = o.Res, o.Err
	}
	return
}

// VerifyChain calls the method with the given parameters
func (a API) VerifyChain(cmd *btcjson.VerifyChainCmd) (err error) {
	RPCHandlers["verifychain"].Call <- API{a.Ch, cmd, nil}
	return
}

// VerifyChainCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) VerifyChainCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan VerifyChainRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// VerifyChainGetRes returns a pointer to the value in the Result field
func (a API) VerifyChainGetRes() (out *bool, err error) {
	out, _ = a.Result.(*bool)
	err, _ = a.Result.(error)
	return
}

// VerifyChainWait calls the method and blocks until it returns or 5 seconds passes
func (a API) VerifyChainWait(cmd *btcjson.VerifyChainCmd) (out *bool, err error) {
	RPCHandlers["verifychain"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan VerifyChainRes):
		out, err = o.Res, o.Err
	}
	return
}

// VerifyMessage calls the method with the given parameters
func (a API) VerifyMessage(cmd *btcjson.VerifyMessageCmd) (err error) {
	RPCHandlers["verifymessage"].Call <- API{a.Ch, cmd, nil}
	return
}

// VerifyMessageCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) VerifyMessageCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan VerifyMessageRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// VerifyMessageGetRes returns a pointer to the value in the Result field
func (a API) VerifyMessageGetRes() (out *bool, err error) {
	out, _ = a.Result.(*bool)
	err, _ = a.Result.(error)
	return
}

// VerifyMessageWait calls the method and blocks until it returns or 5 seconds passes
func (a API) VerifyMessageWait(cmd *btcjson.VerifyMessageCmd) (out *bool, err error) {
	RPCHandlers["verifymessage"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan VerifyMessageRes):
		out, err = o.Res, o.Err
	}
	return
}

// Version calls the method with the given parameters
func (a API) Version(cmd *btcjson.VersionCmd) (err error) {
	RPCHandlers["version"].Call <- API{a.Ch, cmd, nil}
	return
}

// VersionCheck checks if a new message arrived on the result channel and 
// returns true if it does, as well as storing the value in the Result field
func (a API) VersionCheck() (isNew bool) {
	select {
	case o := <-a.Ch.(chan VersionRes):
		if o.Err != nil {
			a.Result = o.Err
		} else {
			a.Result = o.Res
		}
		isNew = true
	default:
	}
	return
}

// VersionGetRes returns a pointer to the value in the Result field
func (a API) VersionGetRes() (out *map[string]btcjson.VersionResult, err error) {
	out, _ = a.Result.(*map[string]btcjson.VersionResult)
	err, _ = a.Result.(error)
	return
}

// VersionWait calls the method and blocks until it returns or 5 seconds passes
func (a API) VersionWait(cmd *btcjson.VersionCmd) (out *map[string]btcjson.VersionResult, err error) {
	RPCHandlers["version"].Call <- API{a.Ch, cmd, nil}
	select {
	case <-time.After(time.Second * 5):
		break
	case o := <-a.Ch.(chan VersionRes):
		out, err = o.Res, o.Err
	}
	return
}

// RunAPI starts up the api handler server that receives rpc.API messages and runs the handler and returns the result
// Note that the parameters are type asserted to prevent the consumer of the API from sending wrong message types not
// because it's necessary since they are interfaces end to end
func RunAPI(server *Server, quit qu.C) {
	nrh := RPCHandlers
	go func() {
		Debug("starting up node cAPI")
		var err error
		var res interface{}
		for {
			select {
			case msg := <-nrh["addnode"].Call:
				if res, err = nrh["addnode"].
					Fn(server, msg.Params.(*btcjson.AddNodeCmd), nil); Check(err) {
				}
				if r, ok := res.(None); ok {
					msg.Ch.(chan AddNodeRes) <- AddNodeRes{&r, err}
				}
			case msg := <-nrh["createrawtransaction"].Call:
				if res, err = nrh["createrawtransaction"].
					Fn(server, msg.Params.(*btcjson.CreateRawTransactionCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan CreateRawTransactionRes) <- CreateRawTransactionRes{&r, err}
				}
			case msg := <-nrh["decoderawtransaction"].Call:
				if res, err = nrh["decoderawtransaction"].
					Fn(server, msg.Params.(*btcjson.DecodeRawTransactionCmd), nil); Check(err) {
				}
				if r, ok := res.(btcjson.TxRawDecodeResult); ok {
					msg.Ch.(chan DecodeRawTransactionRes) <- DecodeRawTransactionRes{&r, err}
				}
			case msg := <-nrh["decodescript"].Call:
				if res, err = nrh["decodescript"].
					Fn(server, msg.Params.(*btcjson.DecodeScriptCmd), nil); Check(err) {
				}
				if r, ok := res.(btcjson.DecodeScriptResult); ok {
					msg.Ch.(chan DecodeScriptRes) <- DecodeScriptRes{&r, err}
				}
			case msg := <-nrh["estimatefee"].Call:
				if res, err = nrh["estimatefee"].
					Fn(server, msg.Params.(*btcjson.EstimateFeeCmd), nil); Check(err) {
				}
				if r, ok := res.(float64); ok {
					msg.Ch.(chan EstimateFeeRes) <- EstimateFeeRes{&r, err}
				}
			case msg := <-nrh["generate"].Call:
				if res, err = nrh["generate"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.([]string); ok {
					msg.Ch.(chan GenerateRes) <- GenerateRes{&r, err}
				}
			case msg := <-nrh["getaddednodeinfo"].Call:
				if res, err = nrh["getaddednodeinfo"].
					Fn(server, msg.Params.(*btcjson.GetAddedNodeInfoCmd), nil); Check(err) {
				}
				if r, ok := res.([]btcjson.GetAddedNodeInfoResultAddr); ok {
					msg.Ch.(chan GetAddedNodeInfoRes) <- GetAddedNodeInfoRes{&r, err}
				}
			case msg := <-nrh["getbestblock"].Call:
				if res, err = nrh["getbestblock"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(btcjson.GetBestBlockResult); ok {
					msg.Ch.(chan GetBestBlockRes) <- GetBestBlockRes{&r, err}
				}
			case msg := <-nrh["getbestblockhash"].Call:
				if res, err = nrh["getbestblockhash"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan GetBestBlockHashRes) <- GetBestBlockHashRes{&r, err}
				}
			case msg := <-nrh["getblock"].Call:
				if res, err = nrh["getblock"].
					Fn(server, msg.Params.(*btcjson.GetBlockCmd), nil); Check(err) {
				}
				if r, ok := res.(btcjson.GetBlockVerboseResult); ok {
					msg.Ch.(chan GetBlockRes) <- GetBlockRes{&r, err}
				}
			case msg := <-nrh["getblockchaininfo"].Call:
				if res, err = nrh["getblockchaininfo"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(btcjson.GetBlockChainInfoResult); ok {
					msg.Ch.(chan GetBlockChainInfoRes) <- GetBlockChainInfoRes{&r, err}
				}
			case msg := <-nrh["getblockcount"].Call:
				if res, err = nrh["getblockcount"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(int64); ok {
					msg.Ch.(chan GetBlockCountRes) <- GetBlockCountRes{&r, err}
				}
			case msg := <-nrh["getblockhash"].Call:
				if res, err = nrh["getblockhash"].
					Fn(server, msg.Params.(*btcjson.GetBlockHashCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan GetBlockHashRes) <- GetBlockHashRes{&r, err}
				}
			case msg := <-nrh["getblockheader"].Call:
				if res, err = nrh["getblockheader"].
					Fn(server, msg.Params.(*btcjson.GetBlockHeaderCmd), nil); Check(err) {
				}
				if r, ok := res.(btcjson.GetBlockHeaderVerboseResult); ok {
					msg.Ch.(chan GetBlockHeaderRes) <- GetBlockHeaderRes{&r, err}
				}
			case msg := <-nrh["getblocktemplate"].Call:
				if res, err = nrh["getblocktemplate"].
					Fn(server, msg.Params.(*btcjson.GetBlockTemplateCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan GetBlockTemplateRes) <- GetBlockTemplateRes{&r, err}
				}
			case msg := <-nrh["getcfilter"].Call:
				if res, err = nrh["getcfilter"].
					Fn(server, msg.Params.(*btcjson.GetCFilterCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan GetCFilterRes) <- GetCFilterRes{&r, err}
				}
			case msg := <-nrh["getcfilterheader"].Call:
				if res, err = nrh["getcfilterheader"].
					Fn(server, msg.Params.(*btcjson.GetCFilterHeaderCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan GetCFilterHeaderRes) <- GetCFilterHeaderRes{&r, err}
				}
			case msg := <-nrh["getconnectioncount"].Call:
				if res, err = nrh["getconnectioncount"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(int32); ok {
					msg.Ch.(chan GetConnectionCountRes) <- GetConnectionCountRes{&r, err}
				}
			case msg := <-nrh["getcurrentnet"].Call:
				if res, err = nrh["getcurrentnet"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan GetCurrentNetRes) <- GetCurrentNetRes{&r, err}
				}
			case msg := <-nrh["getdifficulty"].Call:
				if res, err = nrh["getdifficulty"].
					Fn(server, msg.Params.(*btcjson.GetDifficultyCmd), nil); Check(err) {
				}
				if r, ok := res.(float64); ok {
					msg.Ch.(chan GetDifficultyRes) <- GetDifficultyRes{&r, err}
				}
			case msg := <-nrh["getgenerate"].Call:
				if res, err = nrh["getgenerate"].
					Fn(server, msg.Params.(*btcjson.GetHeadersCmd), nil); Check(err) {
				}
				if r, ok := res.(bool); ok {
					msg.Ch.(chan GetGenerateRes) <- GetGenerateRes{&r, err}
				}
			case msg := <-nrh["gethashespersec"].Call:
				if res, err = nrh["gethashespersec"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(float64); ok {
					msg.Ch.(chan GetHashesPerSecRes) <- GetHashesPerSecRes{&r, err}
				}
			case msg := <-nrh["getheaders"].Call:
				if res, err = nrh["getheaders"].
					Fn(server, msg.Params.(*btcjson.GetHeadersCmd), nil); Check(err) {
				}
				if r, ok := res.([]string); ok {
					msg.Ch.(chan GetHeadersRes) <- GetHeadersRes{&r, err}
				}
			case msg := <-nrh["getinfo"].Call:
				if res, err = nrh["getinfo"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(btcjson.InfoChainResult0); ok {
					msg.Ch.(chan GetInfoRes) <- GetInfoRes{&r, err}
				}
			case msg := <-nrh["getmempoolinfo"].Call:
				if res, err = nrh["getmempoolinfo"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(btcjson.GetMempoolInfoResult); ok {
					msg.Ch.(chan GetMempoolInfoRes) <- GetMempoolInfoRes{&r, err}
				}
			case msg := <-nrh["getmininginfo"].Call:
				if res, err = nrh["getmininginfo"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(btcjson.GetMiningInfoResult); ok {
					msg.Ch.(chan GetMiningInfoRes) <- GetMiningInfoRes{&r, err}
				}
			case msg := <-nrh["getnettotals"].Call:
				if res, err = nrh["getnettotals"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(btcjson.GetNetTotalsResult); ok {
					msg.Ch.(chan GetNetTotalsRes) <- GetNetTotalsRes{&r, err}
				}
			case msg := <-nrh["getnetworkhashps"].Call:
				if res, err = nrh["getnetworkhashps"].
					Fn(server, msg.Params.(*btcjson.GetNetworkHashPSCmd), nil); Check(err) {
				}
				if r, ok := res.([]btcjson.GetPeerInfoResult); ok {
					msg.Ch.(chan GetNetworkHashPSRes) <- GetNetworkHashPSRes{&r, err}
				}
			case msg := <-nrh["getpeerinfo"].Call:
				if res, err = nrh["getpeerinfo"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.([]btcjson.GetPeerInfoResult); ok {
					msg.Ch.(chan GetPeerInfoRes) <- GetPeerInfoRes{&r, err}
				}
			case msg := <-nrh["getrawmempool"].Call:
				if res, err = nrh["getrawmempool"].
					Fn(server, msg.Params.(*btcjson.GetRawMempoolCmd), nil); Check(err) {
				}
				if r, ok := res.([]string); ok {
					msg.Ch.(chan GetRawMempoolRes) <- GetRawMempoolRes{&r, err}
				}
			case msg := <-nrh["getrawtransaction"].Call:
				if res, err = nrh["getrawtransaction"].
					Fn(server, msg.Params.(*btcjson.GetRawTransactionCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan GetRawTransactionRes) <- GetRawTransactionRes{&r, err}
				}
			case msg := <-nrh["gettxout"].Call:
				if res, err = nrh["gettxout"].
					Fn(server, msg.Params.(*btcjson.GetTxOutCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan GetTxOutRes) <- GetTxOutRes{&r, err}
				}
			case msg := <-nrh["help"].Call:
				if res, err = nrh["help"].
					Fn(server, msg.Params.(*btcjson.HelpCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan HelpRes) <- HelpRes{&r, err}
				}
			case msg := <-nrh["node"].Call:
				if res, err = nrh["node"].
					Fn(server, msg.Params.(*btcjson.NodeCmd), nil); Check(err) {
				}
				if r, ok := res.(None); ok {
					msg.Ch.(chan NodeRes) <- NodeRes{&r, err}
				}
			case msg := <-nrh["ping"].Call:
				if res, err = nrh["ping"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(None); ok {
					msg.Ch.(chan PingRes) <- PingRes{&r, err}
				}
			case msg := <-nrh["resetchain"].Call:
				if res, err = nrh["resetchain"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(None); ok {
					msg.Ch.(chan ResetChainRes) <- ResetChainRes{&r, err}
				}
			case msg := <-nrh["restart"].Call:
				if res, err = nrh["restart"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(None); ok {
					msg.Ch.(chan RestartRes) <- RestartRes{&r, err}
				}
			case msg := <-nrh["searchrawtransactions"].Call:
				if res, err = nrh["searchrawtransactions"].
					Fn(server, msg.Params.(*btcjson.SearchRawTransactionsCmd), nil); Check(err) {
				}
				if r, ok := res.([]btcjson.SearchRawTransactionsResult); ok {
					msg.Ch.(chan SearchRawTransactionsRes) <- SearchRawTransactionsRes{&r, err}
				}
			case msg := <-nrh["sendrawtransaction"].Call:
				if res, err = nrh["sendrawtransaction"].
					Fn(server, msg.Params.(*btcjson.SendRawTransactionCmd), nil); Check(err) {
				}
				if r, ok := res.(None); ok {
					msg.Ch.(chan SendRawTransactionRes) <- SendRawTransactionRes{&r, err}
				}
			case msg := <-nrh["setgenerate"].Call:
				if res, err = nrh["setgenerate"].
					Fn(server, msg.Params.(*btcjson.SetGenerateCmd), nil); Check(err) {
				}
				if r, ok := res.(None); ok {
					msg.Ch.(chan SetGenerateRes) <- SetGenerateRes{&r, err}
				}
			case msg := <-nrh["stop"].Call:
				if res, err = nrh["stop"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(None); ok {
					msg.Ch.(chan StopRes) <- StopRes{&r, err}
				}
			case msg := <-nrh["submitblock"].Call:
				if res, err = nrh["submitblock"].
					Fn(server, msg.Params.(*btcjson.SubmitBlockCmd), nil); Check(err) {
				}
				if r, ok := res.(string); ok {
					msg.Ch.(chan SubmitBlockRes) <- SubmitBlockRes{&r, err}
				}
			case msg := <-nrh["uptime"].Call:
				if res, err = nrh["uptime"].
					Fn(server, msg.Params.(*None), nil); Check(err) {
				}
				if r, ok := res.(btcjson.GetMempoolInfoResult); ok {
					msg.Ch.(chan UptimeRes) <- UptimeRes{&r, err}
				}
			case msg := <-nrh["validateaddress"].Call:
				if res, err = nrh["validateaddress"].
					Fn(server, msg.Params.(*btcjson.ValidateAddressCmd), nil); Check(err) {
				}
				if r, ok := res.(btcjson.ValidateAddressChainResult); ok {
					msg.Ch.(chan ValidateAddressRes) <- ValidateAddressRes{&r, err}
				}
			case msg := <-nrh["verifychain"].Call:
				if res, err = nrh["verifychain"].
					Fn(server, msg.Params.(*btcjson.VerifyChainCmd), nil); Check(err) {
				}
				if r, ok := res.(bool); ok {
					msg.Ch.(chan VerifyChainRes) <- VerifyChainRes{&r, err}
				}
			case msg := <-nrh["verifymessage"].Call:
				if res, err = nrh["verifymessage"].
					Fn(server, msg.Params.(*btcjson.VerifyMessageCmd), nil); Check(err) {
				}
				if r, ok := res.(bool); ok {
					msg.Ch.(chan VerifyMessageRes) <- VerifyMessageRes{&r, err}
				}
			case msg := <-nrh["version"].Call:
				if res, err = nrh["version"].
					Fn(server, msg.Params.(*btcjson.VersionCmd), nil); Check(err) {
				}
				if r, ok := res.(map[string]btcjson.VersionResult); ok {
					msg.Ch.(chan VersionRes) <- VersionRes{&r, err}
				}
			case <-quit:
				Debug("stopping wallet cAPI")
				return
			}
		}
	}()
}

// RPC API functions to use with net/rpc

func (c *CAPI) AddNode(req *btcjson.AddNodeCmd, resp None) (err error) {
	nrh := RPCHandlers
	res := nrh["addnode"].Result()
	res.Params = req
	nrh["addnode"].Call <- res
	select {
	case resp = <-res.Ch.(chan None):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) CreateRawTransaction(req *btcjson.CreateRawTransactionCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["createrawtransaction"].Result()
	res.Params = req
	nrh["createrawtransaction"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) DecodeRawTransaction(req *btcjson.DecodeRawTransactionCmd, resp btcjson.TxRawDecodeResult) (err error) {
	nrh := RPCHandlers
	res := nrh["decoderawtransaction"].Result()
	res.Params = req
	nrh["decoderawtransaction"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.TxRawDecodeResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) DecodeScript(req *btcjson.DecodeScriptCmd, resp btcjson.DecodeScriptResult) (err error) {
	nrh := RPCHandlers
	res := nrh["decodescript"].Result()
	res.Params = req
	nrh["decodescript"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.DecodeScriptResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) EstimateFee(req *btcjson.EstimateFeeCmd, resp float64) (err error) {
	nrh := RPCHandlers
	res := nrh["estimatefee"].Result()
	res.Params = req
	nrh["estimatefee"].Call <- res
	select {
	case resp = <-res.Ch.(chan float64):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) Generate(req *None, resp []string) (err error) {
	nrh := RPCHandlers
	res := nrh["generate"].Result()
	res.Params = req
	nrh["generate"].Call <- res
	select {
	case resp = <-res.Ch.(chan []string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetAddedNodeInfo(req *btcjson.GetAddedNodeInfoCmd, resp []btcjson.GetAddedNodeInfoResultAddr) (err error) {
	nrh := RPCHandlers
	res := nrh["getaddednodeinfo"].Result()
	res.Params = req
	nrh["getaddednodeinfo"].Call <- res
	select {
	case resp = <-res.Ch.(chan []btcjson.GetAddedNodeInfoResultAddr):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetBestBlock(req *None, resp btcjson.GetBestBlockResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getbestblock"].Result()
	res.Params = req
	nrh["getbestblock"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.GetBestBlockResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetBestBlockHash(req *None, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["getbestblockhash"].Result()
	res.Params = req
	nrh["getbestblockhash"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetBlock(req *btcjson.GetBlockCmd, resp btcjson.GetBlockVerboseResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getblock"].Result()
	res.Params = req
	nrh["getblock"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.GetBlockVerboseResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetBlockChainInfo(req *None, resp btcjson.GetBlockChainInfoResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getblockchaininfo"].Result()
	res.Params = req
	nrh["getblockchaininfo"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.GetBlockChainInfoResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetBlockCount(req *None, resp int64) (err error) {
	nrh := RPCHandlers
	res := nrh["getblockcount"].Result()
	res.Params = req
	nrh["getblockcount"].Call <- res
	select {
	case resp = <-res.Ch.(chan int64):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetBlockHash(req *btcjson.GetBlockHashCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["getblockhash"].Result()
	res.Params = req
	nrh["getblockhash"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetBlockHeader(req *btcjson.GetBlockHeaderCmd, resp btcjson.GetBlockHeaderVerboseResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getblockheader"].Result()
	res.Params = req
	nrh["getblockheader"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.GetBlockHeaderVerboseResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetBlockTemplate(req *btcjson.GetBlockTemplateCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["getblocktemplate"].Result()
	res.Params = req
	nrh["getblocktemplate"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetCFilter(req *btcjson.GetCFilterCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["getcfilter"].Result()
	res.Params = req
	nrh["getcfilter"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetCFilterHeader(req *btcjson.GetCFilterHeaderCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["getcfilterheader"].Result()
	res.Params = req
	nrh["getcfilterheader"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetConnectionCount(req *None, resp int32) (err error) {
	nrh := RPCHandlers
	res := nrh["getconnectioncount"].Result()
	res.Params = req
	nrh["getconnectioncount"].Call <- res
	select {
	case resp = <-res.Ch.(chan int32):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetCurrentNet(req *None, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["getcurrentnet"].Result()
	res.Params = req
	nrh["getcurrentnet"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetDifficulty(req *btcjson.GetDifficultyCmd, resp float64) (err error) {
	nrh := RPCHandlers
	res := nrh["getdifficulty"].Result()
	res.Params = req
	nrh["getdifficulty"].Call <- res
	select {
	case resp = <-res.Ch.(chan float64):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetGenerate(req *btcjson.GetHeadersCmd, resp bool) (err error) {
	nrh := RPCHandlers
	res := nrh["getgenerate"].Result()
	res.Params = req
	nrh["getgenerate"].Call <- res
	select {
	case resp = <-res.Ch.(chan bool):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetHashesPerSec(req *None, resp float64) (err error) {
	nrh := RPCHandlers
	res := nrh["gethashespersec"].Result()
	res.Params = req
	nrh["gethashespersec"].Call <- res
	select {
	case resp = <-res.Ch.(chan float64):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetHeaders(req *btcjson.GetHeadersCmd, resp []string) (err error) {
	nrh := RPCHandlers
	res := nrh["getheaders"].Result()
	res.Params = req
	nrh["getheaders"].Call <- res
	select {
	case resp = <-res.Ch.(chan []string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetInfo(req *None, resp btcjson.InfoChainResult0) (err error) {
	nrh := RPCHandlers
	res := nrh["getinfo"].Result()
	res.Params = req
	nrh["getinfo"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.InfoChainResult0):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetMempoolInfo(req *None, resp btcjson.GetMempoolInfoResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getmempoolinfo"].Result()
	res.Params = req
	nrh["getmempoolinfo"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.GetMempoolInfoResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetMiningInfo(req *None, resp btcjson.GetMiningInfoResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getmininginfo"].Result()
	res.Params = req
	nrh["getmininginfo"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.GetMiningInfoResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetNetTotals(req *None, resp btcjson.GetNetTotalsResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getnettotals"].Result()
	res.Params = req
	nrh["getnettotals"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.GetNetTotalsResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetNetworkHashPS(req *btcjson.GetNetworkHashPSCmd, resp []btcjson.GetPeerInfoResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getnetworkhashps"].Result()
	res.Params = req
	nrh["getnetworkhashps"].Call <- res
	select {
	case resp = <-res.Ch.(chan []btcjson.GetPeerInfoResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetPeerInfo(req *None, resp []btcjson.GetPeerInfoResult) (err error) {
	nrh := RPCHandlers
	res := nrh["getpeerinfo"].Result()
	res.Params = req
	nrh["getpeerinfo"].Call <- res
	select {
	case resp = <-res.Ch.(chan []btcjson.GetPeerInfoResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetRawMempool(req *btcjson.GetRawMempoolCmd, resp []string) (err error) {
	nrh := RPCHandlers
	res := nrh["getrawmempool"].Result()
	res.Params = req
	nrh["getrawmempool"].Call <- res
	select {
	case resp = <-res.Ch.(chan []string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetRawTransaction(req *btcjson.GetRawTransactionCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["getrawtransaction"].Result()
	res.Params = req
	nrh["getrawtransaction"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) GetTxOut(req *btcjson.GetTxOutCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["gettxout"].Result()
	res.Params = req
	nrh["gettxout"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) Help(req *btcjson.HelpCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["help"].Result()
	res.Params = req
	nrh["help"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) Node(req *btcjson.NodeCmd, resp None) (err error) {
	nrh := RPCHandlers
	res := nrh["node"].Result()
	res.Params = req
	nrh["node"].Call <- res
	select {
	case resp = <-res.Ch.(chan None):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) Ping(req *None, resp None) (err error) {
	nrh := RPCHandlers
	res := nrh["ping"].Result()
	res.Params = req
	nrh["ping"].Call <- res
	select {
	case resp = <-res.Ch.(chan None):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) ResetChain(req *None, resp None) (err error) {
	nrh := RPCHandlers
	res := nrh["resetchain"].Result()
	res.Params = req
	nrh["resetchain"].Call <- res
	select {
	case resp = <-res.Ch.(chan None):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) Restart(req *None, resp None) (err error) {
	nrh := RPCHandlers
	res := nrh["restart"].Result()
	res.Params = req
	nrh["restart"].Call <- res
	select {
	case resp = <-res.Ch.(chan None):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) SearchRawTransactions(req *btcjson.SearchRawTransactionsCmd, resp []btcjson.SearchRawTransactionsResult) (err error) {
	nrh := RPCHandlers
	res := nrh["searchrawtransactions"].Result()
	res.Params = req
	nrh["searchrawtransactions"].Call <- res
	select {
	case resp = <-res.Ch.(chan []btcjson.SearchRawTransactionsResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) SendRawTransaction(req *btcjson.SendRawTransactionCmd, resp None) (err error) {
	nrh := RPCHandlers
	res := nrh["sendrawtransaction"].Result()
	res.Params = req
	nrh["sendrawtransaction"].Call <- res
	select {
	case resp = <-res.Ch.(chan None):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) SetGenerate(req *btcjson.SetGenerateCmd, resp None) (err error) {
	nrh := RPCHandlers
	res := nrh["setgenerate"].Result()
	res.Params = req
	nrh["setgenerate"].Call <- res
	select {
	case resp = <-res.Ch.(chan None):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) Stop(req *None, resp None) (err error) {
	nrh := RPCHandlers
	res := nrh["stop"].Result()
	res.Params = req
	nrh["stop"].Call <- res
	select {
	case resp = <-res.Ch.(chan None):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) SubmitBlock(req *btcjson.SubmitBlockCmd, resp string) (err error) {
	nrh := RPCHandlers
	res := nrh["submitblock"].Result()
	res.Params = req
	nrh["submitblock"].Call <- res
	select {
	case resp = <-res.Ch.(chan string):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) Uptime(req *None, resp btcjson.GetMempoolInfoResult) (err error) {
	nrh := RPCHandlers
	res := nrh["uptime"].Result()
	res.Params = req
	nrh["uptime"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.GetMempoolInfoResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) ValidateAddress(req *btcjson.ValidateAddressCmd, resp btcjson.ValidateAddressChainResult) (err error) {
	nrh := RPCHandlers
	res := nrh["validateaddress"].Result()
	res.Params = req
	nrh["validateaddress"].Call <- res
	select {
	case resp = <-res.Ch.(chan btcjson.ValidateAddressChainResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) VerifyChain(req *btcjson.VerifyChainCmd, resp bool) (err error) {
	nrh := RPCHandlers
	res := nrh["verifychain"].Result()
	res.Params = req
	nrh["verifychain"].Call <- res
	select {
	case resp = <-res.Ch.(chan bool):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) VerifyMessage(req *btcjson.VerifyMessageCmd, resp bool) (err error) {
	nrh := RPCHandlers
	res := nrh["verifymessage"].Result()
	res.Params = req
	nrh["verifymessage"].Call <- res
	select {
	case resp = <-res.Ch.(chan bool):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

func (c *CAPI) Version(req *btcjson.VersionCmd, resp map[string]btcjson.VersionResult) (err error) {
	nrh := RPCHandlers
	res := nrh["version"].Result()
	res.Params = req
	nrh["version"].Call <- res
	select {
	case resp = <-res.Ch.(chan map[string]btcjson.VersionResult):
	case <-time.After(c.Timeout):
	case <-c.quit:
	}
	return
}

// Client call wrappers for a CAPI client with a given Conn

func (r *CAPIClient) AddNode(cmd ...*btcjson.AddNodeCmd) (res None, err error) {
	var c *btcjson.AddNodeCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.AddNode", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) CreateRawTransaction(cmd ...*btcjson.CreateRawTransactionCmd) (res string, err error) {
	var c *btcjson.CreateRawTransactionCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.CreateRawTransaction", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) DecodeRawTransaction(cmd ...*btcjson.DecodeRawTransactionCmd) (res btcjson.TxRawDecodeResult, err error) {
	var c *btcjson.DecodeRawTransactionCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.DecodeRawTransaction", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) DecodeScript(cmd ...*btcjson.DecodeScriptCmd) (res btcjson.DecodeScriptResult, err error) {
	var c *btcjson.DecodeScriptCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.DecodeScript", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) EstimateFee(cmd ...*btcjson.EstimateFeeCmd) (res float64, err error) {
	var c *btcjson.EstimateFeeCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.EstimateFee", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) Generate(cmd ...*None) (res []string, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.Generate", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetAddedNodeInfo(cmd ...*btcjson.GetAddedNodeInfoCmd) (res []btcjson.GetAddedNodeInfoResultAddr, err error) {
	var c *btcjson.GetAddedNodeInfoCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetAddedNodeInfo", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetBestBlock(cmd ...*None) (res btcjson.GetBestBlockResult, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetBestBlock", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetBestBlockHash(cmd ...*None) (res string, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetBestBlockHash", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetBlock(cmd ...*btcjson.GetBlockCmd) (res btcjson.GetBlockVerboseResult, err error) {
	var c *btcjson.GetBlockCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetBlock", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetBlockChainInfo(cmd ...*None) (res btcjson.GetBlockChainInfoResult, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetBlockChainInfo", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetBlockCount(cmd ...*None) (res int64, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetBlockCount", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetBlockHash(cmd ...*btcjson.GetBlockHashCmd) (res string, err error) {
	var c *btcjson.GetBlockHashCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetBlockHash", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetBlockHeader(cmd ...*btcjson.GetBlockHeaderCmd) (res btcjson.GetBlockHeaderVerboseResult, err error) {
	var c *btcjson.GetBlockHeaderCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetBlockHeader", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetBlockTemplate(cmd ...*btcjson.GetBlockTemplateCmd) (res string, err error) {
	var c *btcjson.GetBlockTemplateCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetBlockTemplate", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetCFilter(cmd ...*btcjson.GetCFilterCmd) (res string, err error) {
	var c *btcjson.GetCFilterCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetCFilter", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetCFilterHeader(cmd ...*btcjson.GetCFilterHeaderCmd) (res string, err error) {
	var c *btcjson.GetCFilterHeaderCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetCFilterHeader", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetConnectionCount(cmd ...*None) (res int32, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetConnectionCount", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetCurrentNet(cmd ...*None) (res string, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetCurrentNet", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetDifficulty(cmd ...*btcjson.GetDifficultyCmd) (res float64, err error) {
	var c *btcjson.GetDifficultyCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetDifficulty", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetGenerate(cmd ...*btcjson.GetHeadersCmd) (res bool, err error) {
	var c *btcjson.GetHeadersCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetGenerate", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetHashesPerSec(cmd ...*None) (res float64, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetHashesPerSec", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetHeaders(cmd ...*btcjson.GetHeadersCmd) (res []string, err error) {
	var c *btcjson.GetHeadersCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetHeaders", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetInfo(cmd ...*None) (res btcjson.InfoChainResult0, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetInfo", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetMempoolInfo(cmd ...*None) (res btcjson.GetMempoolInfoResult, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetMempoolInfo", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetMiningInfo(cmd ...*None) (res btcjson.GetMiningInfoResult, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetMiningInfo", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetNetTotals(cmd ...*None) (res btcjson.GetNetTotalsResult, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetNetTotals", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetNetworkHashPS(cmd ...*btcjson.GetNetworkHashPSCmd) (res []btcjson.GetPeerInfoResult, err error) {
	var c *btcjson.GetNetworkHashPSCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetNetworkHashPS", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetPeerInfo(cmd ...*None) (res []btcjson.GetPeerInfoResult, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetPeerInfo", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetRawMempool(cmd ...*btcjson.GetRawMempoolCmd) (res []string, err error) {
	var c *btcjson.GetRawMempoolCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetRawMempool", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetRawTransaction(cmd ...*btcjson.GetRawTransactionCmd) (res string, err error) {
	var c *btcjson.GetRawTransactionCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetRawTransaction", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) GetTxOut(cmd ...*btcjson.GetTxOutCmd) (res string, err error) {
	var c *btcjson.GetTxOutCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.GetTxOut", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) Help(cmd ...*btcjson.HelpCmd) (res string, err error) {
	var c *btcjson.HelpCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.Help", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) Node(cmd ...*btcjson.NodeCmd) (res None, err error) {
	var c *btcjson.NodeCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.Node", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) Ping(cmd ...*None) (res None, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.Ping", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) ResetChain(cmd ...*None) (res None, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.ResetChain", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) Restart(cmd ...*None) (res None, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.Restart", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) SearchRawTransactions(cmd ...*btcjson.SearchRawTransactionsCmd) (res []btcjson.SearchRawTransactionsResult, err error) {
	var c *btcjson.SearchRawTransactionsCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.SearchRawTransactions", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) SendRawTransaction(cmd ...*btcjson.SendRawTransactionCmd) (res None, err error) {
	var c *btcjson.SendRawTransactionCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.SendRawTransaction", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) SetGenerate(cmd ...*btcjson.SetGenerateCmd) (res None, err error) {
	var c *btcjson.SetGenerateCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.SetGenerate", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) Stop(cmd ...*None) (res None, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.Stop", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) SubmitBlock(cmd ...*btcjson.SubmitBlockCmd) (res string, err error) {
	var c *btcjson.SubmitBlockCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.SubmitBlock", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) Uptime(cmd ...*None) (res btcjson.GetMempoolInfoResult, err error) {
	var c *None
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.Uptime", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) ValidateAddress(cmd ...*btcjson.ValidateAddressCmd) (res btcjson.ValidateAddressChainResult, err error) {
	var c *btcjson.ValidateAddressCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.ValidateAddress", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) VerifyChain(cmd ...*btcjson.VerifyChainCmd) (res bool, err error) {
	var c *btcjson.VerifyChainCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.VerifyChain", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) VerifyMessage(cmd ...*btcjson.VerifyMessageCmd) (res bool, err error) {
	var c *btcjson.VerifyMessageCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.VerifyMessage", c, &res); Check(err) {
	}
	return
}

func (r *CAPIClient) Version(cmd ...*btcjson.VersionCmd) (res map[string]btcjson.VersionResult, err error) {
	var c *btcjson.VersionCmd
	if len(cmd) > 0 {
		c = cmd[0]
	}
	if err = r.Call("CAPI.Version", c, &res); Check(err) {
	}
	return
}
