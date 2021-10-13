package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // package for DB connection driver
)

// 		Opening connection to DB with postgres driver
//		and check if error
func Open() *sql.DB {
	psql, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return psql
}
