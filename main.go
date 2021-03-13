// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"log"
	"strconv"
	"strings"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
	"github.com/google/go-github/github"
	r "github.com/irevenko/octostats/rest"
)

func renderList(results *github.UsersSearchResult) {
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
			renderGrid(user[0])
		}

		ui.Render(l)
	}
}

func renderGrid(username string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	ctx, client := r.AuthREST("f204672235aa1d2b321da46238a8498ad9a91d60")

	allRepos := r.AllRepos(ctx, client, username)

	totalStars := r.TotalStars(client, allRepos)

	p := widgets.NewParagraph()
	p.Text = "🍰: " + strconv.Itoa(totalStars)
	p.Title = username + " stats"

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
				renderInput()
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		}
	}
}

func renderInput() {
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
				ui.Clear()
				return
			case "<Left>":
				i.MoveCursorLeft()
			case "<Right>":
				i.MoveCursorRight()
			case "<Backspace>":
				i.Backspace()
			case "<Enter>":
				s := i.GetText()
				results := search(s)
				renderList(results)
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

func main() {
	renderInput()
}

func search(username string) *github.UsersSearchResult {
	ctx, client := r.AuthREST("f204672235aa1d2b321da46238a8498ad9a91d60")

	opts := &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 50}}

	users, _, err := client.Search.Users(ctx, username, opts)
	if err != nil {
		log.Fatal(err)
	}

	return users
}
