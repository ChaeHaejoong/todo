package ui

import "charm.land/lipgloss/v2"

func CalendarCell(cursored bool, content string) string {
	style := lipgloss.NewStyle().
		Width(5).
		Align(lipgloss.Center)

	if cursored {
		style = style.Background(ColorCursored)
	}

	return style.Render(content)
}
