package store

import "github.com/google/uuid"

func LoadTodos() (Data, error) {
	return load()
}

func AppendTodo(todoStr string) error {
	todo := Todo{
		ID: uuid.NewString(),
		Content: todoStr,
	}

	data, err := load()
	if err != nil {
		return err
	}

	data.Todos = append(data.Todos, todo)

	return save(data)
}

func RemoveTodo(todo Todo) error {
	data, err := load()
	if err != nil {
		return err
	}

	remaining := data.Todos[:0]
	for _, item := range data.Todos {
		if item.ID != todo.ID {
			remaining = append(remaining, item)
		}
	}
	data.Todos = remaining

	return save(data)
}
