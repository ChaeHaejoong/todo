package todolist

type Model struct {
}

func New() Model {
	return Model{}
}

func (m Model) View(focused bool) string {
	return renderTodoList("todo list", focused)
}
