package sqlc

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Discover(Name string, ClientID string) (Nodes, error) {

	url := "http://" + ClientID + "/ip"
	jsonStr := []byte(`{"brokerIP": "192.168.0.26"}`)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return Nodes{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return Nodes{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error", err)
			return Nodes{}, err
		}
		var data Nodes
		data.ClientID = sql.NullString{String: ClientID, Valid: true}
		data.Name = sql.NullString{String: Name, Valid: true}
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Error:", err)
			return Nodes{}, err
		}
		return data, nil
	}
	return Nodes{}, nil
}
