package todomodal

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	isModalOpen bool
	input       textinput.Model
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.isModalOpen {
		return m, nil
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok && isCloseModalKey(keyMsg.String()) {
		m.CloseModal()
		return m, nil
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if !m.isModalOpen {
		return ""
	}
	return "Modal"
}

func (m *Model) OpenModal() {
	m.isModalOpen = true
	m.input.Focus()
}

func (m Model) IsModalOpen() bool {
	return m.isModalOpen
}

func (m *Model) CloseModal() {
	m.isModalOpen = false
}
