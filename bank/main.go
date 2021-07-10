package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db := setupDb()
	defer db.Close()
}

func setupDb() *sql.DB {
	pgCredentials := fmt.Sprintf("host=%s port=%s user=%s posword=%s dbname=%s sslmode=disable",
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_DBNAME"),
	)

	db, err := sql.Open("postgres", pgCredentials)

	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}
