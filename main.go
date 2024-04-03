package main

import (
	_ "database/sql"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	_ "log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
)

func main() {
	router := gin.Default()

	// Endpoint to submit Kafka broker address via POST request
	router.POST("/kafka-setup", func(c *gin.Context) {
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
		if isBrokerAvailable(kafkaBroker.Address) {
			fmt.Println("Kafka broker is available.")
			c.JSON(http.StatusOK, gin.H{"status": "running"})
		} else {
			fmt.Println("Kafka broker is not available.")
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
		}
	})
	// Run the server
	router.Run(":3312")
}

func isBrokerAvailable(broker string) bool {
	// Create configuration for the Kafka client
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0 // Set the Kafka version

	// Create a new Kafka client
	client, err := sarama.NewClient([]string{broker}, config)
	if err != nil {
		fmt.Println("Error creating Kafka client:", err)
		return false
	}
	defer client.Close()

	return true
}
