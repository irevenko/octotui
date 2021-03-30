package cmd

import (
	"strings"

	gh "github.com/irevenko/octotui/github"

	tui "github.com/irevenko/octotui/tui"
	"github.com/spf13/cobra"
)

var Search = &cobra.Command{
	Use:   "search",
	Short: "Search for the github profile",
	Long:  `octotui search <USER_OR_ORGANIZATION>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchVal := strings.Join(args, " ")
		results := gh.SearchUser(tui.Ctx, tui.RestClient, searchVal)
		tui.RenderList(results)
	},
}

var ByRemote = &cobra.Command{
	Use:   "by-remote",
	Short: "Get github profile by remote URL",
	Long:  `octotui by-remote [REMOTE_NAME]`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		panic("Not implemented")
	},
}

func AddCommands() {
	RootCmd.AddCommand(Search)
	RootCmd.AddCommand(ByRemote)
}
