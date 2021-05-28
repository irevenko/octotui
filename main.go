package main

import (
	"fmt"
	"os"

	cmd "github.com/irevenko/octotui/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
