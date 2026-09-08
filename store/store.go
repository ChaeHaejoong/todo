package store

import (
	"github.com/chaehaejoong/todo/internal/apptime"
	"github.com/google/uuid"
)

func LoadTodos() (Data, error) {
	return load()
}

func AppendTodo(todoStr string, appTime apptime.Apptime, includeTime bool) error {
	todo := Todo{
		ID:      uuid.NewString(),
		Content: todoStr,
		Date:    appTime.DateString(),
	}
	if includeTime {
		todo.Time = appTime.EndTimeString()
	}

	data, err := load()
	if err != nil {
		return err
	}

	data.Todos = append(data.Todos, todo)

	return save(data)
}

func AppendAppointment(content string, appTime apptime.Apptime) error {
	appointment := Appointment{
		ID:        uuid.NewString(),
		Content:   content,
		Date:      appTime.DateString(),
		StartTime: appTime.StartTimeString(),
		EndTime:   appTime.EndTimeString(),
	}

	data, err := load()
	if err != nil {
		return err
	}

	data.Appointment = append(data.Appointment, appointment)

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
