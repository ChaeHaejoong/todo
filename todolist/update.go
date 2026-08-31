package todolist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/todo"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TodoListLoadedMsg:
		m.todos = msg.Data.Todos
		if len(m.todos) == 0 {
			m.cursor = 0
		} else {
			m.cursor = min(m.cursor, len(m.todos)-1)
		}
		m.loading = false
		m.err = nil
		return m, nil

	case TodoListLoadFailedMsg:
		m.loading = false
		m.err = msg.err
		return m, nil

	case todo.TodoSubmittedMsg:
		return m, loadTodosCmd()

	case todo.TodoRemoveMsg:
		return m, loadTodosCmd()

	case TodoRemoveFailedMsg:
		m.err = msg.err
		return m, nil

	case tea.KeyPressMsg:
		if isOpenModalKey(msg.String()) {
			return m, func() tea.Msg { return OpenTodoModalMsg{} }
		}

		if isDeleteKey(msg.String()) {
			selected, ok := m.selectedTodo()
			if !ok {
				return m, nil
			}

			return m, removeTodoCmd(selected)
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
