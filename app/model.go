package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/appointment"
	"github.com/chaehaejoong/todo/calendar"
	"github.com/chaehaejoong/todo/clock"
	"github.com/chaehaejoong/todo/internal/apptime"
	"github.com/chaehaejoong/todo/internal/focus"
	"github.com/chaehaejoong/todo/todolist"
	"github.com/chaehaejoong/todo/todomodal"
)

type Model struct {
	width  int
	height int

	focus   focus.Manager
	apptime apptime.Apptime

	todolist    todolist.Model
	todomodal   todomodal.Model
	calendar    calendar.Model
	clock       clock.Model
	appointment appointment.Model
}

func New() Model {
	model := Model{
		focus:   focus.New(),
		apptime: apptime.New(),

		todolist:  todolist.New(),
		todomodal: todomodal.New(),
	}

	model.calendar = calendar.New(&model.apptime)
	model.clock = clock.New(&model.apptime)

	return model
}

func (m Model) Init() tea.Cmd {
	return m.todolist.Init()
}
