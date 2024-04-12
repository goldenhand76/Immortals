package database

import (
	"sync/atomic"
)

type Db interface {
	Read()
	// Commit catches the node data; It doesnt save data until the save called.
	Write()
	// Checks if any error or duplicity exists in the data
}

type db struct {
	lastCommit atomic.Value // time.Time - the last time a Node was successfully commited
	options    DbOptions
}

// Connect implements Db.
func (db *db) Read() {
	panic("unimplemented")
}

// Commit implements Db.
func (db *db) Write() {
	panic("unimplemented")
}

func NewClient(o *DbOptions) Db {
	db := &db{}
	db.options = *o
	return db
}
