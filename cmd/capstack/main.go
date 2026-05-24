package main

import (
	"fmt"
	"os"

	"github.com/grokify/prism-capability/cli"
)

func main() {
	// Override the command name for standalone use
	cli.RootCmd.Use = "capstack"
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
