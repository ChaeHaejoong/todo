package calendar

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/chaehaejoong/todo/internal/ui"
)

func (m Model) renderCalendar() string {
	var lines []string

	weekdays := []string{"Sun", "Mon", "Tue", "Wen", "Thu", "Fri", "Sat"}
	weekdayLine := make([]string, 0, len(weekdays))
	for _, weekday := range weekdays {
		weekdayLine = append(weekdayLine, ui.CalendarCell(false, weekday))
	}
	lines = append(lines, strings.Join(weekdayLine, ""))

	week := make([]string, 0, 7)
	for i := 0; i < m.firstWeekday(); i++ {
		week = append(week, ui.CalendarCell(false, ""))
	}

	for day := 1; day <= m.daysInMonth(); day++ {
		week = append(
			week,
			ui.CalendarCell(day == m.day, fmt.Sprintf("%d", day)),
		)
		if len(week) == 7 {
			lines = append(lines, strings.Join(week, ""))
			week = week[:0]
		}
	}

	if len(week) > 0 {
		for len(week) < 7 {
			week = append(week, ui.CalendarCell(false, ""))
		}
		lines = append(lines, strings.Join(week, ""))
	}

	footer := lipgloss.NewStyle().Render(
		fmt.Sprintf("\n > %d-%d-%d", m.year, m.month, m.day),
	)
	lines = append(lines, footer)

	return strings.Join(lines, "\n")
}
