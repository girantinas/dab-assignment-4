package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
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

// Verifies the correctness of the getCookie function.
func TestGetCookie(t *testing.T) {
	// Each test will use a different set of cookies.
	tests := []struct {
		Name             string
		Cookies          []*http.Cookie
		ExpectedResponse string
	}{
		{"Basic Cookie", []*http.Cookie{{Name: "access_token", Value: "test_value"}}, "test_value"},
		{"Wrong Cookie", []*http.Cookie{{Name: "not_access_token", Value: "a_value"}}, ""},
		{"No Cookie", []*http.Cookie{}, ""},
		{"Multiple Cookies", []*http.Cookie{{Name: "access_token", Value: "cool_value"}, {Name: "not_access_token", Value: "a_true_value"}}, "cool_value"},
	}

	// Go through each one of our tests.
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Create a fake request and ResponseWriter.
			req := httptest.NewRequest(http.MethodGet, "/api/getCookie", nil)
			rec := httptest.NewRecorder()

			// Add the cookies for this test to the request.
			for _, cookie := range test.Cookies {
				req.AddCookie(cookie)
			}

			// Call the function using our fake request and recorder.
			getCookie(rec, req)

			// Check that everything matched what we expected.
			err := checkStatusCodeAndBody(http.StatusOK, rec.Result().StatusCode, test.ExpectedResponse, rec.Body.String())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

// Test the correctness of the getQuery function.
func TestGetQuery(t *testing.T) {
	// Each test will change the query parameter present in the request.
	tests := []struct {
		Name             string
		Params           string
		ExpectedResponse string
	}{
		{"Basic Query", "userID=40", "40"},
		{"No Params", "", ""},
		{"Multiple Query Params", "userID=the_id&secondParam=second", "the_id"},
		{"No userID", "aParam=Some_Parameter", ""},
	}

	// Run all the tests.
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Create a fake request with our query params and a ResponseWriter.
			req := httptest.NewRequest(http.MethodGet, "/api/getQuery?"+test.Params, nil)
			rec := httptest.NewRecorder()

			// Call the function.
			getQuery(rec, req)

			// Now test the Status Code and make sure the body is right.
			err := checkStatusCodeAndBody(http.StatusOK, rec.Result().StatusCode, test.ExpectedResponse, rec.Body.String())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

// Tests the correctness of the getJSON function.
func TestGetJSON(t *testing.T) {
	// Make our tests.
	tests := []struct {
		Name               string
		JSON               string
		ExpectedStatusCode int
		ExpectedResponse   string
	}{
		{"Basic JSON", normalJSON, http.StatusOK, "OskiBear\nHoshJug"},
		{"Extra Stuff In JSON", extraJSON, http.StatusOK, "The Finger\nThe Hill"},
		{"Missing Password", missingPasswordJSON, http.StatusBadRequest, ""},
		{"Missing Username", missingUsernameJSON, http.StatusBadRequest, ""},
		{"Empty JSON", emptyJSON, http.StatusBadRequest, ""},
		{"Empty Body", "", http.StatusBadRequest, ""},
	}

	// Run all of the tests.
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Create a fake request with our JSON and a ResponseWriter.
			req := httptest.NewRequest(http.MethodGet, "/api/getJSON", strings.NewReader(test.JSON))
			rec := httptest.NewRecorder()

			// Call the function with our JSON.
			getJSON(rec, req)

			// Now test that the correct code and body were returned.
			err := checkStatusCodeAndBody(test.ExpectedStatusCode, rec.Result().StatusCode, test.ExpectedResponse, rec.Body.String())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

// Tests the correctness of the Signup function.
func TestSignup(t *testing.T) {
	// Tests that the signup function can sign 50 users up.
	t.Run("Basic Signup", func(t *testing.T) {
		// Make sure there are no users already before starting the test.
		clearGlobalSlice()

		// Create a list of 50 test users. For each user call the function
		// so they are added to the global slice.
		users := make([]Credentials, 50)
		for i := 0; i < 50; i++ {
			users[i] = Credentials{strconv.Itoa(i), strconv.Itoa(i)}

			// Create a new request for the method with a JSON for this user.
			req, rec, err := createRequestAndResponseWithJSON(users[i], http.MethodPost, "/api/signup")
			if err != nil {
				t.Fatal(err)
			}

			// Call the function.
			signup(rec, req)

			// Now make sure the code is right.
			if rec.Result().StatusCode != http.StatusCreated {
				t.Fatalf("Failed to signup user %d. Got status code %d. Expected status code %d", i, rec.Result().StatusCode, http.StatusCreated)
			}
		}

		// Check that the slices have the same users in the same order.
		if !reflect.DeepEqual(UserSlice, users) {
			t.Error("Global slice has wrong contents.")
		}
	})

	// Check that the function errors on bad JSON.
	tests := []struct {
		Name string
		JSON string
	}{
		{"Missing Password", missingPasswordJSON},
		{"Missing Username", missingUsernameJSON},
		{"Empty JSON", emptyJSON},
		{"Empty Body", ""},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Setup a request for this bad JSON.
			req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(test.JSON))
			rec := httptest.NewRecorder()

			// Make sure the size of the slice before and after the call is the same.
			sizeBefore := len(UserSlice)
			signup(rec, req)
			if sizeBefore != len(UserSlice) {
				t.Fatal("Global slice got larger for a bad JSON!")
			}

			err := checkStatusCodeAndBody(http.StatusBadRequest, rec.Result().StatusCode, "", rec.Body.String())
			if err != nil {
				t.Fatal(err)
			}
		})
	}

	// Lastly, check that we get a conflict error if the same username is used twice.
	t.Run("Conflict", func(t *testing.T) {
		// Make sure there are no users already before starting the test.
		clearGlobalSlice()

		// The first user should be able to sign up with no problems.
		req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(normalJSON))
		rec := httptest.NewRecorder()
		signup(rec, req)
		if rec.Result().StatusCode != http.StatusCreated {
			t.Fatalf("Failed to signup first user. Got status code %d. Expected status code %d", rec.Result().StatusCode, http.StatusCreated)
		}

		// The second user should fail to signup.
		req = httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(normalJSON))
		rec = httptest.NewRecorder()
		sizeBefore := len(UserSlice)
		signup(rec, req)
		if sizeBefore != len(UserSlice) {
			t.Fatalf("User with conflicting name was able to sign up.")
		}
		if rec.Result().StatusCode != http.StatusConflict {
			t.Fatalf("Second user may have succeeded to sign up. Got status code %d. Expected status code %d", rec.Result().StatusCode, http.StatusConflict)
		}
	})
}

