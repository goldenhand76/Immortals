package main

import (
	db "Immortals/internal/database"
	"Immortals/internal/immo"
	"Immortals/internal/kafka"
	"Immortals/internal/mqtt"
	_ "Immortals/pkg/node"
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
	wg.Add(3) // Number of consumers

	opts := db.NewDbOptions()
	db := db.NewDbContext(opts)

	// Start Immo CLI service
	go immo.SetupImmo(db)
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
