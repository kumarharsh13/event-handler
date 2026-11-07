package event

import (
	"event-handler/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
        INSERT INTO events (name, description, location, date_time, user_id)
        VALUES($1, $2, $3, $4, $5)
        RETURNING id
    `
	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)

	return err
}

func GetAllEvents() []Event {
	return events
}
