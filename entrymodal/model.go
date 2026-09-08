package entrymodal

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
	mode        Mode
	todo        *store.Todo
	appointment *store.Appointment
}

type Mode int

const (
	TodoAdd Mode = iota
	TodoEdit
	AppointmentAdd
	AppointmentEdit
)

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
	return renderTodoModal(width, height, m.input.View(), m.includeTime, m.title(), focused)
}

func (m *Model) OpenModal() tea.Cmd {
	return m.OpenTodoAdd()
}

func (m *Model) OpenTodoAdd() tea.Cmd {
	m.mode = TodoAdd
	m.todo = nil
	m.appointment = nil
	m.includeTime = true
	m.input.SetValue("")
	m.isModalOpen = true
	return m.input.Focus()
}

func (m *Model) OpenTodoEdit(todo store.Todo) tea.Cmd {
	m.mode = TodoEdit
	m.todo = &todo
	m.appointment = nil
	m.includeTime = todo.Time != ""
	m.input.SetValue(todo.Content)
	m.isModalOpen = true
	return m.input.Focus()
}

func (m *Model) OpenAppointmentAdd() tea.Cmd {
	m.mode = AppointmentAdd
	m.todo = nil
	m.appointment = nil
	m.includeTime = true
	m.input.SetValue("")
	m.isModalOpen = true
	return m.input.Focus()
}

func (m *Model) OpenAppointmentEdit(appointment store.Appointment) tea.Cmd {
	m.mode = AppointmentEdit
	m.todo = nil
	m.appointment = &appointment
	m.includeTime = appointment.Time != ""
	m.input.SetValue(appointment.Content)
	m.isModalOpen = true
	return m.input.Focus()
}

func (m Model) IsModalOpen() bool {
	return m.isModalOpen
}

func AppendTodo(todoStr string, appTime apptime.Apptime, includeTime bool) error {
	return store.AppendTodo(todoStr, appTime, includeTime)
}

func (m Model) title() string {
	switch m.mode {
	case TodoEdit:
		return "Edit-Todo"
	case AppointmentAdd:
		return "Append-Appointment"
	case AppointmentEdit:
		return "Edit-Appointment"
	default:
		return "Append-Todo"
	}
}

func (m *Model) CloseModal() {
	m.isModalOpen = false
}
