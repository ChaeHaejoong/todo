package clock

import (
	"fmt"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/chaehaejoong/todo/internal/ui"
)

func renderClock(start, end time.Time, current int, focused bool) string {
	startLine := fmt.Sprintf(
		"\nStart: %s : %s",
		cursor(start.Hour(), current == startHour, focused),
		cursor(start.Minute(), current == startMin, focused),
	)
	endLine := fmt.Sprintf(
		"\nEnd:   %s : %s\n",
		cursor(end.Hour(), current == endHour, focused),
		cursor(end.Minute(), current == endMin, focused),
	)

	return lipgloss.JoinVertical(lipgloss.Left, startLine, endLine)
}

func cursor(value int, active, focused bool) string {
	text := fmt.Sprintf("%02d", value)
	if !active || !focused {
		return text
	}

	return lipgloss.NewStyle().
		Background(ui.ColorCursored).
		Render(text)
}
