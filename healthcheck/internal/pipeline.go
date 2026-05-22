package internal

import (
	"sync"
	"time"
)

type Pipeline struct {
	urls []string
}

type Result struct {
	Url    string
	Status string
	Time   time.Duration
}

func NewPipeline(urls []string) *Pipeline {
	return &Pipeline{urls}
}

func (p *Pipeline) Run() <-chan Result {
	return p.request(p.generate())
}

func (p *Pipeline) generate() <-chan string {
	outChan := make(chan string)

	go func() {
		defer close(outChan)
		for _, url := range p.urls {
			outChan <- url
		}
	}()

	return outChan
}

func (p *Pipeline) request(ch <-chan string) <-chan Result {
	merged := make(chan Result)
	const workers uint8 = 10
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range ch {
				status, elapsed := Request(url)
				merged <- Result{Url: url, Status: status, Time: elapsed}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
