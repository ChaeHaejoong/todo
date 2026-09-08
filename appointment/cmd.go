package appointment

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/store"
)

func loadAppointmentsCmd() tea.Cmd {
	return func() tea.Msg {
		data, err := store.LoadTodos()
		if err != nil {
			return AppointmentLoadFailedMsg{err: err}
		}
		return AppointmentLoadedMsg{Data: data}
	}
}

func removeAppointmentCmd(selected store.Appointment) tea.Cmd {
	return func() tea.Msg {
		if err := store.RemoveAppointment(selected); err != nil {
			return AppointmentRemoveFailedMsg{err: err}
		}
		return AppointmentRemovedMsg{}
	}
}

func (m Model) Init() tea.Cmd {
	return loadAppointmentsCmd()
}
