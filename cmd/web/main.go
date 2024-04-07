package main

import (
	"Immortals/internal/kafka"
	"Immortals/internal/mqtt"
	"context"
	"sync"
)

func main() {

	store := &kafka.NotificationStore{
		Data: make(kafka.UserNotifications),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a wait group to wait for all consumers to finish
	var wg sync.WaitGroup
	wg.Add(2) // Number of consumers

	// Start Kafka consumer
	go kafka.SetupConsumerGroup(ctx, store)
	defer wg.Done()

	// Start MQTT consumer
	go mqtt.SetupMQTTConsumer(store)
	defer wg.Done()

	wg.Wait()
	// Discover nodes using QR code

	// node.DiscoverNode("esp32.local")
}
