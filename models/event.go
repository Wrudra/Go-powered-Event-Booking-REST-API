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
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

/*
	This variable will store a slice of events
*/
var events = []Event{}

/*
	Save such an event to the database later,
	for the moment simply into a variable
*/
func (e Event) Save() {
	// later: add it to a database
	events = append(events, e)
}
