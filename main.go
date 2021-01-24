// Package main is the root of the Parallelcoin Pod software suite
//
// It slices, it dices
//
package main

import (
	"fmt"
	_ "net/http/pprof"
	
	_ "github.com/l0k18/pod/pkg"
	"github.com/l0k18/pod/version"
	
	"github.com/l0k18/pod/cmd"
)

func main() {
	version.URL = URL
	version.GitRef = GitRef
	version.GitCommit = GitCommit
	version.BuildTime = BuildTime
	version.Tag = Tag
	version.Get = GetVersion
	cmd.Main()
}

var (
	URL       string
	GitRef    string
	GitCommit string
	BuildTime string
	Tag       string
)

func GetVersion() string {
	return fmt.Sprintf(
		`ParallelCoin Pod
	repo: %s
	branch: %s
	commit: %s
	built: %s
	tag: %s
`,
		URL, GitRef, GitCommit, BuildTime, Tag)
}
