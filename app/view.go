package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/internal/focus"
)

type AppViews struct {
	TodoList  string
	TodoModal string
	Calendar  string
}

func (m Model) View() tea.View {
	views := AppViews{
		TodoList:  m.todolist.View(m.focus.Is(focus.TodoList)),
		TodoModal: m.todomodal.View(m.focus.Is(focus.TodoModal)),
		Calendar:  m.calendar.View(m.focus.Is(focus.Calendar)),
	}

	view := tea.NewView(renderApp(m.width, m.height, views))
	view.AltScreen = true

	return view
}
