package tui

import (
	"log"
	"os"

	"github.com/briandowns/spinner"
	ui "github.com/gizak/termui"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
)

const (
	token = "fb71a5aa6f42d225f4210fb6b20bdebb81ce7cf5"
)

var (
	ctx, restClient = r.AuthREST(token)
	qlClient        = g.AuthGraphQL(token)
)

func RenderStats(username string, s *spinner.Spinner) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	user := g.UserDetails(qlClient, username)
	allRepos := r.AllRepos(ctx, restClient, username)

	img, images := SetupImage(user.AvatarURL, user.Login)
	p := SetupProfileInfo(user)
	p2 := SetupProfileStats(user, allRepos)
	p3 := SetupReposStats(user, allRepos)
	pc := SetupLangsByCommits(user)
	pc2 := SetupLangsByRepo(user, allRepos)
	bc := SetupStarsPerLangs(user, allRepos)
	bc2 := SetupForksPerLangs(user, allRepos)
	sl := SetupContribsSparkline(user)

	render := func() {
		img.Image = images[0]
		ui.Render(img, p, p2, pc, pc2, sl, bc, bc2, p3)
	}
	s.Stop()
	render()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			ui.Clear()
			ui.Close()
			os.Exit(1)
			return
		case "<Enter>":
			img.Monochrome = !img.Monochrome
		case "<Tab>":
			img.MonochromeInvert = !img.MonochromeInvert
		}
		render()
	}
}
