package models

type Todo struct {
	Title     string
	Completed bool
}

func NewTodo(title string) *Todo {
	return &Todo{
		Title:     title,
		Completed: false,
	}
}
