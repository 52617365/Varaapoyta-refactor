package requests

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"varaapoyta-backend-refactor/responseStructures"
)

func GetRestaurants() ([]responseStructures.Edges, error) {
	api := &Api{
		Name: "restaurant",
	}
	req := GetRequestHandlerFor(api)

	resp, err := getRestaurantsFromApi(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("GetRestaurants - could not get restaurants from api.raflaamo.fi/query")
	}
	restaurants, err := ReadRestaurantApiResponse(resp)

	if err != nil {
		return nil, err
	}
	validRestaurants := filterRestaurants(restaurants)
	return validRestaurants, nil
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

func filterRestaurants(restaurants *responseStructures.RestaurantApiResponse) []responseStructures.Edges {
	validRestaurants := getValidRestaurants(restaurants)
	setReservationIdsToRestaurants(validRestaurants)
	return validRestaurants
}

func getValidRestaurants(restaurants *responseStructures.RestaurantApiResponse) []responseStructures.Edges {
	validRestaurants := make([]responseStructures.Edges, 0, len(restaurants.Data.ListRestaurantsByLocation.Edges))
	for _, restaurant := range restaurants.Data.ListRestaurantsByLocation.Edges {
		if reservationPageExists(restaurant.Links.TableReservationLocalized.FiFI) { // TODO: add city into the filtering later.
			validRestaurants = append(validRestaurants, restaurant)
		}
	}
	return validRestaurants
}

// setReservationIdsToRestaurants the id returned from the endpoint is not the same as the one in the reservation page url.
// we need the latter to access the graph api.
func setReservationIdsToRestaurants(restaurants []responseStructures.Edges) {
	for index := range restaurants {
		reservationPageUrl := restaurants[index].Links.TableReservationLocalized.FiFI
		restaurantId, err := getReservationIdFrom(reservationPageUrl)
		if err != nil {
			// TODO: this should be logged because it should not have a problem finding the id after the check.
			continue
		}
		restaurants[index].ReservationPageID = restaurantId
	}
}

var RegexToMatchRestaurantId = regexp.MustCompile(`fi/(\d+)`)

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
