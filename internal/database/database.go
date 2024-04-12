package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"Immortals/pkg/models"
)

var ErrNoNodeFound = errors.New("no messages found")

type Db interface {
	Read(nodeId string) (map[string]interface{}, error)
	// Commit catches the node data; It doesnt save data until the save called.
	ReadAll() ([]map[string]interface{}, error)

	Write(*models.NodeData) error
	// Checks if any error or duplicity exists in the data
}

type db struct {
	// lastCommit atomic.Value // time.Time - the last time a Node was successfully commited
	options DbOptions
}

// ReadAll implements Db.
func (db *db) ReadAll() ([]map[string]interface{}, error) {
	// Read JSON data from file
	data, err := os.ReadFile(db.options.filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a map
	var jsonData []map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

// Connect implements Db.
func (db *db) Read(nodeId string) (map[string]interface{}, error) {
	jsonData, err := db.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, item := range jsonData {
		if value, ok := item["nodeId"]; ok && value == nodeId {
			return item, nil
		}
	}
	return nil, ErrNoNodeFound
}

// Commit implements Db.
func (db *db) Write(nodeData *models.NodeData) error {
	// Marshal the JSON data
	jsonBytes, err := json.MarshalIndent(nodeData, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	// Write the JSON data to the file
	if err := os.WriteFile(db.options.filePath, jsonBytes, 0644); err != nil {
		return err
	}

	fmt.Printf("Node data written to %s", db.options.filePath)
	return nil
}

func NewClient(o *DbOptions) Db {
	db := &db{}
	db.options = *o
	return db
}
