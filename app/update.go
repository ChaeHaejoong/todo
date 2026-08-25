package app

import tea "github.com/charmbracelet/bubbletea"

func (m Model) windowSizeUpdate(msg tea.Msg) Model {
	if windowSizeMsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = windowSizeMsg.Width
		m.height = windowSizeMsg.Height
	}

	return m
}

func (m Model) appUpdate(msg tea.Msg) (Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		key := keyMsg.String()

		if isQuitKey(key) {
			return m, tea.Quit
		}

		if !m.todomodal.IsModalOpen() && isOpenTodoModalKey(key) {
			m.focus = FocusTodoModal
			return m, m.todomodal.OpenModal()
		}
	}

	var cmd tea.Cmd
	return m, cmd
}
