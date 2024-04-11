package node

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "Immortals/internal/database"
)

var ErrNoNodeFound = errors.New("no messages found")

func Discover(nodeID string) error {
	url := "http://" + nodeID + "/ip"
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
	opts := db.NewDbOptions()
	opts.SetName("Ali")
	r := db.NewClient(opts)
	r.Connect()

	defer resp.Body.Close()
	return nil
}

func List() error {
	fmt.Println("Listing containers")
	return nil
}
