package todomodal

import (
	"charm.land/bubbles/v2/textinput"
	ui "github.com/chaehaejoong/todo/internal/ui"
)

func newTodoInput() textinput.Model {
	input := textinput.New()
	input.Prompt = ""
	input.SetWidth(35)

	styles := textinput.DefaultDarkStyles()
	styles.Cursor.Blink = true
	input.SetStyles(styles)

	return input
}

func renderTodoModal(content string, focused bool) string {
	return ui.TitledBorder(40, 3, "Append-Todo", content, focused)
}
