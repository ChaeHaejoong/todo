package app

import (
	"charm.land/lipgloss/v2"
	ui "github.com/chaehaejoong/todo/internal/ui"
)

func renderApp(width, height int, views AppViews) string {
	contentWidth, contentHeight := appContentSize(width, height)

	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		views.TodoList,
		views.Calendar,
	)

	background := lipgloss.Place(
		contentWidth,
		contentHeight,
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

		modalX := max(0, (contentWidth-modalWidth)/2)
		modalY := max(0, (contentHeight-modalHeight)/2)

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

func appContentSize(width, height int) (int, int) {
	return ui.TitledBorderContentSize(width, height)
}
