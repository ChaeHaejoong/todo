package appointment

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/chaehaejoong/todo/internal/ui"
	"github.com/chaehaejoong/todo/store"
)

func render(width, height int, content string, focused bool) string {
	return ui.TitledBorder(width, height, "[4] Appointment", content, focused)
}

func renderAppointments(width, height int, appointments []store.Appointment, cursor int, focused bool) string {
	if len(appointments) == 0 {
		return "No appointments"
	}

	contentWidth, _ := ui.TitledBorderContentSize(width, height)
	lines := make([]string, 0, len(appointments))
	for index, item := range appointments {
		content := renderAppointment(contentWidth, item)
		lines = append(lines, ui.ListElement(content, index == cursor && focused))
	}
	return strings.Join(lines, "\n")
}

func renderAppointment(width int, appointment store.Appointment) string {
	content := "- " + appointment.Content
	metadata := strings.TrimSpace(fmt.Sprintf("%s %s", appointment.Date, appointment.Time))
	if metadata == "" {
		return content
	}

	contentWidth := lipgloss.Width(content)
	metadataWidth := lipgloss.Width(metadata)
	separatorWidth := width - contentWidth - metadataWidth - 2
	if separatorWidth >= 1 {
		return content + " " + strings.Repeat("-", separatorWidth) + metadata + "~"
	}

	metadataLine := strings.Repeat("-", max(1, width-metadataWidth)) + metadata
	return content + "\n" + metadataLine
}
