/*
All the logic that deals with storing event data in a database
*/
package models

import (
	"root/db"
	"time"
)

/*
Defines the shape of an event
*/
type Event struct {
	ID          int64
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
A method to save events to the database
Returns error if any error occures
*/
func (e Event) Save() error {
	/*
		INSERT
		Storing data in the database
	*/
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)` // Special Syntax for inserting actual values
	/*
		Inserting the actual values
	*/
	stmt, err := db.DB.Prepare(query)
	/*
		Checking Error
	*/
	if err != nil {
		return err
	}
	defer stmt.Close() // Closing the statement without execution
	/*
		Execute the statment if no error
		Passing the event names for values
	*/
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	/*
		Checking Error
	*/
	if err != nil {
		return err
	}
	/*
		Calling the last inserted ID that was inserted
		Using the automated generated event ID on the event
	*/
	id, err := result.LastInsertId()
	e.ID = id

	return err
}

/*
A func to get all events
*/
func GetAllEvents() []Event {
	return events
}
