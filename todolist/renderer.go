package todolist

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/chaehaejoong/todo/internal/ui"
	"github.com/chaehaejoong/todo/store"
)

func renderTodoList(width, height int, content string, focused bool) string {
	return ui.TitledBorder(width, height, "[1] Todo", content, focused)
}

func renderTodos(width, height int, todos []store.Todo, cursor int) string {
	if len(todos) == 0 {
		return "Nothing to do"
	}

	contentWidth, _ := ui.TitledBorderContentSize(width, height)
	lines := make([]string, 0, len(todos))
	for index, todo := range todos {
		content := renderTodo(contentWidth, todo)
		lines = append(lines, ui.ListElement(content, index == cursor))
	}

	return strings.Join(lines, "\n")
}

func renderTodo(width int, todo store.Todo) string {
	content := "- " + todo.Content
	metadata := strings.TrimSpace(fmt.Sprintf("%s %s", todo.Date, todo.Time))
	if metadata == "" {
		return content
	}

	contentWidth := lipgloss.Width(content)
	metadataWidth := lipgloss.Width(metadata)
	separatorWidth := width - contentWidth - metadataWidth - 3
	if separatorWidth >= 1 {
		return content + " " + strings.Repeat("-", separatorWidth) + " ~" + metadata
	}

	metadataLine := strings.Repeat("-", max(1, width-metadataWidth)) + metadata
	return content + "\n" + metadataLine
}
