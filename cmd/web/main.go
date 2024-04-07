package main

import (
	"Immortals/internal/kafka"
	"Immortals/internal/mqtt"
	"Immortals/pkg/api/node"
	"context"
)

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

	// Discover nodes using QR code
	node.DiscoverNode("esp32.local")
}
