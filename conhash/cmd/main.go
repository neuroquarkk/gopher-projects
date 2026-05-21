package main

import (
	"fmt"
	"os"

	"conhash/intenals/pipeline"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run main.go word1 word2 word3...")
		return
	}

	var words []string
	var err error

	if args[0] == "-f" {
		if len(args) < 2 {
			fmt.Println("-f requires a file path")
			return
		}
		words, err = pipeline.WordsFromFile(args[1])
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
	} else {
		words = pipeline.WordsFromArgs(args)
	}

	p := pipeline.New(words)
	for hash := range p.Run() {
		fmt.Printf("%-20s -> %s\n", hash.Word, hash.Hash)
	}
}
