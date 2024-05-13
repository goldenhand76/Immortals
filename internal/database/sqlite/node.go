package database

import (
	"Immortals/pkg/models"
	"log"

	"github.com/pkg/errors"
)

type Node interface {
	ReadNode(nodeId string) (*models.Node, error)
	ReadAllNodes() ([]models.Node, error)
	InsertNode(*models.Node) error
}

func (d *db) ReadAllNodes() ([]models.Node, error) {
	rows, err := d.conn.Query("SELECT id, name, clientID FROM nodes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []models.Node

	for rows.Next() {
		var node models.Node
		err := rows.Scan(&node.Id, &node.Name, &node.ClientID)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	d.nodes = nodes
	return nodes, nil
}

// Insert implements DbContext.
func (d *db) InsertNode(node *models.Node) error {
	tx, err := d.conn.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	log.Print("Connection Successfully Begun")

	stmt, err := tx.Prepare("INSERT INTO nodes (name, clientID) VALUES (?, ?)")
	if err != nil {
		return errors.Wrap(err, "Failed to prepare statement")
	}
	log.Print("Statement Prepared Successfully")

	_, err = stmt.Exec(node.Name, node.ClientID)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to insert node")
	}
	log.Print("Node Inserted Successfully")

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	log.Print("Commited.")

	return nil
}

// Read implements DbContext.
func (d *db) ReadNode(nodeId string) (*models.Node, error) {
	panic("unimplemented")
}
