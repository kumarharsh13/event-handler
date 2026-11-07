package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error
	dbURL := os.Getenv("DB_URL")
	DB, err = sql.Open("pgx", dbURL)

	if err != nil {
		panic("Couldn't open database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping() // actually tries to connect
	if err != nil {
		panic("Couldn't connect to database: " + err.Error())
	}

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Couldn't connect to database: " + err.Error())
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    date_time TIMESTAMP NOT NULL,
    user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Couldn't connect to database: " + err.Error())
	}
}
