package models

import (
	"time"

	"github.com/Soyaib10/eba-event-booking-api/db"
)

// This file is responsible of sql qurey making
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

// Save saves event in a slice of struct named events
func (e Event) Save() error {
	query := 
	`INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // Exac() used when to update or change data
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// GetAllEvents returns all events
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // Query() used to get bunch of rows back
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}


func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id) // QueryRow for exactly one row

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) Update() error {
	query := 
	`UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

    stmt, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(e.ID)
    if err != nil {
        return err
    }
	return nil
}