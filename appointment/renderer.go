package appointment

import "github.com/chaehaejoong/todo/internal/ui"

func render(width, height int, content string, focused bool) string {
	return ui.TitledBorder(width, height,"[4] Appointment", content, focused)
}
