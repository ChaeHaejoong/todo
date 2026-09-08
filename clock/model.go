package clock

import (
	"time"

	"github.com/chaehaejoong/todo/internal/apptime"
)

const (
	startHour = iota
	startMin
	endHour
	endMin
)

type Model struct {
	appTime *apptime.Apptime

	cursor int
}

func New(appTime *apptime.Apptime) Model {
	return Model{
		appTime: appTime,
	}
}

func (m Model) StartTime() string {
	return formatTime(m.appTime.GetStart())
}

func (m Model) EndTime() string {
	return formatTime(m.appTime.GetEnd())
}

func formatTime(t time.Time) string {
	return t.Format("15 : 04")
}
