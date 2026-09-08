package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/appointment"
	"github.com/chaehaejoong/todo/calendar"
	"github.com/chaehaejoong/todo/clock"
	"github.com/chaehaejoong/todo/entrymodal"
	"github.com/chaehaejoong/todo/internal/apptime"
	"github.com/chaehaejoong/todo/internal/focus"
	"github.com/chaehaejoong/todo/todolist"
)

type Model struct {
	width  int
	height int

	focus            focus.Manager
	apptime          apptime.Apptime
	modalReturnFocus focus.Target

	todolist    todolist.Model
	entrymodal  entrymodal.Model
	calendar    calendar.Model
	clock       clock.Model
	appointment appointment.Model
}

func New() Model {
	model := Model{
		focus:   focus.New(),
		apptime: apptime.New(),

		todolist:    todolist.New(),
		appointment: appointment.New(),
	}

	model.entrymodal = entrymodal.New(&model.apptime)

	model.calendar = calendar.New(&model.apptime)
	model.clock = clock.New(&model.apptime)

	return model
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.todolist.Init(), m.appointment.Init())
}
