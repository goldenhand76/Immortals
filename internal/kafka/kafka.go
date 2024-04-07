package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

func IsBrokerAvailable(broker string) bool {
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
