package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "docker"
	password = "XdccDa85_JK"
	dbname   = "docker"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println("----------- Connecting to : " + connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	fmt.Println("----------- Creating table -----------")

	sqlCreate := `CREATE TABLE currencies(id SERIAL PRIMARY KEY, name VARCHAR(3))`
	_, err = db.Exec(sqlCreate)
	if err != nil {
		panic(err)
	}

	fmt.Println("----------- INSERT with binding of variables -----------")

	sqlStatement := `INSERT INTO currencies (name) VALUES ($1)`
	_, err = db.Exec(sqlStatement, "EUR")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(sqlStatement, "CHF")
	_, err = db.Exec(sqlStatement, "USD")
	//ommitting errors here - the first worked

	fmt.Println("----------- QUERY -----------")

	rows, err := db.Query("SELECT id, name FROM currencies LIMIT $1", 10)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, name)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}
