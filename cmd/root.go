package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "octotui",
	Short: "GitHub stats in your terminal",
	Long:  `Complete documentation is available at https://github.com/irevenko/octotui`,
}
