package todolist

import (
	"fmt"
	"strings"

	"github.com/chaehaejoong/todo/internal/ui"
	"github.com/chaehaejoong/todo/store"
)

func renderTodoList(content string, focused bool) string {
	return ui.TitledBorder(100, 10, "Todo", content, focused)
}

func renderTodos(todos []store.Todo) string {
	if len(todos) == 0 {
		return "할 일이 없어요"
	}

	lines := make([]string, 0, len(todos))
	for i, todo := range todos {

		lines = append(lines, fmt.Sprintf("%d. %s", i+1, todo.Content))
	}

	return strings.Join(lines, "\n")
}
