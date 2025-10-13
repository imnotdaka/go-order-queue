package consumer

import (
	"log"

	"github.com/imnotdaka/go-order-queue/internal/order"
)

type Worker struct {
	queue *order.MessageQueue
}

func NewWorker(queue *order.MessageQueue) *Worker {
	return &Worker{
		queue: queue,
	}
}

func (w *Worker) Work(stop <-chan bool) {
	for {
		select {
		case <-stop:
			return
		case order := <-w.queue.Messages:
			log.Println("worker received the message: ", order)
		}
	}
}
