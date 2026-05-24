package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	broker := NewBroker()

	var subWg sync.WaitGroup
	var pubWg sync.WaitGroup
	quitChan := make(chan struct{})

	spawnSub := func(id int, topic string) {
		subWg.Add(1)
		go func() {
			defer subWg.Done()
			myChan := broker.Subscribe(topic)
			for msg := range myChan {
				fmt.Printf("[Sub %d | %s] Received: %s\n", id, topic, msg)
			}
			fmt.Printf("[Sub %d] Channel closed\n", id)
		}()
	}

	spawnSub(1, "logs")
	spawnSub(2, "logs")
	spawnSub(3, "logs")
	spawnSub(4, "alerts")
	spawnSub(5, "alerts")

	for i := range 10 {
		pubWg.Add(1)
		go func(id int) {
			defer pubWg.Done()
			for {
				select {
				case <-quitChan:
					return
				default:
					topic := "logs"
					if rand.Intn(2) == 0 {
						topic = "alerts"
					}

					msg := fmt.Sprintf("System update from publisher %d", id)
					broker.Publish(topic, msg)

					time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
				}
			}
		}(i + 1)
	}

	fmt.Println("Broker running")
	time.Sleep(2 * time.Second)

	fmt.Println("Initiating Graceful Shutdown")

	close(quitChan)
	pubWg.Wait()

	broker.Shutdown()
	subWg.Wait()

	fmt.Println("Broker shutdown complete")
}
