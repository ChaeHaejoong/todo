package todomodal

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/store"
)

type Model struct {
	isModalOpen bool
	input       textinput.Model
}

func New() Model {
	return Model{
		isModalOpen: false,
		input:       newTodoInput(),
	}
}

func (m Model) View(focused bool) string {
	if !m.isModalOpen {
		return ""
	}
	return renderTodoModal(m.input.View(), focused)
}

func (m *Model) OpenModal() tea.Cmd {
	m.isModalOpen = true
	return m.input.Focus()
}

func (m Model) IsModalOpen() bool {
	return m.isModalOpen
}

func (m Model) AppendTodo(todoStr string) error {
	return store.AppendTodo(todoStr)
}

func (m *Model) CloseModal() {
	m.isModalOpen = false
}
