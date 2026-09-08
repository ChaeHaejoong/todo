package clock

import (
	"charm.land/lipgloss/v2"
	"github.com/chaehaejoong/todo/internal/ui"
)

func (m Model) View(width, height int, focused bool) string {
	content := renderClock(
		m.appTime.GetStart(),
		m.appTime.GetEnd(),
		m.cursor,
		focused,
	)
	contentWidth, contentHeight := ui.TitledBorderContentSize(width, height)
	content = lipgloss.Place(
		contentWidth,
		contentHeight,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)

	return ui.TitledBorder(width, height, "[3] Clock", content, focused)
}
