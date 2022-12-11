package requests

import (
	"errors"
	"net/http"
	"testing"
	"varaapoyta-backend-refactor/responseStructures"
)

func TestGetValidRestaurants(t *testing.T) {
	sendRequest = func(req *http.Request) (*http.Response, error) {
		resp := &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
		}
		return resp, nil
	}
	ReadRestaurantApiResponse = func(res *http.Response) (*responseStructures.RestaurantApiResponse, error) {
		return &responseStructures.RestaurantApiResponse{}, nil
	}

	_, err := GetRestaurants()
	if err != nil {
		t.Errorf("GetRestaurants - Threw an unexpected error.")
	}
}
func TestGetInvalidRestaurants(t *testing.T) {
	sendRequest = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("GetRestaurants - could not get response from api.raflaamo.fi/query")
	}

	_, err := GetRestaurants()
	if err == nil {
		t.Errorf("GetRestaurants - Did not throw an error when we expected it to.")
	}
}

func TestGetReservationPageIdFrom(t *testing.T) {
	urlToMatch := "https://s-varaukset.fi/online/reservation/fi/72?_ga=2.161416895.382807502.1612853101-189045693.1611044564"
	expected := "72"
	actual, err := getReservationIdFrom(urlToMatch)
	if err != nil {
		t.Errorf("getReservationIdFrom - Threw an unexpected error.")
	}
	if actual != expected {
		t.Errorf("getReservationPageIdFrom - expected %s, got %s", expected, actual)
	}
}
func TestErrorGetReservationPageIdFrom(t *testing.T) {
	urlToMatch := "https://s-varaukset.fi/"
	actual, err := getReservationIdFrom(urlToMatch)
	if err == nil {
		t.Errorf("getReservationIdFrom - expected error, got %s", actual)
	}
}
func TestReservationPageExistsTrue(t *testing.T) {
	urlToMatch := "https://s-varaukset.fi/online/reservation/fi/72?_ga=2.161416895.382807502.1612853101-189045693.1611044564"
	actual := reservationPageExists(urlToMatch)
	if !actual {
		t.Errorf("reservationPageExists - expected true, got %t", actual)
	}
}
func TestReservationPageExistsFalse(t *testing.T) {
	urlToMatch := ""
	actual := reservationPageExists(urlToMatch)
	if actual {
		t.Errorf("reservationPageExists - expected false, got %t", actual)
	}
}
func TestReservationPageExistsFalse2(t *testing.T) {
	urlToMatch := "randomlink.com"
	actual := reservationPageExists(urlToMatch)
	if actual {
		t.Errorf("reservationPageExists - expected false, got %t", actual)
	}
}
