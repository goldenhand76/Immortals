package database

import (
	"encoding/json"
	"os"
)

type Db interface {
	Read() (map[string]interface{}, error)
	// Commit catches the node data; It doesnt save data until the save called.
	Write(map[string]interface{}) error
	// Checks if any error or duplicity exists in the data
}

type db struct {
	// lastCommit atomic.Value // time.Time - the last time a Node was successfully commited
	options DbOptions
}

// Connect implements Db.
func (db *db) Read() (map[string]interface{}, error) {
	// Read JSON data from file
	data, err := os.ReadFile(db.options.filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a map
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

// Commit implements Db.
func (db *db) Write(jsonData map[string]interface{}) error {
	// Marshal the JSON data
	jsonBytes, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	if err := os.WriteFile(db.options.filePath, jsonBytes, 0644); err != nil {
		return err
	}

	return nil
}

func NewClient(o *DbOptions) Db {
	db := &db{}
	db.options = *o
	return db
}
