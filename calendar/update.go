package calendar

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "left", "h":
		return m.withDay(m.currentDay() - 1), nil
	case "right", "l":
		return m.withDay(m.currentDay() + 1), nil
	case "up", "k":
		return m.withDay(m.currentDay() - 7), nil
	case "down", "j":
		return m.withDay(m.currentDay() + 7), nil
	case "H":
		return m.previousMonth(), nil
	case "L":
		return m.nextMonth(), nil
	}

	return m, nil
}

func (m Model) previousMonth() Model {
	date := m.appTime.GetStart()
	return m.withMonth(date.Year(), date.Month()-1)
}

func (m Model) nextMonth() Model {
	date := m.appTime.GetStart()
	return m.withMonth(date.Year(), date.Month()+1)
}

func (m Model) currentDay() int {
	return m.appTime.GetStart().Day()
}
