package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
		Sending a response in JSON format
		Status code 200 = StatusOK
		H map with message
	*/
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
