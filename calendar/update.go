package calendar

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "left", "h":
		return m.withDay(m.day - 1), nil
	case "right", "l":
		return m.withDay(m.day + 1), nil
	case "up", "k":
		return m.withDay(m.day - 7), nil
	case "down", "j":
		return m.withDay(m.day + 7), nil
	case "H":
		return m.previousMonth(), nil
	case "L":
		return m.nextMonth(), nil
	}

	return m, nil
}

func (m Model) previousMonth() Model {
	return m.withMonth(m.year, m.month-1)
}

func (m Model) nextMonth() Model {
	return m.withMonth(m.year, m.month+1)
}
