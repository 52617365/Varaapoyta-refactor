package requests

import (
	"errors"
	"golang.org/x/exp/slices"
	"net/http"
	"testing"
)

func TestGetRequestHandler(t *testing.T) {
	expectedHeaderKeys := []string{"Content-Type", "Client_id", "User-Agent"}
	expectedHeaderValues := []string{"application/json", "jNAWMvWD9rp637RaR", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"}

	req := GetRequestHandler()

	for key, value := range req.Header {
		if !slices.Contains(expectedHeaderKeys, key) {
			t.Errorf("GetRequestHandler - header key %v not found", key)
		}
		if !slices.Contains(expectedHeaderValues, value[0]) {
			t.Errorf("GetRequestHandler - header value %v not found", value)
		}
	}
	if req.Method != "POST" {
		t.Errorf("GetRequestHandler - expected method to be POST, got %v", req.Method)
	}
	if req.URL.String() != "https://api.raflaamo.fi/query" {
		t.Errorf("GetRequestHandler - expected url to be https://api.raflaamo.fi/query, got %v", req.URL.String())
	}

}

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
