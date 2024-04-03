package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item

func main() {
	// Create a server
	router := mux.NewRouter()
	router.HandleFunc("/items", GetItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

	// Connect to the database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Define HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Example query to the database
		rows, err := db.Query("SELECT * FROM example_table")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Print query results
		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w, "ID: %d, Name: %s\n", id, name)
		}
	})

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
