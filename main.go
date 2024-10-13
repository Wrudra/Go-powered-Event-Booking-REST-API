package main

import (
	"net/http"
	"root/db"     //Database package
	"root/models" //Event model package

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		Initializing database after calling the database package
	*/
	db.InitDB()

	/*
		Default() essentially configues an HTTP server for us with basic functionalities,
		to log incoming requests and to automatically recover if some part of our program should crash.
		Stored in a variable named 'server'
	*/
	server := gin.Default()

	/*
		Handler for an incoming HTTP GET request.
		Incoming GET requests to "/events"
		Set up a function as second argument that will be executed, if such GET request is sent to "/events"
	*/
	server.GET("/events", getEvents)

	/*
		Registering a second endpoint which handles incoming POST requests
		Self explanatory
	*/
	server.POST("/events", createEvent)

	/*
		When the server is started, and starts listening for incoming requests.
		Passed a string 8080 to make sure that we listen to incoming requests on some domain.
		Requests to come to port 8080
	*/
	server.Run(":8080") // localhost:8080
}

/*
The response func if we get a request in our HTTP server
*/
func getEvents(context *gin.Context) { //pointer
	/*
		After importing Models package,
		We get back the events that are managed in that separate file
	*/
	events := models.GetAllEvents()

	/*
		Sending a response in JSON format
		Status code 200 = StatusOK
		We send back the events by simply passing events as a value to context.Json
	*/
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	/*
		ShouldBindJSON() works like scan func from the FMT package
		Passing the pointer of the variable, for storung data in that variable
	*/
	err := context.ShouldBindJSON(&event)

	/*
		Checking error if any data is missing for creating an event
	*/
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	/*
		After passing error test,
		server-generated ID as it is not required by the user for creating an event.
		For now used dummy value 1, that will change later once we use a database.
	*/
	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event}) //"event":event sending back that event that was created
}
