package main

import (
	"Immortals/internal/immo"
	"Immortals/internal/kafka"
	"Immortals/internal/mqtt"
	_ "Immortals/pkg/api/node"
	"context"
	"sync"
)

func main() {
	// err := node.DiscoverNode("esp32.local")
	// if err != nil {
	// 	log.Printf("Failed to discover the node %v", err)
	// }

	store := &kafka.NotificationStore{
		Data: make(kafka.UserNotifications),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a wait group to wait for all consumers to finish
	var wg sync.WaitGroup
	wg.Add(3) // Number of consumers

	go immo.SetupImmo()
	defer wg.Done()

	// Start Kafka consumer
	go kafka.SetupConsumerGroup(ctx, store)
	defer wg.Done()

	// Start MQTT consumer
	go mqtt.SetupMQTTConsumer(store)
	defer wg.Done()

	wg.Wait()
	// Discover nodes using QR code
}
