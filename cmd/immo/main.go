package main

import (
	"fmt"

	"Immortals/internal/kafka" // Import the Kafka package
)

func main() {
	// Use the Kafka package to check if the broker is available
	broker := "your_kafka_broker_address"
	if kafka.IsBrokerAvailable(broker) {
		fmt.Println("Kafka broker is available.")
	} else {
		fmt.Println("Kafka broker is not available.")
	}

	// Add your CLI application logic here
}
