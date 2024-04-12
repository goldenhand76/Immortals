package node

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	db "Immortals/internal/database"
	"Immortals/pkg/models"
)

var ErrNoNodeFound = errors.New("no messages found")

func Discover(nodeID string) (*models.NodeData, error) {
	url := "http://" + nodeID + "/ip"
	jsonStr := []byte(`{"brokerIP": "192.168.0.26"}`)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error", err)
			return nil, err
		}
		var data models.NodeData
		data.NodeID = nodeID
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		return &data, nil
	}

	return nil, nil
}

func Add(nodeID string) (*models.NodeData, error) {
	fmt.Println("Adding Node...")
	nodeData, err := Discover(nodeID)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	opts := db.NewDbOptions()
	r := db.NewClient(opts)

	if err := r.Write(nodeData); err != nil {
		return nil, err
	}
	return nodeData, nil
}

func List() error {
	fmt.Println("Listing containers")
	return nil
}
