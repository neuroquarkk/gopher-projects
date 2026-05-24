package main

import "sync"

type Broker struct {
	subscribers map[string][]chan string
	mu          sync.RWMutex
	shutdown    sync.Once
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]chan string),
		mu:          sync.RWMutex{},
		shutdown:    sync.Once{},
	}
}

func (b *Broker) Subscribe(topic string) chan string {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan string)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

func (b *Broker) Publish(topic, message string) {
	b.mu.RLock()
	subs := b.subscribers[topic]
	chanCopy := make([]chan string, len(subs))
	copy(chanCopy, subs)
	b.mu.RUnlock()

	for _, ch := range chanCopy {
		ch <- message
	}
}

func (b *Broker) Shutdown() {
	b.shutdown.Do(func() {
		b.mu.Lock()
		defer b.mu.Unlock()

		for _, v := range b.subscribers {
			for _, ch := range v {
				close(ch)
			}
		}
	})
}
