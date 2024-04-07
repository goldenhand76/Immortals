package node

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"time"
)

var ErrNoNodeFound = errors.New("no messages found")

func DiscoverNode(deviceID string) error {
	url := "http://" + deviceID + "/ip"
	jsonStr := []byte(`{"brokerIP": "192.168.0.26"}`)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
