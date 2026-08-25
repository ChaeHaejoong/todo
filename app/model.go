package app

import (
	todomodal "github.com/chaehaejoong/todo/todo_modal"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	focus     int
	todomodal todomodal.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m, cmd = m.appUpdate(msg)
	m.todomodal, cmd = m.todomodal.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return "hello" + m.todomodal.View()
}
