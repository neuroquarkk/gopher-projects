package main

import (
	"flag"
	"fmt"
	"ipcli/internal/api"
	"ipcli/internal/output"
)

func main() {
	var jsonOutput bool

	flag.BoolVar(&jsonOutput, "json", false, "output results in JSON format")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ipcli [--json] <ip|me>")
		return
	}

	query := args[0]
	if query == "me" {
		query = ""
	}

	data, err := api.Request(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	if jsonOutput {
		output.JsonOutput(data)
	} else {
		output.Print(data, query)
	}
}
