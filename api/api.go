package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Declares a global slice of Credentials you will modify in the functions below.
// See credentials.go
var UserSlice []Credentials

// Given a gorilla/mux Router, registers the required HTTP endpoints
// for each of the routes in our server.
func RegisterRoutes(router *mux.Router) {
	// We have done the first 3 routes for you. Register the remaining ones
	// based on the API given in API.md after reading over all the functions below.
	router.HandleFunc("/api/getCookie", getCookie).Methods(http.MethodGet)
	router.HandleFunc("/api/getQuery", getQuery).Methods(http.MethodGet)
	router.HandleFunc("/api/getJSON", getJSON).Methods(http.MethodGet)

	/* YOUR CODE HERE */
}

// Obtain the "access_token" cookie's value and write it to the response.
// If there is no such cookie, write an empty string to the response.
func getCookie(response http.ResponseWriter, request *http.Request) {
	/*YOUR CODE HERE*/
}

// Obtain the "userID" query paramter and write it to the response.
// If there is no such query parameter, write an empty string to the response.
func getQuery(response http.ResponseWriter, request *http.Request) {
	/*YOUR CODE HERE*/
}

// Our JSON file will look like this:
//
// {
// 	 "username" : <username>,
// 	 "password" : <password>
// }
//
// Decode this JSON file into an instance of Credentials.
// Then, write the username and password to the response, separated by a newline.
//
// Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest".
// What kind of errors can we expect here?
func getJSON(response http.ResponseWriter, request *http.Request) {
	/*YOUR CODE HERE*/
}

// Our JSON file will look like this:
//
// {
//	 "username" : <username>,
//	 "password" : <password>
// }
//
// Decode this JSON file into an instance of Credentials.
// Then store it ("append" it) to the global slice of Credentials.
//
// Make sure to error check! If there are any errors, call http.Error(), and pass in the appropriate error (See API.md).
// What kind of errors can we expect here?
//
// If you aren't sure how to append to a slice, check this out: https://tour.golang.org/moretypes/15.
// On success, make sure the status code is 201 Status Created!
func signup(response http.ResponseWriter, request *http.Request) {
	/*YOUR CODE HERE*/
}

// Our JSON file will look like this:
//
// {
//	 "username" : <username>
// }
//
// Decode this JSON file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)
// Return the array index of the Credentials object in the global Credentials slice.
//
// The index will be of type integer, but we can only write strings to the response. What library and function was used to get around this?
//
// Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest".
// What kind of errors can we expect here?
func getIndex(response http.ResponseWriter, request *http.Request) {
	/*YOUR CODE HERE*/
}

// Our JSON file will look like this:
//
// {
//	 "username" : <username>
// }
//
// Decode this JSON file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)
// Write the password of the specific user to the response.
//
// Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest".
// What kind of errors can we expect here?
func getPassword(response http.ResponseWriter, request *http.Request) {
	/*YOUR CODE HERE*/
}

// Our JSON file will look like this:
//
// {
// 	"username" : <username>,
// 	"password" : <password,
// }
//
// Decode this JSON file into an instance of Credentials.
// The password in the JSON file is the new password they want to replace the old password with.
// You don't need to return anything in this.
//
// Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest".
// What kind of errors can we expect here?
func updatePassword(response http.ResponseWriter, request *http.Request) {
	/*YOUR CODE HERE*/
}

// Our JSON file will look like this:
//
// {
// 	"username" : <username>,
// 	"password" : <password,
// }
//
// Decode this JSON file into an instance of Credentials.
// Remove this user from the array. Preserve the original order. You may want to create a helper function.
//
// This wasn't covered in lecture, so you may want to read the following:
// 	- https://gobyexample.com/slices
// 	- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/
//
// Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest".
// What kind of errors can we expect here?
func deleteUser(response http.ResponseWriter, request *http.Request) {
	/*YOUR CODE HERE*/
}
