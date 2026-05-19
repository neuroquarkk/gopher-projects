package repl

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-cli/internal/storage"
)

func handleAdd(args []string) {
	title := strings.Join(args, " ")
	title = strings.TrimSpace(title)

	if title == "" {
		fmt.Println("Title cannot be empty")
		return
	}

	storage.AddTodo(title)
	fmt.Println("Todo added successfully")
}

func handleList() {
	todos := storage.GetTodos()
	if len(todos) == 0 {
		fmt.Println("No todo to display")
		return
	}

	for i, todo := range todos {
		var status string
		if todo.Completed {
			status = "[x]"
		} else {
			status = "[ ]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, todo.Title)
	}
}

func handleToggle(args []string) {
	if len(args) == 0 {
		fmt.Println("Id cannot be empty")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid id:", err)
		return
	}

	if id <= 0 {
		fmt.Println("Id should be greater than 0")
		return
	}

	if err := storage.ToggleTodo(id); err != nil {
		fmt.Println("Failed to mark todo:", err)
		return
	}

	fmt.Println("Todo toggled successfully")
}

func handleSave() {
	if err := storage.SaveTodo(); err != nil {
		fmt.Println("Failed to save")
	} else {
		fmt.Println("Saved successfully")
	}
}

func handleDelete(args []string) {
	if len(args) == 0 {
		fmt.Println("Id cannot be empty")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid id:", err)
		return
	}

	if id <= 0 {
		fmt.Println("Id should be greater than 0")
		return
	}

	if err := storage.DeleteTodo(id); err != nil {
		fmt.Println("Failed to delete todo:", err)
		return
	}

	fmt.Println("Tod deleted successfully")
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	if err := storage.LoadTodo(); err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Print("[todo]> ")

		if !scanner.Scan() {
			break
		}

		input := strings.Fields(scanner.Text())
		if len(input) == 0 {
			continue
		}

		cmd := strings.ToLower(input[0])
		args := input[1:]

		switch cmd {
		case "add":
			handleAdd(args)
		case "list":
			handleList()
		case "toggle":
			handleToggle(args)
		case "delete":
			handleDelete(args)
		case "save":
			handleSave()
		case "exit":
			storage.SaveTodo()
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
}
