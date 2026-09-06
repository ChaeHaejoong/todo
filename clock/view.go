package clock

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"
	"github.com/chaehaejoong/todo/internal/ui"
)

const (
	startInput = iota
	endInput
)

type Model struct {
	start   textinput.Model
	end     textinput.Model
	current int
}

func New() Model {
	start := newTimeInput()
	end := newTimeInput()
	start.Focus()

	return Model{
		start: start,
		end:   end,
	}
}

func (m Model) View(width, height int, focused bool) string {
	startRow := inputRow("Start", m.start.View(), m.current == startInput && focused)
	endRow := inputRow("End", m.end.View(), m.current == endInput && focused)
	content := lipgloss.JoinVertical(lipgloss.Left, startRow, endRow)

	return ui.TitledBorder(width, height, "[4] Clock", content, focused)
}

func (m Model) StartTime() string {
	return m.start.Value()
}

func (m Model) EndTime() string {
	return m.end.Value()
}

func newTimeInput() textinput.Model {
	input := textinput.New()
	input.Prompt = ""
	input.Placeholder = "HH:MM"
	input.CharLimit = 5
	input.SetWidth(5)

	styles := textinput.DefaultDarkStyles()
	styles.Cursor.Blink = true
	input.SetStyles(styles)

	return input
}

func inputRow(label, input string, active bool) string {
	labelStyle := lipgloss.NewStyle().Width(6)
	if active {
		labelStyle = labelStyle.Foreground(ui.ColorFocused)
	}

	return labelStyle.Render(label+":") + input
}
