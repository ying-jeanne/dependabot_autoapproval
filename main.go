package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// SQLite Connection
	sqliteDB, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer sqliteDB.Close()

	// Create SQLite table
	_, err = sqliteDB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			age INTEGER
		)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("SQLite table created successfully.")

	// PostgreSQL Connection
	pgDB, err := sql.Open("postgres", "host=localhost port=5432 user=youruser password=yourpassword dbname=yourdbname sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer pgDB.Close()

	// Create PostgreSQL table
	_, err = pgDB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			age INTEGER
		)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("PostgreSQL table created successfully.")
	fmt.Println("test test.")
}
