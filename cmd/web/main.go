package main

import (
	"bytes"
	"fmt"
	"net/http"

	_ "database/sql"
	_ "encoding/json"
	_ "log"

	"github.com/IBM/sarama"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}

func main() {
	// mqtt topic
	topic := "test/"
	payload := []byte("Bye")
	var qos byte = 1

	opts := mqtt.NewClientOptions()
	opts.SetClientID("Go Client")
	opts.SetBinaryWill("/Go/Will", payload, qos, false)
	opts.AddBroker("broker:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to MQTT broker:", token.Error()))
	}

	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error subscribing to topic:", token.Error()))
	}
	fmt.Println("Subscribed to topic:", topic)

	discoverDevice("esp32.local")

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
	router.Run("0.0.0.0:8088")
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

func discoverDevice(deviceID string) bool {
	url := "http://" + deviceID + "/ip"
	jsonStr := []byte(`{"brokerIP": "192.168.0.26"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return true
}
