/*
All the logic that deals with storing event data in a database
*/
package models

import "time"

/*
	Defines the shape of an event
*/
type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

/*
	This variable will store a slice of events
*/
var events = []Event{}

/*
	A method to save events to the database later,
	for the moment simply into a variable
*/
func (e Event) Save() {
	// later: add it to a database
	events = append(events, e)
}

/*
	A func to get all events
*/
func GetAllEvents() []Event {
	return events
}
