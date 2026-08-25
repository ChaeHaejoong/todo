package app

import (
	todomodal "github.com/chaehaejoong/todo/todo_modal"
	tea "github.com/charmbracelet/bubbletea"
)

type Focus int

const (
	FocusMain Focus = iota
	FocusTodoModal
)

type Model struct {
	width  int
	height int

	focus     Focus
	todomodal todomodal.Model
}

func New() Model {
	return Model{
		focus:     FocusMain,
		todomodal: todomodal.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m = m.windowSizeUpdate(msg)

	var appCmd tea.Cmd
	m, appCmd = m.appUpdate(msg)

	var componentCmd tea.Cmd
	if m.focus == FocusTodoModal {
		m.todomodal, componentCmd = m.todomodal.Update(msg)
	}

	return m, tea.Batch(appCmd, componentCmd)
}

func (m Model) View() string {
	return "hello" + m.todomodal.View()
}
