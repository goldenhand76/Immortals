package database

import (
	"sync/atomic"
)

type Db interface {
	Connect()
	// Commit catches the node data; It doesnt save data until the save called.
	Commit()
	// Checks if any error or duplicity exists in the data
	IsValid() bool
	// Store data to the json file
	Save() bool
}

type db struct {
	lastCommit atomic.Value // time.Time - the last time a Node was successfully commited
	options    DbOptions
}

// Connect implements Db.
func (db *db) Connect() {
	panic("unimplemented")
}

// Commit implements Db.
func (db *db) Commit() {

}

// IsValid implements Db.
func (db *db) IsValid() bool {
	panic("unimplemented")
}

// Save implements Db.
func (db *db) Save() bool {
	panic("unimplemented")
}

type NodeData struct {
	NodeID   string            `json:"nodeId"`   // Node ID that set on device and not changable
	Sensor   map[string]string `json:"sensor"`   // List of sensors that exists and their topics
	Actuator map[string]string `json:"actuator"` // List of actuators that exists on node and their topics
}

func NewClient(o *DbOptions) Db {
	db := &db{}
	db.options = *o
	return db
}
