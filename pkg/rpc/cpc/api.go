package api

import (
	"github.com/l0k18/pod/pkg/rpc/btcjson"
	"github.com/l0k18/pod/pkg/rpc/chainrpc"
	qu "github.com/l0k18/pod/pkg/util/quit"
)

// StartAPI starts up the api handler server that receives rpc.API messages and runs the handler and returns the result
// Note that the parameters are type asserted to prevent the consumer of the API from sending wrong message types not
// because it's necessary since they are interfaces end to end
func StartAPI(server *chainrpc.Server, quit qu.C) {
	nrh := chainrpc.RPCHandlers
	go func() {
		var err error
		var res interface{}
		for {
			select {
			case msg := <-nrh["addnode"].Call:
				if _, err = nrh["addnode"].
					Fn(
						server, msg.Params.(btcjson.AddNodeCmd),
						nil,
				); Check(err) {
				}
				msg.Ch.(chan chainrpc.AddNodeRes) <- chainrpc.AddNodeRes{
					Res: nil, Err: err,
				}
			case msg := <-nrh["createrawtransaction"].Call:
				if res, err = nrh["createrawtransaction"].
					Fn(server, msg.Params.(btcjson.CreateRawTransactionCmd), nil); Check(err) {
				}
				msg.Ch.(chan chainrpc.CreateRawTransactionRes) <- chainrpc.CreateRawTransactionRes{
					Res: res.(*string), Err: err,
				}
			case msg := <-nrh["decoderawtransaction"].Call:
				var ret btcjson.TxRawDecodeResult
				if res, err = nrh["decoderawtransaction"].Fn(
					server, msg.Params.(btcjson.DecodeRawTransactionCmd),
					nil,
				); Check(err) {
				} else {
					ret = res.(btcjson.TxRawDecodeResult)
				}
				msg.Ch.(chan chainrpc.DecodeRawTransactionRes) <- chainrpc.DecodeRawTransactionRes{
					Res: &ret, Err: err,
				}
			case msg := <-nrh["decodescript"].Call:
				if res, err = nrh["decodescript"].Fn(server, msg.Params.(btcjson.DecodeScriptCmd), nil); Check(err) {
				}
				msg.Ch.(chan chainrpc.DecodeScriptRes) <- chainrpc.DecodeScriptRes{
					Res: res.(*btcjson.DecodeScriptResult), Err: err,
				}
			case msg := <-nrh["estimatefee"].Call:
				if res, err = nrh["estimatefee"].
					Fn(
						server, msg.Params.(btcjson.EstimateFeeCmd),
						nil,
				); Check(err) {
				}
				msg.Ch.(chan chainrpc.EstimateFeeRes) <- chainrpc.EstimateFeeRes{
					Res: res.(*float64), Err: err,
				}
			case <-quit:
				return
			}
		}
	}()
}
