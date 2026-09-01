package calendar

import "github.com/chaehaejoong/todo/internal/ui"

func (m Model) View(focused bool) string {
	return ui.TitledBorder(39, 11, "[2] Calendar", m.renderCalendar(), focused)
}
