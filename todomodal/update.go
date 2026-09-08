package todomodal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/todo"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if keyPressMsg, ok := msg.(tea.KeyPressMsg); ok {
		key := keyPressMsg.String()

		if isToggleTimeKey(key) {
			m.includeTime = !m.includeTime
			return m, nil
		}

		if isCloseModalKey(key) {
			m.CloseModal()
			return m, nil
		}

		if isAppendTodoKey(key) {
			if err := AppendTodo(m.input.Value(), *m.appTime, m.includeTime); err != nil {
				return m, nil
			}
			m.CloseModal()
			m.input.SetValue("")
			return m, func() tea.Msg { return todo.TodoSubmittedMsg{} }
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
