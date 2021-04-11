package api

// We've left in some packages that you may find helpful while
// implementing this assignment. You're free to use whatever packages
// you'd like, but these are the ones we used to do this. To use the package,
// just remove the underscore in front of it.
import (
	"encoding/json"
	"errors"
	"fmt"
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
	router.HandleFunc("/api/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/api/getIndex", getIndex).Methods(http.MethodGet)
	router.HandleFunc("/api/getPW", getPassword).Methods(http.MethodGet)
	router.HandleFunc("/api/updatePW", updatePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/deleteUser", deleteUser).Methods(http.MethodDelete)
}

// Obtain the "access_token" cookie's value and write it to the response.
// If there is no such cookie, write an empty string to the response.
func getCookie(response http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("access_token")
	if err != nil {
		fmt.Fprint(response, "")
	} else {
		fmt.Fprint(response, cookie.Value)
	}
}

// Obtain the "userID" query parameter and write it to the response.
// If there is no such query parameter, write an empty string to the response.
func getQuery(response http.ResponseWriter, request *http.Request) {
	userIDQuery := request.URL.Query().Get("userID")
	fmt.Fprint(response, userIDQuery)
}

//Reads an HTTP Request as a credentials pointer,
// passing back an error in the case of problems.
func readJSON(request *http.Request) (*Credentials, error) {
	var creds Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		return &creds, errors.New("Bad JSON")
	} else if creds.Password == "" {
		return &creds, errors.New("No Password")
	} else if creds.Username == "" {
		return &creds, errors.New("No Username")
	} else {
		return &creds, nil
	}
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
// Make sure to error check! What kind of errors can we expect here?
func getJSON(response http.ResponseWriter, request *http.Request) {
	creds, err := readJSON(request)
	if err != nil {
		http.Error(response, "", http.StatusBadRequest)
	} else {
		fmt.Fprint(response, creds.Username+"\n"+creds.Password)
	}
}

//Returns the index of a user with a given username
func findUser(username string) (int, error) {
	for i, creds := range UserSlice {
		if creds.Username == username {
			return i, nil
		}
	}
	return -1, errors.New("User Not Found")
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
// Make sure to error check! What kind of errors can we expect here?
//
// If you aren't sure how to append to a slice, check this out: https://tour.golang.org/moretypes/15.
// On success, make sure the status code is 201 Status Created!
func signup(response http.ResponseWriter, request *http.Request) {
	creds, err := readJSON(request)
	if err != nil {
		http.Error(response, "", http.StatusBadRequest)
	} else {
		_, userErr := findUser(creds.Username)
		if userErr != nil {
			UserSlice = append(UserSlice, *creds)
			response.WriteHeader(201)
		} else {
			http.Error(response, "", http.StatusConflict)
		}
	}
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
// Make sure to error check! What kind of errors can we expect here?
func getIndex(response http.ResponseWriter, request *http.Request) {
	creds, err := readJSON(request)
	if err != nil && err.Error() != "No Password" {
		http.Error(response, "", http.StatusBadRequest)
	} else {
		index, userErr := findUser(creds.Username)
		if userErr != nil {
			http.Error(response, "", http.StatusBadRequest)
			fmt.Printf("Got here")
		} else {
			fmt.Fprintf(response, "%d", index)
		}
	}
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
// Make sure to error check! What kind of errors can we expect here?
func getPassword(response http.ResponseWriter, request *http.Request) {
	creds, err := readJSON(request)
	if err != nil && err.Error() != "No Password" {
		http.Error(response, "", http.StatusBadRequest)
	} else {
		index, userErr := findUser(creds.Username)
		if userErr != nil {
			http.Error(response, "", http.StatusBadRequest)
		} else {
			fmt.Fprint(response, UserSlice[index].Password)
		}
	}
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
// Make sure to error check! What kind of errors can we expect here?
func updatePassword(response http.ResponseWriter, request *http.Request) {
	creds, err := readJSON(request)
	if err != nil {
		http.Error(response, "", http.StatusBadRequest)
	} else {
		index, userErr := findUser(creds.Username)
		if userErr != nil {
			http.Error(response, "", http.StatusBadRequest)
		} else {
			UserSlice[index].Password = creds.Password
		}
	}
}

func remove(slice []Credentials, index int) []Credentials {
	end := len(slice) - 1
	slice[index] = slice[end]
	return slice[:end]
}

// Our JSON file will look like this:
//
// {
// 	"username" : <username>
// }
//
// Decode this JSON file into an instance of Credentials.
// Remove this user from the array. Preserve the original order. You may want to create a helper function.
//
// This wasn't covered in lecture, so you may want to read the following:
// 	- https://gobyexample.com/slices
// 	- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/
//
// Make sure to error check! What kind of errors can we expect here?
func deleteUser(response http.ResponseWriter, request *http.Request) {
	creds, err := readJSON(request)
	if err != nil {
		http.Error(response, "", http.StatusBadRequest)
	} else {
		index, userErr := findUser(creds.Username)
		if userErr != nil {
			http.Error(response, "", http.StatusBadRequest)
		} else {
			UserSlice = remove(UserSlice, index)
		}
	}
}
