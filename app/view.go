package app

import (
	tea "charm.land/bubbletea/v2"
)



func (m Model) View() tea.View {
	view := tea.NewView(m.renderApp())
	view.AltScreen = true

	return view
}
