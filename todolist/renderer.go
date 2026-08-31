package todolist

import (
	"fmt"
	"strings"

	"github.com/chaehaejoong/todo/internal/ui"
	"github.com/chaehaejoong/todo/store"
)

func renderTodoList(content string, focused bool) string {
	return ui.TitledBorder(100, 10, "[1] Todo", content, focused)
}

func renderTodos(todos []store.Todo) string {
	if len(todos) == 0 {
		return "Nothing to do"
	}

	lines := make([]string, 0, len(todos))
	for _, todo := range todos {
		lines = append(lines, fmt.Sprintf("- %s", todo.Content))
	}

	return strings.Join(lines, "\n")
}
