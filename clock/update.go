package clock

import "time"

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "h":
		m.cursor = wrap(m.cursor-1, endMin+1)
	case "l":
		m.cursor = wrap(m.cursor+1, endMin+1)
	case "j":
		m.adjust(-1)
	case "k":
		m.adjust(+1)
	case "J":
		m.adjust(-10)
	case "K":
		m.adjust(+10)
	case "0":
		m.setCurrent(0)
	}

	return m, nil
}

func (m Model) adjust(delta int) {
	switch m.cursor {
	case startHour:
		start := m.appTime.GetStart()
		m.appTime.SetStartClock(time.Date(
			start.Year(), start.Month(), start.Day(),
			wrap(start.Hour()+delta, 24), start.Minute(),
			0, 0, start.Location(),
		))
	case startMin:
		start := m.appTime.GetStart()
		m.appTime.SetStartClock(time.Date(
			start.Year(), start.Month(), start.Day(),
			start.Hour(), wrap(start.Minute()+delta, 60),
			0, 0, start.Location(),
		))
	case endHour:
		end := m.appTime.GetEnd()
		m.appTime.SetEndClock(time.Date(
			end.Year(), end.Month(), end.Day(),
			wrap(end.Hour()+delta, 24), end.Minute(),
			0, 0, end.Location(),
		))
	case endMin:
		end := m.appTime.GetEnd()
		m.appTime.SetEndClock(time.Date(
			end.Year(), end.Month(), end.Day(),
			end.Hour(), wrap(end.Minute()+delta, 60),
			0, 0, end.Location(),
		))
	}
}

func (m Model) setCurrent(value int) {
	switch m.cursor {
	case startHour, startMin:
		start := m.appTime.GetStart()
		hour, minute := start.Hour(), start.Minute()
		if m.cursor == startHour {
			hour = value
		} else {
			minute = value
		}
		m.appTime.SetStartClock(time.Date(
			start.Year(), start.Month(), start.Day(),
			hour, minute, 0, 0, start.Location(),
		))
	case endHour, endMin:
		end := m.appTime.GetEnd()
		hour, minute := end.Hour(), end.Minute()
		if m.cursor == endHour {
			hour = value
		} else {
			minute = value
		}
		m.appTime.SetEndClock(time.Date(
			end.Year(), end.Month(), end.Day(),
			hour, minute, 0, 0, end.Location(),
		))
	}
}

func wrap(value, size int) int {
	value %= size
	if value < 0 {
		value += size
	}
	return value
}
