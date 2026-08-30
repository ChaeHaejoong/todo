package todomodal

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if keyPressMsg, ok := msg.(tea.KeyPressMsg); ok {
		key := keyPressMsg.String()

		if isCloseModalKey(key) {
			m.CloseModal()
			return m, nil
		}

		if isAppendTodoKey(key) {
			if err := m.AppendTodo(m.input.Value()); err != nil {
				return m, nil
			}
			m.CloseModal()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
