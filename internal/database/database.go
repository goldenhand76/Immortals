package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"Immortals/pkg/models"
)

var ErrNoNodeFound = errors.New("no messages found")

type DbContext interface {
	Read(nodeId string) (*models.NodeData, error)
	// Commit catches the node data; It doesnt save data until the save called.
	ReadAll() ([]models.NodeData, error)

	Write(*[]models.NodeData) error
	// Checks if any error or duplicity exists in the data

	Append(*models.NodeData) error

	NotExists(nodeId string) bool
}

type db struct {
	nodes []models.NodeData
	// lastCommit atomic.Value // time.Time - the last time a Node was successfully commited
	options DbOptions
}

// ReadAll implements Db.
func (db *db) ReadAll() ([]models.NodeData, error) {
	// Read JSON data from file
	data, err := os.ReadFile(db.options.filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a map
	var jsonData []models.NodeData
	if err := json.Unmarshal([]byte(data), &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

// Connect implements Db.
func (db *db) Read(nodeId string) (*models.NodeData, error) {
	for _, item := range db.nodes {
		if item.NodeID == nodeId {
			return &item, nil
		}
	}
	return nil, ErrNoNodeFound
}

// Commit implements Db.
func (db *db) Write(nodeData *[]models.NodeData) error {
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

// NotExists implements Db.
func (db *db) NotExists(nodeId string) bool {
	for _, item := range db.nodes {
		if item.NodeID == nodeId {
			return false
		}
	}
	return true
}

// WriteMany implements Db.
func (db *db) Append(nodeData *models.NodeData) error {
	db.nodes = append(db.nodes, *nodeData)

	if err := db.Write(&db.nodes); err != nil {
		return err
	}
	return nil
}

func NewDbContext(o *DbOptions) DbContext {
	db := &db{}
	db.options = *o
	nodes, err := (*db).ReadAll()
	if err != nil {
		fmt.Println("Error Retrieving nodes:", err)
	}
	db.nodes = nodes
	return db
}
