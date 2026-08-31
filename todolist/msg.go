package todolist

import "github.com/chaehaejoong/todo/store"

type OpenTodoModalMsg struct{}

type TodoListLoadedMsg struct {
	Data store.Data
}

type TodoListLoadFailedMsg struct {
	err error
}

type TodoRemoveFailedMsg struct {
	err error
}
