package tui

import (
	"log"
	"strconv"
	"strings"

	gh "../github"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
	"github.com/google/go-github/github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
)

var (
	ctx, client = r.AuthREST("")
	qlClient    = g.AuthGraphQL("")
)

func RenderList(results *github.UsersSearchResult) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var rowsData []string

	for _, v := range results.Users {
		login := *v.Login
		accountType := *v.Type
		rowsData = append(rowsData, login+" ("+strings.ToLower(accountType)+")")
	}

	l := widgets.NewList()

	l.Title = "Search Results"
	l.Rows = rowsData
	l.TextStyle = ui.NewStyle(ui.ColorGreen)
	l.WrapText = true
	l.SetRect(1, 4, 100, 30)

	ui.Render(l)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Clear()
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "e", "<Enter>":
			user := strings.Split(l.Rows[l.SelectedRow], " ")
			RenderGrid(user[0])
		}

		ui.Render(l)
	}
}

func RenderGrid(username string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.WrapText = true

	langs, stars, forks := gh.FetchStats(ctx, client, qlClient, username)

	p.Text = "Total Stars: " + strconv.Itoa(stars) + "\nTotal Forks: " + strconv.Itoa(forks) + "\nMost used langs: \n" + strings.Join(langs[:], ",")

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, p),
		),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				ui.Clear()
				RenderInput()
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		}
	}
}

func RenderInput() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	i := widgets.NewTextBox()
	i.SetRect(1, 1, 40, 4)
	i.ShowCursor = true

	ui.Render(i)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<C-c>":
				ui.Close()
				return
			case "<Left>":
				i.MoveCursorLeft()
			case "<Right>":
				i.MoveCursorRight()
			case "<Backspace>":
				i.Backspace()
			case "<Enter>":
				user := i.GetText()
				results := gh.SearchUser(ctx, client, user)
				RenderList(results)
			case "<Space>":
				i.InsertText(" ")
			default:
				if ui.ContainsString(ui.PRINTABLE_KEYS, e.ID) {
					i.InsertText(e.ID)
				}
			}
			ui.Render(i)
		}
	}
}
