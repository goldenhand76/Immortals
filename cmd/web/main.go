package main

import (
	db "Immortals/internal/database/sqlite/sqlc"
	"Immortals/internal/immo"
	"Immortals/internal/kafka"
	"Immortals/internal/mqtt"
	"context"
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://leo:Goldenhand76@localhost:5432/immo?sslmode=disable"
)

func main() {
	store := &kafka.NotificationStore{
		Data: make(kafka.UserNotifications),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Printf("cannot Connect to database: %s", err)
	}

	dbStore := db.NewStore(conn)

	// Create a wait group to wait for all consumers to finish
	var wg sync.WaitGroup
	wg.Add(3) // Number of consumers

	// Start Immo CLI service
	go immo.SetupImmo(dbStore)
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
