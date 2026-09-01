package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/appointment"
	"github.com/chaehaejoong/todo/calendar"
	"github.com/chaehaejoong/todo/internal/focus"
	"github.com/chaehaejoong/todo/todolist"
	"github.com/chaehaejoong/todo/todomodal"
)

type Model struct {
	width  int
	height int

	focus focus.Manager

	todolist    todolist.Model
	todomodal   todomodal.Model
	calendar    calendar.Model
	appointment appointment.Model
}

func New() Model {
	return Model{
		focus:     focus.New(),
		todolist:  todolist.New(),
		todomodal: todomodal.New(),
		calendar:  calendar.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.todolist.Init()
}
