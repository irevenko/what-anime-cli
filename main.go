package main

import (
	"fmt"
	"os"

	"github.com/irevenko/what-anime-cli/cli"
)

func main() {
	cli.AddCommands()

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
