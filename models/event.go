package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

// Save saves event in a slice of struct named events
func (e Event) Save() {
	events = append(events, e)
}

// GetAllEvents returns all events
func GetAllEvents() []Event {
	return events
}
