package pipeline

import (
	"bufio"
	"os"
	"strings"
)

func WordsFromArgs(args []string) []string {
	return args
}

func WordsFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var words []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words = append(words, line)
		}
	}
	return words, nil
}
