package calendar

import "github.com/chaehaejoong/todo/internal/ui"

func (m Model) View(width, height int, focused bool) string {
	return ui.TitledBorder(width, height, "[2] Calendar", m.renderCalendar(), focused)
}
