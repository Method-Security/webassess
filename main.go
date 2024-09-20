package main

import (
	"flag"
	"os"

	"github.com/Method-Security/webassess/cmd"
)

var version = "none"

func main() {
	flag.Parse()

	webassess := cmd.NewWebAssess(version)
	webassess.InitRootCommand()
	webassess.InitURLAssess()

	if err := webassess.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
