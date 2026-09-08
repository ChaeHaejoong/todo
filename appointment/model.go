package appointment

import "github.com/chaehaejoong/todo/store"

type Model struct {
	appointments []store.Appointment
	cursor       int
	loading      bool
	err          error
}

func New() Model {
	return Model{loading: true}
}

func (m Model) selectedAppointment() (store.Appointment, bool) {
	if m.cursor < 0 || m.cursor >= len(m.appointments) {
		return store.Appointment{}, false
	}
	return m.appointments[m.cursor], true
}

func (m Model) SelectedAppointment() (store.Appointment, bool) {
	return m.selectedAppointment()
}
