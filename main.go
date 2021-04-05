package main

import (
	"log"
	"net/http"

	"github.com/BearCloud/sp21-assignment-4/api"
	"github.com/gorilla/mux"
)

// Starts the server and has it listen for requests.
func main() {
	// Create a new mux for routing api calls
	router := mux.NewRouter()

	//Register our endpoints
	//See api/api.go
	api.RegisterRoutes(router)

	//Print log to output, very similar to fmt.Println
	//What are the differences?
	log.Println("starting go server")

	// Has the server listen on port 80 using the routes
	// registered earlier.
	http.ListenAndServe(":80", router)
}
