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
