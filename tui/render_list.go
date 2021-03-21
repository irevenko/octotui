package tui

import (
	"log"
	"strings"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
	"github.com/google/go-github/github"
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
	l.TextStyle = ui.NewStyle(ui.ColorRed)
	l.WrapText = true
	l.SetRect(0, 7, 100, 30)

	ui.Render(l)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "e", "<Enter>":
			user := strings.Split(l.Rows[l.SelectedRow], " ")
			ui.Clear()
			RenderStats(user[0])
		}

		ui.Render(l)
	}
}
