package tui

import (
	"log"

	ui "github.com/gizak/termui"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
)

const (
	token = ""
)

var (
	ctx, client = r.AuthREST(token)
	qlClient    = g.AuthGraphQL(token)
)

func RenderStats(username string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	user := g.UserDetails(qlClient, username)
	p := SetupProfileInfo(user)
	p2 := SetupProfileStats(user)
	pc := SetupLangsByCommits(user)

	img, images := SetupImage(user.AvatarURL, user.Login)

	render := func() {
		img.Image = images[0]
		ui.Render(img, p, p2, pc)
	}
	render()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Clear()
			RenderInput()
		case "<Enter>":
			img.Monochrome = !img.Monochrome
		case "<Tab>":
			img.MonochromeInvert = !img.MonochromeInvert
		}
		render()
	}
}
