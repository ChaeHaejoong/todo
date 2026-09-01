package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/internal/focus"
	"github.com/chaehaejoong/todo/todolist"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.setWindowSize(msg)

	m, cmd := m.appUpdate(msg)
	if cmd != nil {
		return m, cmd
	}

	if m.focus.Is(focus.TodoModal) {
		var todoModalCmd tea.Cmd
		m.todomodal, todoModalCmd = m.todomodal.Update(msg)

		if !m.todomodal.IsModalOpen() {
			m.focus.Set(focus.TodoList)
		}

		return m, todoModalCmd
	}

	if m.focus.Is(focus.TodoList) {
		var todoListCmd tea.Cmd
		m.todolist, todoListCmd = m.todolist.Update(msg)

		return m, todoListCmd
	}

	if m.focus.Is(focus.Calendar) {
		var calendarCmd tea.Cmd
		m.calendar, calendarCmd = m.calendar.Update(msg)

		return m, calendarCmd
	}

	return m, nil
}

func (m Model) appUpdate(msg tea.Msg) (Model, tea.Cmd) {
	if keyPressMsg, ok := msg.(tea.KeyPressMsg); ok {
		key := keyPressMsg.String()
		if isQuitKey(key) {
			return m, tea.Quit
		}

		if isTodoListFocusKey(key) {
			m.focus.Set(focus.TodoList)
			return m, nil
		}

		if isCalendarFocusKey(key) {
			m.focus.Set(focus.Calendar)
			return m, nil
		}
	}

	if _, ok := msg.(todolist.OpenTodoModalMsg); ok {
		m.focus.Set(focus.TodoModal)
		return m, m.todomodal.OpenModal()
	}

	return m, nil
}

func (m *Model) setWindowSize(msg tea.Msg) {
	if windowSizeMsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = windowSizeMsg.Width
		m.height = windowSizeMsg.Height
	}
}
