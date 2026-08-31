package todolist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/todo"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) { case TodoListLoadedMsg:
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
	}

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
