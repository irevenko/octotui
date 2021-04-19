package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	h "github.com/irevenko/octotui/helpers"
)

var RootCmd = &cobra.Command{
	Use:   "octotui",
	Short: "GitHub stats in your terminal",
	Long:  `Complete documentation is available at https://github.com/irevenko/octotui`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(h.LoadOwner())
	},
}
