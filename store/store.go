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

func AppendAppointment(content string, appTime apptime.Apptime, includeTime bool) error {
	appointment := Appointment{
		ID:      uuid.NewString(),
		Content: content,
		Date:    appTime.DateString(),
	}
	if includeTime {
		appointment.StartTime = appTime.StartTimeString()
		appointment.EndTime = appTime.EndTimeString()
	}

	data, err := load()
	if err != nil {
		return err
	}

	data.Appointment = append(data.Appointment, appointment)

	return save(data)
}

func UpdateTodo(todo Todo, content string, appTime apptime.Apptime, includeTime bool) error {
	data, err := load()
	if err != nil {
		return err
	}

	for index := range data.Todos {
		if data.Todos[index].ID != todo.ID {
			continue
		}

		updated := data.Todos[index]
		updated.Content = content
		if includeTime {
			if updated.Time == "" {
				updated.Date = appTime.DateString()
				updated.Time = appTime.EndTimeString()
			}
		} else {
			updated.Time = ""
		}
		data.Todos[index] = updated
		return save(data)
	}

	return nil
}

func UpdateAppointment(appointment Appointment, content string, appTime apptime.Apptime, includeTime bool) error {
	data, err := load()
	if err != nil {
		return err
	}

	for index := range data.Appointment {
		if data.Appointment[index].ID != appointment.ID {
			continue
		}

		updated := data.Appointment[index]
		updated.Content = content
		if includeTime {
			if updated.StartTime == "" && updated.EndTime == "" {
				updated.Date = appTime.DateString()
				updated.StartTime = appTime.StartTimeString()
				updated.EndTime = appTime.EndTimeString()
			}
		} else {
			updated.StartTime = ""
			updated.EndTime = ""
		}
		data.Appointment[index] = updated
		return save(data)
	}

	return nil
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

func RemoveAppointment(appointment Appointment) error {
	data, err := load()
	if err != nil {
		return err
	}

	remaining := data.Appointment[:0]
	for _, item := range data.Appointment {
		if item.ID != appointment.ID {
			remaining = append(remaining, item)
		}
	}
	data.Appointment = remaining

	return save(data)
}
