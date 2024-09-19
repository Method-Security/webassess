package main

import (
	"flag"
	"os"

	"github.com/Method-Security/aiassess/cmd"
)

var version = "none"

func main() {
	flag.Parse()

	aiassess := cmd.NewAIAssess(version)
	aiassess.InitRootCommand()
	aiassess.InitURLAssess()

	if err := aiassess.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
