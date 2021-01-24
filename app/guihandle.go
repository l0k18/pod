package app

import (
	"github.com/urfave/cli"
	
	"github.com/l0k18/pod/app/config"
	"github.com/l0k18/pod/app/conte"
	"github.com/l0k18/pod/cmd/gui"
)

func walletGUIHandle(cx *conte.Xt) func(c *cli.Context) (err error) {
	return func(c *cli.Context) (err error) {
		// Debug(os.Args)
		config.Configure(cx, c.Command.Name, true)
		// interrupt.AddHandler(func() {
		// 	Debug("wallet gui is shut down")
		// })
		if err := gui.Main(cx, c); Check(err) {
		}
		Debug("pod gui finished")
		return
	}
}
