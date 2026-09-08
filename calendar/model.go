package calendar

import (
	"time"

	"github.com/chaehaejoong/todo/internal/apptime"
)

type Model struct {
	appTime *apptime.Apptime
}

func New(appTime *apptime.Apptime) Model {
	return Model{
		appTime: appTime,
	}
}

func (m Model) daysInMonth() int {
	date := m.appTime.GetStart()

	return time.Date(
		date.Year(), date.Month()+1, 0,
		0, 0, 0, 0, time.Local,
	).Day()
}

func (m Model) firstWeekday() int {
	date := m.appTime.GetStart()

	return int(
		time.Date(
			date.Year(), date.Month(), 1,
			0, 0, 0, 0, time.Local,
		).Weekday(),
	)
}

func (m Model) withDay(day int) Model {
	date := m.appTime.GetStart()

	if day < 1 {
		day = 1
	}
	if day > m.daysInMonth() {
		day = m.daysInMonth()
	}

	m.appTime.SetDate(time.Date(
		date.Year(), date.Month(), day,
		0, 0, 0, 0, date.Location(),
	))
	return m
}

func (m Model) withMonth(year int, month time.Month) Model {
	day := m.appTime.GetStart().Day()
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	m.appTime.SetDate(date)
	return m.withDay(day)
}
