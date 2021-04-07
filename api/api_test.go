package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Verifies that all of the proper routes have been registered
// by the server with the correct HTTP method.
func TestRegisterRoutes(t *testing.T) {
	// Create a list that contains all of the routes used by our server.
	// Recall a route is a combination of an API endpoint and an HTTP method.
	routes := []struct {
		Endpoint string
		Method   string
	}{
		{"/api/getCookie", http.MethodGet},
		{"/api/getQuery", http.MethodGet},
		{"/api/getJSON", http.MethodGet},
		{"/api/signup", http.MethodPost},
		{"/api/getIndex", http.MethodGet},
		{"/api/getPW", http.MethodGet},
		{"/api/updatePW", http.MethodPut},
		{"/api/deleteUser", http.MethodDelete},
	}

	// Create a new mux router and register all the routes on it.
	router := mux.NewRouter()
	RegisterRoutes(router)

	// Now check that the router responds to all the routes.
	for _, route := range routes {

		// Make a fake request we can use to probe the router.
		req := httptest.NewRequest(route.Method, route.Endpoint, nil)

		// Make a RouteMatch struct that will hold information about the
		// route that matched our request.
		match := &mux.RouteMatch{}

		// Check that some route matched our request in the router.
		if matched := router.Match(req, match); matched == false {
			if match.MatchErr == mux.ErrMethodMismatch {
				t.Errorf("Endpoint %s does not respond to the correct method. It should respond to method %s.", route.Endpoint, route.Method)
			} else {
				t.Errorf("Could not find a route associated with endpoint %s.", route.Endpoint)
			}
			continue
		}

		// Get all the methods on the route so we can verify them.
		methods, err := match.Route.GetMethods()

		// Check that the endpoint has at least one method registered to it.
		if err != nil {
			t.Errorf("Endpoint %s has no methods registered.", route.Endpoint)
			continue
		}

		// Each endpoint in our API has exactly one HTTP method associated with it.
		if len(methods) > 1 {
			t.Errorf("Endpoint %s has %d methods registered to it. It should only have 1.", route.Endpoint, len(methods))
		}
	}
}

// TODO: Write more tests!

func TestGetIndex(t *testing.T) {
	clearGlobalSlice()
	handlerGetIndex := http.HandlerFunc(getIndex)

	// First try to get index of nonexistent user
	var jsonGetIndex = []byte(`{"username":"student1"}`)

	req, err := http.NewRequest("GET", "/api/getIndex", bytes.NewBuffer(jsonGetIndex))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlerGetIndex.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	signUserUp("student1", "dab")

	// Now get index of registered user
	req, err = http.NewRequest("GET", "/api/getIndex", bytes.NewBuffer(jsonGetIndex))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handlerGetIndex.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `0`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	clearGlobalSlice()
}

func TestGetPW(t *testing.T) {
	clearGlobalSlice()
	handlerGetPW := http.HandlerFunc(getPassword)

	// Get password of registered user
	signUserUp("student1", "dab")

	var jsonGetPW = []byte(`{"username":"student1"}`)

	req, err := http.NewRequest("GET", "/api/getPW", bytes.NewBuffer(jsonGetPW))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlerGetPW.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `dab`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Get password of unregistered user
	var jsonGetPWBad = []byte(`{"username":"student001"}`)

	req, err = http.NewRequest("GET", "/api/getPW", bytes.NewBuffer(jsonGetPWBad))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handlerGetPW.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	clearGlobalSlice()
}

func TestUpdatePW(t *testing.T) {
	clearGlobalSlice()
	handlerGetPW := http.HandlerFunc(getPassword)
	handlerUpdatePW := http.HandlerFunc(updatePassword)

	// Sign user up
	signUserUp("student1", "dab")

	var jsonGetPW = []byte(`{"username":"student1"}`)

	req, err := http.NewRequest("GET", "/api/getPW", bytes.NewBuffer(jsonGetPW))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlerGetPW.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `dab`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Update password
	var jsonUpdatePW = []byte(`{"username":"student1", "password":"dabdab"}`)

	req, err = http.NewRequest("GET", "/api/updatePW", bytes.NewBuffer(jsonUpdatePW))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handlerUpdatePW.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verify password was changed
	var jsonGetNewPW = []byte(`{"username":"student1"}`)

	req, err = http.NewRequest("GET", "/api/getPW", bytes.NewBuffer(jsonGetNewPW))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handlerGetPW.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected = `dabdab`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Try to update password of unregistered user
	var jsonUpdatePWBad = []byte(`{"username":"student001"}`)

	req, err = http.NewRequest("GET", "/api/updatePW", bytes.NewBuffer(jsonUpdatePWBad))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handlerUpdatePW.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	clearGlobalSlice()
}

func TestDeleteUser(t *testing.T) {
	clearGlobalSlice()
	handlerDelete := http.HandlerFunc(deleteUser)

	signUserUp("student1", "dab")

	var jsonDelete = []byte(`{"username":"student1"}`)
	req, _ := http.NewRequest("DELETE", "/api/deleteUser", bytes.NewBuffer(jsonDelete))

	rr := httptest.NewRecorder()
	handlerDelete.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Try to delete user again
	req, _ = http.NewRequest("DELETE", "/api/deleteUser", bytes.NewBuffer(jsonDelete))

	rr = httptest.NewRecorder()
	handlerDelete.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func signUserUp(username string, password string) {
	var jsonStr = []byte(fmt.Sprintf(`{"username":"%v", "password":"%v"}`, username, password))

	rr := httptest.NewRecorder()
	handlerSignup := http.HandlerFunc(signup)

	req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(jsonStr))

	handlerSignup.ServeHTTP(rr, req)
}

func clearUser(username string) {
	var jsonStr = []byte(fmt.Sprintf(`{"username":"%v"}`, username))

	rr := httptest.NewRecorder()
	handlerDelete := http.HandlerFunc(deleteUser)

	req, _ := http.NewRequest("DELETE", "/api/deleteUser", bytes.NewBuffer(jsonStr))

	handlerDelete.ServeHTTP(rr, req)
}
