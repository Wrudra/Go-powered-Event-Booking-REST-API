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
		Preparing the query.
		Inserting the actual values, Prepare is efficient as it is stored in the memory of sql package
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
func GetAllEvents() ([]Event, error) {
	/*
		SELECT
		Getting data from the database
	*/
	query := "SELECT * FROM events"
	/*
		Preparing the query.
		Prepare() will work as well
	*/
	rows, err := db.DB.Query(query)
	/*
		Checking Error
	*/
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Closing the statement without execution

	/*
		Looping through all the rows to step by step read them,
		and populate the "events" slice with that row data
	*/
	var events []Event
	for rows.Next() { // Next() method returns a boolean, which is true as long as there are rows left
		var event Event // Custom Struct
		/*
			Scan works like the FMT package
			Passing pointer, one for every columns that can be found in the row
		*/
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		/*
			Checking Error
		*/
		if err != nil {
			return nil, err
		}
		/*
			Populating the "events" slice using the "event" custom one
		*/
		events = append(events, event)
	}

	return events, nil
}
