package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/go-order-queue/internal/handlers"
	"github.com/imnotdaka/go-order-queue/internal/order"
	"github.com/imnotdaka/go-order-queue/internal/producer"
)

func main() {
	router := gin.Default()

	mq := order.NewMessageQueue("orders", 5)
	producer := producer.NewProducer(mq)
	stop := make(chan bool)

	router.POST("/order", handlers.OrderHandler(mq, producer, stop))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run()
}
