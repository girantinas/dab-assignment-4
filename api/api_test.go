package api

import (
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
