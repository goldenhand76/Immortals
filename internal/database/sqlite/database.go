package database

import (
	"Immortals/pkg/models"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DbContext interface {
	Node
}

type db struct {
	conn    *sql.DB
	nodes   []models.Node
	options DbOptions
}

func NewDbContext(o *DbOptions) (DbContext, error) {
	db := &db{}
	db.options = *o

	tx, err := sql.Open("sqlite3", o.filePath)
	if err != nil {
		return nil, err
	}
	db.conn = tx
	log.Printf("Transaction Successfully started.")

	_, err = db.conn.Exec(SQL_NODES_TABLE)
	if err != nil {
		log.Printf("Error : %s", err)
		return nil, err
	}
	log.Printf("Nodes Table successfully Created.")

	return db, nil
}

var SQL_NODES_TABLE = `CREATE TABLE IF NOT EXISTS nodes (
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	name TEXT, 
	agentID INTEGER, 
	clientID TEXT UNIQUE, 
	is_online INTEGER)`
