package tui

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
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
	l.TextStyle = ui.NewStyle(ui.ColorBlue)
	l.WrapText = true
	l.SetRect(0, 7, 100, 30)

	ui.Render(l)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			ui.Clear()
			ui.Close()
			os.Exit(1)
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "e", "<Enter>":
			user := strings.Split(l.Rows[l.SelectedRow], " ")
			ui.Clear()
			ui.Close()

			s := spinner.New(spinner.CharSets[30], 100*time.Millisecond)
			s.Prefix = "fetching github data "
			s.FinalMSG = "done"
			s.Start()

			RenderStats(user[0], s)
		}

		ui.Render(l)
	}
}
