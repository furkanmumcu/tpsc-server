package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
		return nil, err
	}
	return db, nil
}
