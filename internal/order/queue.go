package order

import (
	"log"
	"sync"
)

type MessageQueue struct {
	Name     string
	Messages chan *Order
	Mu       sync.Mutex
}

func NewMessageQueue(name string, size int) *MessageQueue {
	return &MessageQueue{
		Name:     name,
		Messages: make(chan *Order, size),
	}
}

func (mq *MessageQueue) Publish(data *Order, stop <-chan bool) {
	select {
	case <-stop:
		log.Println("stopping")
		return
	case mq.Messages <- data:
		log.Println("message sent", mq.Name)
		return
	}
}
