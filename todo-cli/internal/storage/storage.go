package storage

import (
	"encoding/json"
	"errors"
	"os"
	"todo-cli/internal/models"
)

const filepath string = "data/todo.json"

var todos []*models.Todo

func AddTodo(title string) {
	todo := models.NewTodo(title)
	todos = append(todos, todo)
}

func GetTodos() []*models.Todo {
	return todos
}

func ToggleTodo(id int) error {
	if id > len(todos) {
		return errors.New("id is out of range")
	}
	idx := id - 1
	todo := todos[idx]
	todo.Completed = !todo.Completed
	return nil
}

func DeleteTodo(id int) error {
	if id > len(todos) {
		return errors.New("id is out of range")
	}
	idx := id - 1
	todos = append(todos[:idx], todos[idx+1:]...)
	return nil
}

func LoadTodo() error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return errors.New("failed to read the file data")
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		return errors.New("failed to unmarshal the json data")
	}

	if err := validateTodos(); err != nil {
		return err
	}

	return nil
}

func SaveTodo() error {
	b, err := json.Marshal(todos)
	if err != nil {
		return errors.New("failed to marshal the data")
	}

	if err := os.WriteFile(filepath, b, 0644); err != nil {
		return errors.New("failed to write data")
	}

	return nil
}

func validateTodos() error {
	for _, todo := range todos {
		if todo.Title == "" {
			return errors.New("corrupted json file, missing title field")
		}
	}

	return nil
}
