package todomodal

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/internal/apptime"
	"github.com/chaehaejoong/todo/store"
)

type Model struct {
	isModalOpen bool
	input       textinput.Model
	appTime     *apptime.Apptime
	includeTime bool
}

func New(appTime *apptime.Apptime) Model {
	return Model{
		isModalOpen: false,
		input:       newTodoInput(),
		appTime:     appTime,
		includeTime: true,
	}
}

func (m Model) View(width, height int, focused bool) string {
	if !m.isModalOpen {
		return ""
	}
	return renderTodoModal(width, height, m.input.View(), m.includeTime, focused)
}

func (m *Model) OpenModal() tea.Cmd {
	m.isModalOpen = true
	return m.input.Focus()
}

func (m Model) IsModalOpen() bool {
	return m.isModalOpen
}

func AppendTodo(todoStr string, appTime apptime.Apptime, includeTime bool) error {
	return store.AppendTodo(todoStr, appTime, includeTime)
}

func (m *Model) CloseModal() {
	m.isModalOpen = false
}
