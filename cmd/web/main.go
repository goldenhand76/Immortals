package main

import (
	"fmt"
	"net/http"
	"os"

	_ "database/sql"
	_ "encoding/json"
	_ "log"

	"github.com/IBM/sarama"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/mdns"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}

func main() {
	// mqtt topic
	topic := "test/"

	opts := mqtt.NewClientOptions()
	opts.AddBroker("localhost:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to MQTT broker:", token.Error()))
	}

	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error subscribing to topic:", token.Error()))
	}
	fmt.Println("Subscribed to topic:", topic)

	// Setup our service export
	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	service, _ := mdns.NewMDNSService(host, "MQQTT.tcp", "", "", 1883, nil, info)
	println(host)
	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()

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
		if isBrokerAvailable(kafkaBroker.Address) {
			fmt.Println("Kafka broker is available.")
			c.JSON(http.StatusOK, gin.H{"status": "running"})
		} else {
			fmt.Println("Kafka broker is not available.")
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
		}
	})
	// Run the server
	router.Run(":8088")
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
