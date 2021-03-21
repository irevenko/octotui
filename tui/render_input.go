package tui

import (
	"log"

	gh "../github"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func RenderInput() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.WrapText = true
	p.Text = "Enter GitHub username or organization name"
	p.Border = true
	p.SetRect(0, 0, 45, 4)

	i := widgets.NewTextBox()
	i.SetRect(0, 4, 45, 7)
	i.ShowCursor = true

	ui.Render(i, p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<C-c>":
				return
			case "<Left>":
				i.MoveCursorLeft()
			case "<Right>":
				i.MoveCursorRight()
			case "<Backspace>":
				i.Backspace()
			case "<Enter>":
				user := i.GetText()
				results := gh.SearchUser(ctx, restClient, user)
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