// Tests the correctness of the getIndex function.
func TestGetIndex(t *testing.T) {

	// This test makes sure the function returns an error when it tries to get the index of a user that doesn't exist.
	t.Run("No User", func(t *testing.T) {
		clearGlobalSlice()
		req := httptest.NewRequest(http.MethodGet, "/api/getIndex", strings.NewReader(normalJSON))
		rr := httptest.NewRecorder()

		getIndex(rr, req)

		err := checkStatusCodeAndBody(http.StatusBadRequest, rr.Result().StatusCode, "", rr.Body.String())
		if err != nil {
			t.Fatal(err)
		}
	})

	// Tests the basic functionality of the function.
	t.Run("Basic Index Retrieval", func(t *testing.T) {
		clearGlobalSlice()

		// Add a user to the global slice as if they had signed up.
		creds := Credentials{"student1", "dab"}
		UserSlice = append(UserSlice, creds)

		req, rr, err := createRequestAndResponseWithJSON(creds, http.MethodGet, "/api/getIndex")
		if err != nil {
			t.Fatal(err)
		}

		getIndex(rr, req)

		// We should get 0 back.
		err = checkStatusCodeAndBody(http.StatusOK, rr.Result().StatusCode, "0", rr.Body.String())
		if err != nil {
			t.Fatal(err)
		}
	})
}

// Tests the correctness of the getPassword function.
func TestGetPW(t *testing.T) {
	// Basic functionality test.
	t.Run("Basic Get Password", func(t *testing.T) {
		// Get 1 user into the global slice.
		clearGlobalSlice()
		creds := Credentials{"student1", "dab"}
		UserSlice = append(UserSlice, creds)

		req := httptest.NewRequest(http.MethodGet, "/api/getPassword", strings.NewReader(`{"username":"student1"}`))
		rr := httptest.NewRecorder()

		getPassword(rr, req)

		err := checkStatusCodeAndBody(http.StatusOK, rr.Result().StatusCode, "dab", rr.Body.String())
		if err != nil {
			t.Fatal(err)
		}
	})

	// Get password of unregistered user
	t.Run("Nonexistent User", func(t *testing.T) {
		clearGlobalSlice()
		creds := Credentials{"student1", "dab"}
		UserSlice = append(UserSlice, creds)

		req := httptest.NewRequest(http.MethodGet, "/api/getPassword", strings.NewReader("student001"))
		rr := httptest.NewRecorder()

		getPassword(rr, req)

		err := checkStatusCodeAndBody(http.StatusBadRequest, rr.Result().StatusCode, "", rr.Body.String())
		if err != nil {
			t.Fatal(err)
		}
	})
}

