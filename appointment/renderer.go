package appointment

import (
	"strings"

	"github.com/chaehaejoong/todo/internal/ui"
	"github.com/chaehaejoong/todo/store"
)

func render(width, height int, content string, focused bool) string {
	return ui.TitledBorder(width, height, "[4] Appointment", content, focused)
}

func renderAppointments(appointments []store.Appointment, cursor int) string {
	if len(appointments) == 0 {
		return "No appointments"
	}

	lines := make([]string, 0, len(appointments))
	for index, item := range appointments {
		content := "- " + item.Content
		lines = append(lines, ui.ListElement(content, index == cursor))
	}
	return strings.Join(lines, "\n")
}
