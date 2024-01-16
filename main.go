package main

import (
	"database/sql"
	"fmt"
	"log"
	"vhellman/sandbox/scripts"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "viktortest.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(scripts.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// INSERT QUERY - insert some test data
	statement, err := db.Prepare("INSERT INTO your_table_name (entity_id, status) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	/* _, err = statement.Exec("123456789", "pending")
	if err != nil {
		log.Fatal(err)
	} */

	// SELECT QUERY
	// Prepare the select statement
	selectStmt, err := db.Prepare("SELECT * FROM your_table_name WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer selectStmt.Close()

	// Set the id you want to select
	idToSelect := 1 // Replace with the id you want to query

	// Execute the query
	row := selectStmt.QueryRow(idToSelect)

	// Create variables to hold the query results
	var id int
	var entityID, status string

	// Scan the result into the variables
	err = row.Scan(&id, &entityID, &status)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows were returned!")
		} else {
			log.Fatal(err)
		}
	}

	// Print the results
	fmt.Printf("ID: %d, Entity ID: %s, Status: %s\n", id, entityID, status)

}
