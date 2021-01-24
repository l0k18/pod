// +build headless

package app

import (
	"os"
	
	"github.com/urfave/cli"
	
	"github.com/l0k18/pod/app/conte"
)

var walletGUIHandle = func(cx *conte.Xt) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		Warn("GUI was disabled for this build (server only version)")
		os.Exit(1)
		return nil
	}
}
