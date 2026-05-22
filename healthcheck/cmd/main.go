package main

import (
	"fmt"
	"healthcheck/internal"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <urls.txt>")
		return
	}

	arg := os.Args[1]

	if arg == "" {
		fmt.Println("Usage: go run main.go <urls.txt>")
		return
	}

	input, err := internal.ReadFile(arg)
	if err != nil {
		fmt.Printf("[ERROR] failed to read file: %v\n", err)
		return
	}

	pipeline := internal.NewPipeline(input)

	for result := range pipeline.Run() {
		fmt.Printf("%-40s - %-20s (%v)\n", result.Url, result.Status, result.Time)
	}
}
