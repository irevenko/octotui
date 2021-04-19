package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"

	h "github.com/irevenko/octotui/helpers"
)

var RootCmd = &cobra.Command{
	Use:   "octotui",
	Short: "GitHub stats in your terminal",
	Long:  `Complete documentation is available at https://github.com/irevenko/octotui`,
	Run: func(cmd *cobra.Command, args []string) {
		owner := h.LoadOwner()

		if owner == "" {
			log.Fatalf("Owner is empty. Either add data or use the search subcommand.")
		}

		nameAndType := strings.Split(owner, ":")

		if len(nameAndType) != 2 {
			log.Fatalf("Default owner must be in format \"name:type\" where type is either %q or %q", h.Org, h.User)
		}
	},
}
