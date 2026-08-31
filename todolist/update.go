package todolist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/todo"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TodoListLoadedMsg:
		m.todos = msg.Data.Todos
		m.loading = false
		m.err = nil
		return m, nil

	case TodoListLoadFailedMsg:
		m.loading = false
		m.err = msg.err
		return m, nil

	case todo.TodoSubmittedMsg:
		return m, loadTodosCmd()

	case tea.KeyPressMsg:
		if isOpenModalKey(msg.String()) {
			return m, func() tea.Msg { return OpenTodoModalMsg{} }
		}

		if isCursorDownKey(msg.String()) {
			m.cursor = min(len(m.todos)-1, m.cursor+1)
			return m, nil
		}

		if isCursorUpKey(msg.String()) {
			m.cursor = max(0, m.cursor-1)
			return m, nil
		}
	}

	return m, nil
}
