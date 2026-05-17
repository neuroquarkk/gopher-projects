package storage

import (
	"errors"
	"todo-cli/internal/models"
)

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
