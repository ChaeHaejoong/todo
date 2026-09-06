package app

import (
	"charm.land/lipgloss/v2"
	"github.com/chaehaejoong/todo/internal/focus"
)

type AppViews struct {
	TodoList    string
	TodoModal   string
	Calendar    string
	Clock       string
	Appointment string
}

func (m Model) renderApp() string {
	appWidth := m.width
	appHeight := m.height

	calendarWidth := 39
	calendarHeight := 11
	clockWidth := 18
	clockHeight := appHeight - calendarHeight
	rightWidth := max(calendarWidth, clockWidth)

	appointmentWidth := 39
	appointmentHeight := appHeight
	todoListWidth := appWidth - appointmentWidth - rightWidth
	todoListHeight := appHeight

	modalWidth := 40
	modalHeight := 3

	views := AppViews{
		TodoList:    m.todolist.View(todoListWidth, todoListHeight, m.focus.Is(focus.TodoList)),
		TodoModal:   m.todomodal.View(modalWidth, modalHeight, m.focus.Is(focus.TodoModal)),
		Calendar:    m.calendar.View(rightWidth, calendarHeight, m.focus.Is(focus.Calendar)),
		Clock:       m.clock.View(rightWidth, clockHeight, m.focus.Is(focus.Clock)),
		Appointment: m.appointment.View(appointmentWidth, appointmentHeight, m.focus.Is(focus.Appointment)),
	}

	rightContent := lipgloss.JoinVertical(
		lipgloss.Left,
		views.Calendar,
		views.Clock,
	)

	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		views.TodoList,
		views.Appointment,
		rightContent,
	)

	background := lipgloss.Place(
		appWidth,
		appHeight,
		lipgloss.Left,
		lipgloss.Top,
		mainContent,
	)

	layers := []*lipgloss.Layer{
		lipgloss.NewLayer(background).
			X(0).
			Y(0).
			Z(0),
	}

	if views.TodoModal != "" {
		modalWidth := lipgloss.Width(views.TodoModal)
		modalHeight := lipgloss.Height(views.TodoModal)

		modalX := max(0, (appWidth-modalWidth)/2)
		modalY := max(0, (appHeight-modalHeight)/2)

		layers = append(
			layers,
			lipgloss.NewLayer(views.TodoModal).
				X(modalX).
				Y(modalY).
				Z(1),
		)
	}

	content := lipgloss.NewCompositor(layers...).Render()
	return content
}
