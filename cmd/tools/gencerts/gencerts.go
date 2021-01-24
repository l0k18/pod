package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
	
	"github.com/jessevdk/go-flags"
	
	"github.com/l0k18/pod/app/appdata"
	"github.com/l0k18/pod/pkg/util"
)

type config struct {
	Directory    string   `short:"d" long:"directory" description:"Directory to write certificate pair"`
	Years        int      `short:"y" long:"years" description:"How many years a certificate is valid for"`
	Organization string   `short:"o" long:"org" description:"Organization in certificate"`
	ExtraHosts   []string `short:"H" long:"host" description:"Additional hosts/IPs to create certificate for"`
	Force        bool     `short:"f" long:"force" description:"Force overwriting of any old certs and keys"`
}

func main() {
	cfg := config{
		Years:        10,
		Organization: "gencerts",
	}
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		Error(err)
		if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			parser.WriteHelp(os.Stderr)
		}
		return
	}
	if cfg.Directory == "" {
		var err error
		cfg.Directory, err = os.Getwd()
		if err != nil {
			Error(err)
			_, _ = fmt.Fprintf(os.Stderr, "no directory specified and cannot get working directory\n")
			os.Exit(1)
		}
	}
	cfg.Directory = cleanAndExpandPath(cfg.Directory)
	certFile := filepath.Join(cfg.Directory, "rpc.cert")
	caFile := filepath.Join(cfg.Directory, "ca.cert")
	keyFile := filepath.Join(cfg.Directory, "rpc.key")
	if !cfg.Force {
		if fileExists(certFile) || fileExists(keyFile) {
			_, _ = fmt.Fprintf(os.Stderr, "%v: certificate and/or key files exist; use -f to force\n", cfg.Directory)
			os.Exit(1)
		}
	}
	validUntil := time.Now().Add(time.Duration(cfg.Years) * 365 * 24 * time.Hour)
	var cert, key []byte
	cert, key, err = util.NewTLSCertPair(cfg.Organization, validUntil, cfg.ExtraHosts)
	if err != nil {
		Error(err)
		_, _ = fmt.Fprintf(os.Stderr, "cannot generate certificate pair: %v\n", err)
		os.Exit(1)
	}
	// Write cert and key files.
	if err = ioutil.WriteFile(certFile, cert, 0666); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot write cert: %v\n", err)
		os.Exit(1)
	}
	// Write cert and key files.
	if err = ioutil.WriteFile(caFile, cert, 0666); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot write ca cert: %v\n", err)
		os.Exit(1)
	}
	if err = ioutil.WriteFile(keyFile, key, 0600); err != nil {
		if err := os.Remove(certFile); Check(err) {
		}
		_, _ = fmt.Fprintf(os.Stderr, "cannot write key: %v\n", err)
		os.Exit(1)
	}
}

// cleanAndExpandPath expands environement variables and leading ~ in the passed path, cleans the result, and returns
// it.
func cleanAndExpandPath(
	path string) string {
	// Expand initial ~ to OS specific home directory.
	if strings.HasPrefix(path, "~") {
		appHomeDir := appdata.Dir("gencerts", false)
		homeDir := filepath.Dir(appHomeDir)
		path = strings.Replace(path, "~", homeDir, 1)
	}
	// NOTE: The os.ExpandEnv doesn't work with Windows-style %VARIABLE%, but they variables can still be expanded via
	// POSIX-style $VARIABLE.
	return filepath.Clean(os.ExpandEnv(path))
}

// filesExists reports whether the named file or directory exists.
func fileExists(
	name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
