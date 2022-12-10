package requests

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetValidRestaurants(t *testing.T) {
	GetRestaurantsFromApi = func(req *http.Request) (*http.Response, error) {
		resp := &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
		}
		return resp, nil
	}
	ReadResponseBuffer = func(res *http.Response) (*RestaurantApiResponse, error) {
		return &RestaurantApiResponse{}, nil
	}

	_, err := GetRestaurants()
	if err != nil {
		t.Errorf("GetRestaurants - Threw an unexpected error.")
	}
}
func TestGetInvalidRestaurants(t *testing.T) {
	GetRestaurantsFromApi = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("GetRestaurants - could not get response from api.raflaamo.fi/query")
	}

	_, err := GetRestaurants()
	if err == nil {
		t.Errorf("GetRestaurants - Did not throw an error when we expected it to.")
	}
}
