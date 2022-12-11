package requests

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"varaapoyta-backend-refactor/responseStructures"
)

func GetRestaurants() (*responseStructures.RestaurantApiResponse, error) {
	api := &Api{
		Name: "restaurant",
	}
	req := GetRequestHandlerFor(api)

	resp, err := getRestaurantsFromApi(req)
	if err != nil {
		return &responseStructures.RestaurantApiResponse{}, err
	}
	if resp.StatusCode != 200 {
		return &responseStructures.RestaurantApiResponse{}, errors.New("GetRestaurants - could not get response from api.raflaamo.fi/query")
	}
	response, err := ReadRestaurantApiResponse(resp)

	if err != nil {
		return &responseStructures.RestaurantApiResponse{}, err
	}
	setReservationIdsToRestaurants(response)
	return response, nil
}

func getRestaurantsFromApi(req *http.Request) (*http.Response, error) {
	resp, err := sendRequest(req)
	return resp, err
}

var ReadRestaurantApiResponse = func(res *http.Response) (*responseStructures.RestaurantApiResponse, error) {
	readBuffer, err := ReadResponseBuffer(res)
	if err != nil {
		return &responseStructures.RestaurantApiResponse{}, err
	}
	deserializedResponse, err := deserializeRestaurantApiResponse(readBuffer)
	if err != nil {
		return &responseStructures.RestaurantApiResponse{}, err
	}
	return deserializedResponse, nil
}

func deserializeRestaurantApiResponse(response []byte) (*responseStructures.RestaurantApiResponse, error) {
	responseStructure := responseStructures.RestaurantApiResponse{}
	deserializedResponse, err := deserializeResponse(response, responseStructure)
	if err != nil {
		return nil, err
	}
	result := deserializedResponse.(*responseStructures.RestaurantApiResponse)
	return result, nil
}

var RegexToMatchRestaurantId = regexp.MustCompile(`fi/(\d+)`)

// setReservationIdsToRestaurants the id returned from the endpoint is not the same as the one in the reservation page url.
// we need the latter to access the graph api.
func setReservationIdsToRestaurants(restaurants *responseStructures.RestaurantApiResponse) {
	for index, restaurant := range restaurants.Data.ListRestaurantsByLocation.Edges {
		reservationPageUrl := restaurant.Links.TableReservationLocalized.FiFI
		if reservationPageExists(reservationPageUrl) {
			restaurantId, err := getReservationIdFrom(reservationPageUrl)
			if err != nil {
				// TODO: this should be logged because it should not have a problem finding the id after the check.
				continue
			}
			restaurants.Data.ListRestaurantsByLocation.Edges[index].ReservationPageID = restaurantId
		}
	}
}

func getReservationIdFrom(reservationPageUrl string) (string, error) {
	restaurantIdMatch := RegexToMatchRestaurantId.FindAllStringSubmatch(reservationPageUrl, -1)
	if restaurantIdMatch == nil {
		return "", errors.New("getReservationIdFrom - could not get restaurant id from reservation page url")
	}
	restaurantId := restaurantIdMatch[0][1]
	return restaurantId, nil
}
func reservationPageExists(reservationPageUrl string) bool {
	if reservationPageUrl == "" {
		return false
	}
	if !strings.Contains(reservationPageUrl, "https://s-varaukset.fi/online/reservation/fi/") {
		return false
	}
	return true
}
