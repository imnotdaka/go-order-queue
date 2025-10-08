package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/go-order-queue/internal/order"
	"github.com/imnotdaka/go-order-queue/internal/producer"
)

func OrderHandler(queue *order.MessageQueue, producer *producer.Producer, stop <-chan bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order order.Order
		err := c.ShouldBindJSON(&order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		order.CreatedAt = time.Now()

		producer.Produce(&order, stop)

		c.JSON(http.StatusOK, gin.H{"message": "order created"})
	}
}
