package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	ui "github.com/gizak/termui/v3"
	"github.com/spf13/cobra"

	h "github.com/irevenko/octotui/helpers"
	tui "github.com/irevenko/octotui/tui"
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

		name := nameAndType[0]
		ownerType := h.OwnerType(nameAndType[1])

		s := spinner.New(spinner.CharSets[30], 100*time.Millisecond)
		s.Prefix = "fetching github data "
		s.FinalMSG = "done"
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()
		switch {
		case ownerType.IsOrg():
			s.Start()
			tui.RenderOrganization(name, s)
		case ownerType.IsUser():
			s.Start()
			tui.RenderUser(name, s)
		default:
			log.Fatalf("Expected either %q or %q, got %q in default_owner config file", h.Org, h.User, ownerType)
		}
	},
}
