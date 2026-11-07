package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", "postgres://postgres:password@localhost:5432/event-handler")

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
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    date_time TIMESTAMP NOT NULL,
    user_id INTEGER NOT NULL
	)`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Couldn't connect to database: " + err.Error())
	}
}
