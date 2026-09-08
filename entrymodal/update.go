package entrymodal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/store"
	"github.com/chaehaejoong/todo/todo"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if keyPressMsg, ok := msg.(tea.KeyPressMsg); ok {
		key := keyPressMsg.String()

		if isToggleTimeKey(key) {
			m.includeTime = !m.includeTime
			return m, nil
		}

		if isCloseModalKey(key) {
			m.CloseModal()
			return m, nil
		}

		if isAppendTodoKey(key) {
			var err error
			switch m.mode {
			case TodoEdit:
				err = store.UpdateTodo(*m.todo, m.input.Value(), *m.appTime, m.includeTime)
			case AppointmentAdd:
				err = store.AppendAppointment(m.input.Value(), *m.appTime, m.includeTime)
			case AppointmentEdit:
				err = store.UpdateAppointment(*m.appointment, m.input.Value(), *m.appTime, m.includeTime)
			default:
				err = AppendTodo(m.input.Value(), *m.appTime, m.includeTime)
			}

			if err != nil {
				return m, nil
			}
			m.CloseModal()
			m.input.SetValue("")
			return m, func() tea.Msg {
				switch m.mode {
				case AppointmentAdd, AppointmentEdit:
					return AppointmentSubmittedMsg{}
				default:
					return todo.TodoSubmittedMsg{}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
