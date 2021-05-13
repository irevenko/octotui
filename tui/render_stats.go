package tui

import (
	"log"
	"os"

	"github.com/briandowns/spinner"
	ui "github.com/gizak/termui/v3"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
	h "github.com/irevenko/octotui/helpers"
)

var (
	token = h.LoadToken()
)

var (
	Ctx, RestClient = r.AuthREST(token)
	qlClient        = g.AuthGraphQL(token)
)

func RenderStats(username string, accType string, s *spinner.Spinner) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	width, height := ui.TerminalDimensions()
	grid := ui.NewGrid()
	grid.SetRect(0, 0, width, height)

	if accType == "(user)" {
		renderUser(username, s, grid)
	}

	if accType == "(organization)" {
		renderOrganization(username, s)
	}
}

func renderUser(username string, s *spinner.Spinner, grid *ui.Grid) {
	user, err := g.UserDetails(qlClient, username)
	if err != nil {
		log.Fatalf("Couldn't get user details for: %v: %v", username, err)
	}
	allRepos, err := r.AllRepos(Ctx, RestClient, username)
	if err != nil {
		log.Fatalf("Couldn't get all repos for: %v: %v", username, err)
	}

	img, images := SetupImage(user.AvatarURL, user.Login)
	p := SetupProfileInfo(user)
	p2 := SetupProfileStats(user, allRepos)
	p3 := SetupReposStats(user, allRepos)
	pc := SetupLangsByCommits(user)
	pc2 := SetupLangsByRepo(allRepos, "user")
	bc := SetupStarsPerLangs(allRepos, "user")
	bc2 := SetupForksPerLangs(allRepos, "user")
	sl := SetupContribsSparkline(user)

	render := func() {
		img.Image = images[0]
		grid.Set(
			ui.NewCol(
				0.25,
				ui.NewRow(1.0/3, img),
				ui.NewRow(2.0/3, p),
			),
			ui.NewCol(
				0.25,
				ui.NewRow(0.625, p2),
				ui.NewRow(0.375, pc),
			),
			ui.NewCol(
				0.5,
				ui.NewRow(
					0.25,
					sl,
				),
				ui.NewRow(
					0.75,
					ui.NewCol(
						0.5,
						ui.NewRow(0.5, p3),
						ui.NewRow(0.5, pc2),
					),
					ui.NewCol(
						0.5,
						ui.NewRow(0.5, bc2),
						ui.NewRow(0.5, bc),
					),
				),
			),
		)
		ui.Render(grid)
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

func renderOrganization(username string, s *spinner.Spinner) {
	org, err := g.OrganizationDetails(qlClient, username)
	if err != nil {
		log.Fatalf("Couldn't get org details for: %v: %v", username, err)
	}
	allRepos, err := r.AllRepos(Ctx, RestClient, username)
	if err != nil {
		log.Fatalf("Couldn't get all repos for: %v: %v", username, err)
	}

	img, images := SetupImage(org.AvatarURL, org.Login)
	p := SetupOrgInfo(org)
	p2 := SetupOrgStats(org, allRepos)
	p3 := SetupOrgRepos(org, allRepos)
	bc := SetupStarsPerLangs(allRepos, "organization")
	bc2 := SetupForksPerLangs(allRepos, "organization")
	pc2 := SetupLangsByRepo(allRepos, "organization")

	render := func() {
		img.Image = images[0]
		ui.Render(img, p, p2, p3, bc, bc2, pc2)
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
