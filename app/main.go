// Package app is a multi-function universal binary that does all the things.
//
// Parallelcoin Pod
//
// This is the heart of configuration and coordination of the parts that compose the parallelcoin Pod - Ctl, Node and
// Wallet, and the extended, combined Shell and the Gio GUI.
package app

import (
	"os"
	
	"github.com/l0k18/pod/app/conte"
)

const (
	Name              = "pod"
	confExt           = ".json"
	appLanguage       = "en"
	podConfigFilename = Name + confExt
	PARSER            = "json"
)

// Main is the entrypoint for the pod AiO suite
func Main() int {
	cx := conte.GetNewContext(Name, appLanguage, "main")
	cx.App = GetApp(cx)
	if e := cx.App.Run(os.Args); Check(e) {
		return 1
	}
	return 0
}
