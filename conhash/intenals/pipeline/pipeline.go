package pipeline

import (
	"crypto/sha256"
	"fmt"
)

type Result struct {
	Word string
	Hash string
}

type Pipeline struct {
	words []string
}

func New(words []string) *Pipeline {
	return &Pipeline{words}
}

func (p *Pipeline) GenerateWords() <-chan string {
	out := make(chan string)
	go func() {
		for _, word := range p.words {
			out <- word
		}
		close(out)
	}()
	return out
}

func (p *Pipeline) HashWords(inChan <-chan string) <-chan Result {
	out := make(chan Result)
	go func() {
		for word := range inChan {
			sum := sha256.Sum256([]byte(word))
			hashStr := fmt.Sprintf("%x", sum)
			result := Result{Word: word, Hash: hashStr}
			out <- result
		}
		close(out)
	}()
	return out
}

func (p *Pipeline) Run() <-chan Result {
	return p.HashWords(p.GenerateWords())
}
