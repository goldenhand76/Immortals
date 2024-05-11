package node

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	db "Immortals/internal/database"
	"Immortals/pkg/models"
)

var ErrNoNodeFound = errors.New("no messages found")
var ErrNodeExists = errors.New("node already exists")

func Discover(nodeName string, nodeID string) (*models.NodeData, error) {
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
		data.NodeName = nodeName
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		return &data, nil
	}

	return nil, nil
}

func Add(db db.DbContext, nodeName string, nodeID string) (*models.NodeData, error) {
	log.Printf("Adding Node...")
	if db.NotExists(nodeID) {
		nodeData, err := Discover(nodeName, nodeID)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		if err := db.Append(nodeData); err != nil {
			return nil, err
		}
		return nodeData, nil
	}
	log.Printf("Node already exists\n")
	return nil, ErrNodeExists
}

func List(db db.DbContext) ([]models.NodeData, error) {
	fmt.Println("Listing nodes...")
	nodes, err := db.ReadAll()
	return nodes, err
}
