package todolist

import "github.com/chaehaejoong/todo/store"

type Model struct {
	todos   []store.Todo
	cursor  int
	loading bool
	err     error
}

func New() Model {
	return Model{
		loading: true,
	}
}

func (m Model) View(focused bool) string {
	if m.loading {
		return renderTodoList("fetching data...", focused)
	}

	if m.err != nil {
		return renderTodoList(
			"Todo를 불러오지 못했습니다: "+m.err.Error(),
			focused,
		)
	}

	return renderTodoList(renderTodos(m.todos, m.cursor), focused)
}
