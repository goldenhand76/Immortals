package node

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Node struct {
	Name     string `json:"name"`
	ClientID int    `json:"clientId"`
	Address  string `json:"address"`
}

var ErrNoNodeFound = errors.New("no messages found")

func Discover(deviceID string) error {
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

func List() error {
	fmt.Println("Listing containers")
	return nil
}
