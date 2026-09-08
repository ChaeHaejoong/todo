package appointment

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/entrymodal"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AppointmentLoadedMsg:
		m.appointments = msg.Data.Appointment
		m.cursor = min(m.cursor, max(0, len(m.appointments)-1))
		m.loading = false
		m.err = nil
		return m, nil
	case AppointmentLoadFailedMsg:
		m.loading = false
		m.err = msg.err
		return m, nil
	case entrymodal.AppointmentSubmittedMsg:
		return m, loadAppointmentsCmd()
	case AppointmentRemovedMsg:
		return m, loadAppointmentsCmd()
	case AppointmentRemoveFailedMsg:
		m.err = msg.err
		return m, nil
	case tea.KeyPressMsg:
		if isDeleteKey(msg.String()) {
			selected, ok := m.selectedAppointment()
			if !ok {
				return m, nil
			}
			return m, removeAppointmentCmd(selected)
		}
		if isCursorDownKey(msg.String()) {
			m.cursor = min(len(m.appointments)-1, m.cursor+1)
			return m, nil
		}
		if isCursorUpKey(msg.String()) {
			m.cursor = max(0, m.cursor-1)
			return m, nil
		}
	}

	return m, nil
}
