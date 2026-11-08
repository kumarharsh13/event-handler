package models

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
	UserID      int64     `json:"user_id"`
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

func (e *Event) Update(id int64) error {
	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3
		WHERE id = $4
		RETURNING id
  `

	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, id).Scan(&e.ID)

	return err
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	return err
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events WHERE id = $1`
	var e Event
	err := db.DB.QueryRow(query, id).Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (e *Event) Register(userID int64) error {
	query := `
		INSERT INTO registration (event_id, user_id)
		VALUES ($1, $2)
		RETURNING id
	`
	err := db.DB.QueryRow(query, e.ID, userID).Scan(&e.ID)
	return err
}

func (e *Event) CancelRegistration(eventID int64, userID int64) error {
	query := `
		DELETE FROM registration WHERE event_id = $1 AND user_id = $2
	`
	_, err := db.DB.Exec(query, eventID, userID)
	return err
}
