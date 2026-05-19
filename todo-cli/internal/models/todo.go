package models

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTodo(title string) *Todo {
	return &Todo{
		Title:     title,
		Completed: false,
	}
}
