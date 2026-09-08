package appointment

import "github.com/chaehaejoong/todo/store"

type AppointmentLoadedMsg struct {
	Data store.Data
}

type AppointmentLoadFailedMsg struct {
	err error
}

type AppointmentRemovedMsg struct{}

type AppointmentRemoveFailedMsg struct {
	err error
}
