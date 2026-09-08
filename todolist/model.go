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

func (m Model) View(width, height int, focused bool) string {
	if m.loading {
		return renderTodoList(width, height, "fetching data...", focused)
	}

	if m.err != nil {
		return renderTodoList(
			width,
			height,
			"Todo를 불러오지 못했습니다: "+m.err.Error(),
			focused,
		)
	}

	return renderTodoList(width, height, renderTodos(width, height, m.todos, m.cursor), focused)
}

func (m Model) selectedTodo() (store.Todo, bool) {
	if m.cursor < 0 || m.cursor >= len(m.todos) {
		return store.Todo{}, false
	}

	return m.todos[m.cursor], true
}

func (m Model) SelectedTodo() (store.Todo, bool) {
	return m.selectedTodo()
}
