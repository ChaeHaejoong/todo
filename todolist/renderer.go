package todolist

import (
	"github.com/chaehaejoong/todo/internal/ui"
)

func renderTodoList(content string, focused bool) string {
	return ui.TitledBorder(100, 10, "Todo", content, focused)
}