// Tests the correctness of the updatePassword function.
func TestUpdatePW(t *testing.T) {
	t.Run("Basic Update", func(t *testing.T) {
		clearGlobalSlice()
		creds := Credentials{"student1", "dab"}
		UserSlice = append(UserSlice, creds)

		// Change their password to something else.
		creds = Credentials{"student1", "dabdab"}

		req, rr, err := createRequestAndResponseWithJSON(creds, http.MethodPut, "/api/updatePW")
		if err != nil {
			t.Fatal(err)
		}

		updatePassword(rr, req)

		if rr.Result().StatusCode != http.StatusOK {
			t.Fatalf("Incorrect status code returned! Expected: %d Actual: %d", http.StatusOK, rr.Result().StatusCode)
		}

		// Check that the UserSlice is still the same length and the password has been updated.
		if len(UserSlice) != 1 || UserSlice[0].Password != "dabdab" {
			t.Fatal("Password not updated!")
		}
	})

	// Try to update password of unregistered user
	t.Run("Update Non Existent", func(t *testing.T) {
		clearGlobalSlice()
		creds := Credentials{"student1", "dab"}
		UserSlice = append(UserSlice, creds)

		creds = Credentials{"student001", "dabdab"}

		req, rr, err := createRequestAndResponseWithJSON(creds, http.MethodPut, "/api/updatePW")
		if err != nil {
			t.Fatal(err)
		}

		updatePassword(rr, req)

		if rr.Result().StatusCode != http.StatusBadRequest {
			t.Fatalf("Incorrect status code returned! Expected: %d Actual: %d", http.StatusBadRequest, rr.Result().StatusCode)
		}

		if len(UserSlice) != 1 || UserSlice[0].Password != "dab" {
			t.Fatal("Password updated when an error occurred!")
		}
	})
}

// Tests the correctness of the deleteUser function.
func TestDeleteUser(t *testing.T) {
	// Simply deletes a user that's in the slice.
	t.Run("Delete Basic", func(t *testing.T) {
		clearGlobalSlice()
		creds := Credentials{"student1", "dab"}
		UserSlice = append(UserSlice, creds)

		req, rr, err := createRequestAndResponseWithJSON(creds, http.MethodDelete, "/api/deleteUser")
		if err != nil {
			t.Fatal(err)
		}

		deleteUser(rr, req)

		if rr.Result().StatusCode != http.StatusOK {
			t.Fatalf("Incorrect status code returned! Expected: %d Actual: %d", http.StatusOK, rr.Result().StatusCode)
		}
		if len(UserSlice) != 0 {
			t.Fatalf("User was not deleted from slice!")
		}
	})

	// Tries to delete a user that doesn't exist.
	t.Run("Delete Non-Existent", func(t *testing.T) {
		clearGlobalSlice()
		creds := Credentials{"student1", "dab"}
		UserSlice = append(UserSlice, creds)

		creds = Credentials{"student0001", ""}
		req, rr, err := createRequestAndResponseWithJSON(creds, http.MethodDelete, "/api/deleteUser")
		if err != nil {
			t.Fatal(err)
		}

		deleteUser(rr, req)

		if rr.Result().StatusCode != http.StatusBadRequest {
			t.Fatalf("Incorrect status code returned! Expected: %d Actual: %d", http.StatusBadRequest, rr.Result().StatusCode)
		}
		if len(UserSlice) == 0 {
			t.Fatalf("User was deleted from slice!")
		}
	})
}

// Helper Methods and JSON

// Make some test JSON objects
const (
	normalJSON          = `{"username":"OskiBear", "password":"HoshJug"}`
	extraJSON           = `{"username":"The Finger", "password":"The Hill", "the sum of all":"human knowledge"}`
	missingPasswordJSON = `{"username":"No Password"}`
	missingUsernameJSON = `{"password":"oops forgot my username"}`
	emptyJSON           = `{}`
)

// Checks to make sure the body and status code are what we expect.
// Gives an error with a descriptive message if they don't match.
func checkStatusCodeAndBody(expectedCode, actualCode int, expectedBody, actualBody string) error {
	if actualCode != expectedCode {
		return fmt.Errorf("Incorrect status code returned! Expected: %d Actual: %d", expectedCode, actualCode)
	}
	if actualBody != expectedBody {
		return fmt.Errorf("Incorrect body returned! Expected: \"%s\" Actual: \"%s\"", expectedBody, actualBody)
	}
	return nil
}

// Sets the global slice of Credentials to an empty slice. Useful for
// ensuring the tests stay independent.
func clearGlobalSlice() {
	UserSlice = make([]Credentials, 0)
}

// Given an object that can be marshalled by a JSON, creates a request and response
// pair to be sent to a function with that object in a JSON.
func createRequestAndResponseWithJSON(jsonObj interface{}, method, endpoint string) (*http.Request, *httptest.ResponseRecorder, error) {
	jsonBytes, err := json.Marshal(jsonObj)
	if err != nil {
		return nil, nil, fmt.Errorf("An error occurred while marshalling the JSON: %s", err)
	}
	b := bytes.NewBuffer(jsonBytes)
	req := httptest.NewRequest(method, endpoint, b)
	rec := httptest.NewRecorder()
	return req, rec, nil
}
