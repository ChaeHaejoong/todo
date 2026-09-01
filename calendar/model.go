package calendar

import "time"

type Model struct {
	year  int
	month time.Month
	day   int
}

func New() Model {
	now := time.Now()

	return Model{
		year:  now.Year(),
		month: now.Month(),
		day:   now.Day(),
	}
}

func (m Model) SelectedDate() time.Time {
	return time.Date(m.year, m.month, m.day, 0, 0, 0, 0, time.Local)
}

func (m Model) daysInMonth() int {
	return time.Date(m.year, m.month+1, 0, 0, 0, 0, 0, time.Local).Day()
}

func (m Model) firstWeekday() int {
	return int(time.Date(m.year, m.month, 1, 0, 0, 0, 0, time.Local).Weekday())
}

func (m Model) withDay(day int) Model {
	if day < 1 {
		day = 1
	}
	if day > m.daysInMonth() {
		day = m.daysInMonth()
	}

	m.day = day
	return m
}

func (m Model) withMonth(year int, month time.Month) Model {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	m.year = date.Year()
	m.month = date.Month()
	return m.withDay(m.day)
}
