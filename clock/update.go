package clock

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok {
		switch key.String() {
		case "tab", "down":
			return m.focusInput(endInput)
		case "shift+tab", "up":
			return m.focusInput(startInput)
		}
	}

	var cmd tea.Cmd
	if m.current == startInput {
		m.start, cmd = m.start.Update(msg)
	} else {
		m.end, cmd = m.end.Update(msg)
	}

	return m, cmd
}

func (m Model) focusInput(input int) (Model, tea.Cmd) {
	m.start.Blur()
	m.end.Blur()
	m.current = input

	if input == startInput {
		return m, m.start.Focus()
	}

	return m, m.end.Focus()
}
