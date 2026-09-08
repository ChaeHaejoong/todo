package clock

import "github.com/chaehaejoong/todo/internal/ui"

func (m Model) View(width, height int, focused bool) string {
	content := renderClock(
		m.appTime.GetStart(),
		m.appTime.GetEnd(),
		m.cursor,
		focused,
	)
	return ui.TitledBorder(width, height, "[3] Clock", content, focused)
}
