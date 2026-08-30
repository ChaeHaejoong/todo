package app

import (
	"github.com/chaehaejoong/todo/internal/focus"
	"github.com/chaehaejoong/todo/todolist"
	"github.com/chaehaejoong/todo/todomodal"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	width  int
	height int

	focus focus.Manager

	todolist  todolist.Model
	todomodal todomodal.Model
}

func New() Model {
	return Model{
		focus:     focus.New(),
		todolist:  todolist.New(),
		todomodal: todomodal.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.todolist.Init() 
}



