package todolist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/store"
	"github.com/chaehaejoong/todo/todo"
)

func loadTodosCmd() tea.Cmd {
	return func() tea.Msg {
		data, err := store.LoadTodos()
		if err != nil {
			return TodoListLoadFailedMsg{err: err}
		}

		return TodoListLoadedMsg{Data: data}
	}
}

func removeTodoCmd(selected store.Todo) tea.Cmd {
	return func() tea.Msg {
		if err := store.RemoveTodo(selected); err != nil {
			return TodoRemoveFailedMsg{err: err}
		}

		return todo.TodoRemoveMsg{}
	}
}
