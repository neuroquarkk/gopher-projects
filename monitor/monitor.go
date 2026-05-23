package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Metric struct {
	name string
	cpu  int
}

func serviceGenerator(
	serviceName string,
	wg *sync.WaitGroup,
	metricChan chan<- Metric,
	quitChan <-chan struct{},
) {
	defer wg.Done()
	for {
		t := rand.Intn(200) + 300
		time.Sleep(time.Duration(t) * time.Millisecond)

		metric := Metric{
			name: serviceName,
			cpu:  rand.Intn(101),
		}

		select {
		case metricChan <- metric:
		case <-quitChan:
			fmt.Printf("Shutting down %s generator...\n", serviceName)
			return
		default:
			fmt.Println("Metric channel is full")
		}
	}
}

func worker(metricChan <-chan Metric, alertChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for metric := range metricChan {
		if metric.cpu > 80 {
			alertChan <- fmt.Sprintf("%s is using more than 80%% cpu", metric.name)
		}
	}
}
