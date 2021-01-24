package app

import (
	"fmt"
	"io/ioutil"
	"os"
	
	"github.com/urfave/cli"
	
	"github.com/l0k18/pod/app/config"
	
	"github.com/l0k18/pod/app/apputil"
	"github.com/l0k18/pod/app/conte"
	"github.com/l0k18/pod/cmd/node"
	"github.com/l0k18/pod/cmd/walletmain"
	"github.com/l0k18/pod/pkg/wallet"
)

func ShellHandle(cx *conte.Xt) func(c *cli.Context) (err error) {
	return func(c *cli.Context) (err error) {
		config.Configure(cx, c.Command.Name, true)
		Debug("starting shell")
		if *cx.Config.TLS || *cx.Config.ServerTLS {
			// generate the tls certificate if configured
			if apputil.FileExists(*cx.Config.RPCCert) && apputil.FileExists(*cx.Config.RPCKey) &&
				apputil.FileExists(*cx.Config.CAFile) {
				
			} else {
				_, _ = walletmain.GenerateRPCKeyPair(cx.Config, true)
			}
		}
		dbFilename :=
			*cx.Config.DataDir + slash +
				cx.ActiveNet.Params.Name + slash +
				wallet.DbName
		if !apputil.FileExists(dbFilename) && !cx.IsGUI {
			// log.SetLevel("off", false)
			if err := walletmain.CreateWallet(cx.ActiveNet, cx.Config); err != nil {
				Error("failed to create wallet", err)
			}
			fmt.Println("restart to complete initial setup")
			os.Exit(1)
		}
		// for security with apps launching the wallet, the public password can be set with a file that is deleted after
		walletPassPath := *cx.Config.DataDir + slash + cx.ActiveNet.Params.Name + slash + "wp.txt"
		Debug("reading password from", walletPassPath)
		if apputil.FileExists(walletPassPath) {
			var b []byte
			if b, err = ioutil.ReadFile(walletPassPath); !Check(err) {
				*cx.Config.WalletPass = string(b)
				Debug("read password '" + string(b) + "'")
				for i := range b {
					b[i] = 0
				}
				if err = ioutil.WriteFile(walletPassPath, b, 0700); Check(err) {
				}
				if err = os.Remove(walletPassPath); Check(err) {
				}
				Debug("wallet cookie deleted", *cx.Config.WalletPass)
			}
		}
		if !*cx.Config.NodeOff {
			go func() {
				err = node.Main(cx)
				if err != nil {
					Error("error starting node ", err)
				}
			}()
			Info("starting node")
			if !*cx.Config.DisableRPC {
				cx.RPCServer = <-cx.NodeChan
			}
			Info("node started")
		}
		if !*cx.Config.WalletOff {
			go func() {
				err = walletmain.Main(cx)
				if err != nil {
					fmt.Println("error running wallet:", err)
				}
			}()
			Info("starting wallet")
			if !*cx.Config.DisableRPC {
				cx.WalletServer = <-cx.WalletChan
			}
			Info("wallet started")
		}
		Debug("shell started")
		// cx.WaitGroup.Wait()
		cx.WaitWait()
		return nil
	}
}
