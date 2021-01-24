package app

import (
	"os"
	
	"github.com/l0k18/pod/pkg/util/interrupt"
	
	"github.com/l0k18/pod/app/config"
	
	"github.com/urfave/cli"
	
	"github.com/l0k18/pod/cmd/kopach"
	"github.com/l0k18/pod/pkg/chain/config/netparams"
	"github.com/l0k18/pod/pkg/chain/fork"
	
	"github.com/l0k18/pod/app/conte"
)

func KopachHandle(cx *conte.Xt) func(c *cli.Context) (err error) {
	return func(c *cli.Context) (err error) {
		Info("starting up kopach standalone miner for parallelcoin")
		Debug(os.Args)
		config.Configure(cx, c.Command.Name, true)
		if cx.ActiveNet.Name == netparams.TestNet3Params.Name {
			fork.IsTestnet = true
		}
		// quit := qu.T()
		// interrupt.AddHandler(func() {
		// 	Debug("Handle interrupt")
		defer cx.KillAll.Q()
		// os.Exit(0)
		// })
		err = kopach.Handle(cx)(c)
		// Debug(interrupt.GoroutineDump())
		<-interrupt.HandlersDone
		Debug("kopach main finished")
		return
	}
}
