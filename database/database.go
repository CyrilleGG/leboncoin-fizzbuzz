package database

import (
	"log"
	"os"
	"database/sql"        // package for SQL DB operations
	_ "github.com/lib/pq" // package for DB connection driver

	"github.com/joho/godotenv" // package for interacting with .env files
)

// 		Fetching connection info for DB
var err = godotenv.Load(".env")
var dbHost, _ = os.LookupEnv("DB_HOST")
var dbName, _ = os.LookupEnv("DB_DATABASE")
var user, _ = os.LookupEnv("DB_USER")
var pwd, _ = os.LookupEnv("DB_PASSWORD")

// 		Declaring config connection for DB
var connStr = "host=" + dbHost + " port=5432 dbname=" + dbName + " user=" + user + " password=" + pwd + " sslmode=require"

// 		Opening connection to DB with postgres driver
//		and check if error
func Open() *sql.DB {
	if err != nil {
		log.Fatal(".env file not found...")
		return nil
	}
	psql, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return psql
}
