package node

import (
	"bytes"
	"net/http"
)

func DiscoverNode(deviceID string) bool {
	url := "http://" + deviceID + "/ip"
	jsonStr := []byte(`{"brokerIP": "192.168.0.26"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return true
}
