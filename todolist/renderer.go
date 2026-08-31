package todolist

import (
	"strings"

	"github.com/chaehaejoong/todo/internal/ui"
	"github.com/chaehaejoong/todo/store"
)

func renderTodoList(content string, focused bool) string {
	return ui.TitledBorder(100, 10, "[1] Todo", content, focused)
}

func renderTodos(todos []store.Todo, cursor int) string {
	if len(todos) == 0 {
		return "Nothing to do"
	}

	lines := make([]string, 0, len(todos))
	for index, todo := range todos {
		content := "- " + todo.Content
		lines = append(lines, ui.ListElement(content, index == cursor))
	}

	return strings.Join(lines, "\n")
}
