package todolist

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	keyPressMsg, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}

	if isOpenModalKey(keyPressMsg.String()) {
		return m, func() tea.Msg {
			return OpenTodoModalMsg{}
		}
	}

	return m, nil
}
