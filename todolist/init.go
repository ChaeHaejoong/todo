package todolist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/store"
)

func (m Model) Init() tea.Cmd {
	return loadTodosCmd()
}

func loadTodosCmd() tea.Cmd {
	return func() tea.Msg {
		data, err := store.LoadTodos()
		if err != nil {
			return TodoListLoadFailedMsg{err: err}
		}

		return TodoListLoadedMsg{Data: data}
	}
}
