package node

import (
	"bytes"
	"errors"
	"log"
	"net/http"
)

var ErrNoNodeFound = errors.New("no messages found")

func DiscoverNode(deviceID string) {
	url := "http://" + deviceID + "/ip"
	jsonStr := []byte(`{"brokerIP": "192.168.0.26"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("failed to unmarshal notification: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed to unmarshal notification: %v", err)
	}
	defer resp.Body.Close()
}
