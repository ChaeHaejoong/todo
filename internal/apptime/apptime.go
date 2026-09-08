package apptime

import "time"

type Apptime struct {
	start time.Time
	end   time.Time
}

func New() Apptime {
	return Apptime{}
}

func (a *Apptime) SetDate(t time.Time) {
	a.start = time.Date(
		t.Year(), t.Month(), t.Day(),
		a.start.Hour(), a.start.Minute(),
		0, 0, a.start.Location(),
	)

	a.end = time.Date(
		t.Year(), t.Month(), t.Day(),
		a.end.Hour(), a.end.Minute(),
		0, 0, a.end.Location(),
	)
}

func (a *Apptime) SetStartClock(t time.Time) {
	a.start = time.Date(
		a.start.Year(), a.start.Month(), a.start.Day(),
		t.Hour(), t.Minute(),
		0, 0, a.start.Location(),
	)
}

func (a *Apptime) SetEndClock(t time.Time) {
	a.end = time.Date(
		a.end.Year(), a.end.Month(), a.end.Day(),
		t.Hour(), t.Minute(),
		0, 0, a.start.Location(),
	)
}

func (a Apptime) GetStart() time.Time {
	return a.start
}

func (a Apptime) GetEnd() time.Time {
	return a.end
}
