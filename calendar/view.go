package calendar

import (
	"charm.land/lipgloss/v2"
	"github.com/chaehaejoong/todo/internal/ui"
)

func (m Model) View(width, height int, focused bool) string {
	contentWidth, contentHeight := ui.TitledBorderContentSize(width, height)
	content := lipgloss.Place(
		contentWidth,
		contentHeight,
		lipgloss.Center,
		lipgloss.Center,
		m.renderCalendar(),
	)

	return ui.TitledBorder(width, height, "[2] Calendar", content, focused)
}
