package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	services := []string{"AuthService", "DBService", "WebService"}

	var serviceWg sync.WaitGroup
	var workerWg sync.WaitGroup

	metricChan := make(chan Metric, 100)
	alertChan := make(chan string)
	quitChan := make(chan struct{})

	for _, service := range services {
		serviceWg.Add(1)
		go serviceGenerator(service, &serviceWg, metricChan, quitChan)
	}

	for range 3 {
		workerWg.Add(1)
		go worker(metricChan, alertChan, &workerWg)
	}

	go func() {
		for alert := range alertChan {
			fmt.Println(alert)
		}
	}()

	<-time.After(5 * time.Second)
	close(quitChan)

	serviceWg.Wait()
	close(metricChan)

	workerWg.Wait()
	close(alertChan)
}
