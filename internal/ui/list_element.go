package ui

import "charm.land/lipgloss/v2"

func ListElement(content string, cursored bool) string {
	if !cursored {
		return content
	}

	return lipgloss.NewStyle().
		Background(ColorCursored).
		Render(content)
}
