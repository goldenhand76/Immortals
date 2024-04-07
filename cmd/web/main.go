package main

import (
	"Immortals/internal/kafka"
	"Immortals/internal/mqtt"
	"Immortals/pkg/api/node"
	"Immortals/pkg/models"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ConsumerPort = ":8081"

func handleNotifications(ctx *gin.Context, store *kafka.NotificationStore) {
	userID := ctx.Param("userID")
	if userID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No user ID provided"})
		return
	}

	notes := store.Get(userID)
	if len(notes) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message":       "No notifications found for user",
			"notifications": []models.Notification{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"notifications": notes})
}

// ============== HELPER FUNCTIONS ==============
var ErrNoMessagesFound = errors.New("no messages found")

func getUserIDFromRequest(ctx *gin.Context) (string, error) {
	userID := ctx.Param("userID")
	if userID == "" {
		return "", ErrNoMessagesFound
	}
	return userID, nil
}

func main() {

	store := &kafka.NotificationStore{
		Data: make(kafka.UserNotifications),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start Kafka consumer
	go kafka.SetupConsumerGroup(ctx, store)

	// Start MQTT consumer
	go mqtt.SetupMQTTConsumer(store)

	node.DiscoverNode("esp32.local")

	router := gin.Default()
	// Endpoint to submit Kafka broker address via POST request
	router.POST("/kafka-host", func(c *gin.Context) {
		var kafkaBroker struct {
			Address string `json:"address"`
		}
		// Bind JSON body to struct
		if err := c.BindJSON(&kafkaBroker); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			println(err.Error())
			return
		}

		// Check if the broker is available
		if kafka.IsBrokerAvailable(kafkaBroker.Address) {
			fmt.Println("Kafka broker is available.")
			c.JSON(http.StatusOK, gin.H{"status": "running"})
		} else {
			fmt.Println("Kafka broker is not available.")
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
		}
	})

	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		handleNotifications(ctx, store)
	})

	fmt.Printf("Kafka CONSUMER (Group: %s) 👥📥 "+
		"started at http://localhost%s\n", kafka.ConsumerGroup, ConsumerPort)

	if err := router.Run(ConsumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
	// Run the server
	router.Run("0.0.0.0:8088")
}
