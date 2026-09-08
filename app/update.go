package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/appointment"
	"github.com/chaehaejoong/todo/entrymodal"
	"github.com/chaehaejoong/todo/internal/focus"
	"github.com/chaehaejoong/todo/todolist"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.setWindowSize(msg)

	m, cmd := m.appUpdate(msg)
	if cmd != nil {
		return m, cmd
	}

	switch msg.(type) {
	case appointment.AppointmentLoadedMsg,
		appointment.AppointmentLoadFailedMsg,
		appointment.AppointmentRemovedMsg,
		appointment.AppointmentRemoveFailedMsg,
		entrymodal.AppointmentSubmittedMsg:
		var appointmentCmd tea.Cmd
		m.appointment, appointmentCmd = m.appointment.Update(msg)
		return m, appointmentCmd
	}

	if m.focus.Is(focus.TodoModal) {
		var todoModalCmd tea.Cmd
		m.entrymodal, todoModalCmd = m.entrymodal.Update(msg)

		if !m.entrymodal.IsModalOpen() {
			m.focus.Set(m.modalReturnFocus)
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

	if m.focus.Is(focus.Appointment) {
		var appointmentCmd tea.Cmd
		m.appointment, appointmentCmd = m.appointment.Update(msg)

		return m, appointmentCmd
	}

	if m.focus.Is(focus.Clock) {
		var clockCmd tea.Cmd
		m.clock, clockCmd = m.clock.Update(msg)

		return m, clockCmd
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

		if isAppointmentFocusKey(key) {
			m.focus.Set(focus.Appointment)
			return m, nil
		}

		if isClockFocusKey(key) {
			m.focus.Set(focus.Clock)
			return m, nil
		}

		if m.focus.Is(focus.TodoList) {
			switch key {
			case "a":
				m.modalReturnFocus = focus.TodoList
				m.focus.Set(focus.TodoModal)
				return m, m.entrymodal.OpenTodoAdd()
			case "r":
				selected, ok := m.todolist.SelectedTodo()
				if ok {
					m.modalReturnFocus = focus.TodoList
					m.focus.Set(focus.TodoModal)
					return m, m.entrymodal.OpenTodoEdit(selected)
				}
			}
		}

		if m.focus.Is(focus.Appointment) {
			switch key {
			case "a":
				m.modalReturnFocus = focus.Appointment
				m.focus.Set(focus.TodoModal)
				return m, m.entrymodal.OpenAppointmentAdd()
			case "r":
				selected, ok := m.appointment.SelectedAppointment()
				if ok {
					m.modalReturnFocus = focus.Appointment
					m.focus.Set(focus.TodoModal)
					return m, m.entrymodal.OpenAppointmentEdit(selected)
				}
			}
		}
	}

	if _, ok := msg.(todolist.OpenTodoModalMsg); ok {
		m.modalReturnFocus = focus.TodoList
		m.focus.Set(focus.TodoModal)
		return m, m.entrymodal.OpenTodoAdd()
	}

	return m, nil
}

func (m *Model) setWindowSize(msg tea.Msg) {
	if windowSizeMsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = windowSizeMsg.Width
		m.height = windowSizeMsg.Height
	}
}
