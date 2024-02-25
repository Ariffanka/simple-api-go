package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:4kunPostgre@localhost/users?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err !=nil{
		log.Fatal(err)
	}

	// defer db.Close()
	return db
}
