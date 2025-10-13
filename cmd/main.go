package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/go-order-queue/internal/consumer"
	"github.com/imnotdaka/go-order-queue/internal/handlers"
	"github.com/imnotdaka/go-order-queue/internal/order"
	"github.com/imnotdaka/go-order-queue/internal/producer"
)

func main() {
	router := gin.Default()

	mq := order.NewMessageQueue("orders", 5)
	producer := producer.NewProducer(mq)
	stop := make(chan bool)

	worker := consumer.NewWorker(mq)
	go worker.Work(stop)

	router.POST("/order", handlers.OrderHandler(mq, producer, stop))

	router.POST("/stop", handlers.Stop(stop))

	router.Run()
}
