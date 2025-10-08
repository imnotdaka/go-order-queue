package producer

import (
	"log"

	"github.com/imnotdaka/go-order-queue/internal/order"
)

type Producer struct {
	queue *order.MessageQueue
}

func NewProducer(queue *order.MessageQueue) *Producer {
	return &Producer{
		queue: queue,
	}
}

func (p *Producer) Produce(order *order.Order, stop <-chan bool) {
	select {
	case <-stop:
		log.Println("stopping")
		return
	default:
		p.queue.Publish(order, stop)
	}
}
