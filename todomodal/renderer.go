package todomodal

import (
	"fmt"

	"charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"
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

func renderTodoModal(width, height int, content string, includeTime, focused bool) string {
	contentWidth, contentHeight := ui.TitledBorderContentSize(width, height)
	option := fmt.Sprintf("[%c] time", map[bool]rune{true: 'x', false: ' '}[includeTime])
	option = lipgloss.PlaceHorizontal(contentWidth, lipgloss.Right, option)
	content = lipgloss.JoinVertical(lipgloss.Left, content, option)
	content = lipgloss.Place(contentWidth, contentHeight, lipgloss.Left, lipgloss.Top, content)

	return ui.TitledBorder(width, height, "Append-Todo", content, focused)
}
