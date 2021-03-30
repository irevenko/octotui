package cmd

import (
	"log"
	"os/exec"
	"strings"
	"time"

	gh "github.com/irevenko/octotui/github"
	help "github.com/irevenko/octotui/helpers"

	"github.com/briandowns/spinner"
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

var remoteName string

var ByRemote = &cobra.Command{
	Use:   "by-remote",
	Short: "Search for github profile by remote URL",
	Run: func(cmd *cobra.Command, args []string) {
		gitArgs := []string{"remote", "get-url", remoteName}
		gitCmd := exec.Command("git", gitArgs...)
		remoteBytes, err := gitCmd.Output()

		if err != nil {
			log.Fatalf("failed to get remote URL: %v", err)
		}

		owner, err := help.OwnerFromRemote(string(remoteBytes))

		if err != nil {
			log.Fatalf("failed to get owner from remote URL: %v", err)
		}

		results := gh.SearchUser(tui.Ctx, tui.RestClient, owner)
		total := results.GetTotal()

		if total == 0 {
			log.Fatalf("No user/organization found matching %q", owner)
		} else if total > 1 {
			tui.RenderList(results)
		} else {
			user := results.Users[0]
			username := *user.Login
			accountType := "(" + strings.ToLower(*user.Type) + ")"

			s := spinner.New(spinner.CharSets[30], 100*time.Millisecond)
			s.Prefix = "fetching github data "
			s.FinalMSG = "done"
			s.Start()
			tui.RenderStats(username, accountType, s)
		}

	},
}

func AddCommands() {
	RootCmd.AddCommand(Search)
	RootCmd.AddCommand(ByRemote)
}

func init() {
	ByRemote.PersistentFlags().StringVarP(
		&remoteName,
		"remote",
		"r",
		"origin",
		"remote containing the user/organization to search",
	)
}
