package models //All the logic that deals with storing event data in a database
import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

func Save() {

}
