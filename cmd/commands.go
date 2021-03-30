package cmd

import (
	"log"
	"os/exec"
	"strings"

	gh "github.com/irevenko/octotui/github"
	help "github.com/irevenko/octotui/helpers"

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
		gitArgs := []string{"remote", "get-url"}
		gitArgs = append(gitArgs, args...)
		gitCmd := exec.Command("git", gitArgs...)
		remoteBytes, err := gitCmd.Output()

		if err != nil {
			log.Fatalf("failed to get remote URL: %v", err)
		}

		owner, err := help.OwnerFromRemote(string(remoteBytes))

		if err != nil {
			log.Fatalf("failed to get owner from remote URL: %v", err)
		}

		println("owner:", owner)

		panic("Not implemented")
	},
}

func AddCommands() {
	RootCmd.AddCommand(Search)
	RootCmd.AddCommand(ByRemote)
}