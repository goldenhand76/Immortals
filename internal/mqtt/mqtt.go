package mqtt

import (
	"Immortals/internal/kafka"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	MQTTBrokerAddress = "broker:1883"
	MQTTTopic         = "test/"
)

func onMessageReceived(message mqtt.Message, store *kafka.NotificationStore) {
	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}

func SetupMQTTConsumer(store *kafka.NotificationStore) {
	// mqtt topic
	payload := []byte("Bye")
	var qos byte = 1
	opts := mqtt.NewClientOptions()
	opts.SetClientID("Go Client")
	opts.SetBinaryWill("/Go/Will", payload, qos, false)
	opts.AddBroker("broker:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to MQTT broker: %s", token.Error()))
	}

	if token := client.Subscribe(MQTTTopic, 0, func(client mqtt.Client, message mqtt.Message) {
		onMessageReceived(message, store)
	}); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error subscribing to topic: %s", token.Error()))
	}
	fmt.Println("Subscribed to topic:", MQTTTopic)
}
