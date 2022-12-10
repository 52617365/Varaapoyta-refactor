package requests

import (
	"errors"
	"net/http"
)

func GetRestaurants() (*RestaurantApiResponse, error) {
	api := &Api{
		Name: "restaurant",
	}
	req := GetRequestHandlerFor(api)

	resp, err := getRestaurantsFromApi(req)
	if err != nil {
		return &RestaurantApiResponse{}, err
	}
	if resp.StatusCode != 200 {
		return &RestaurantApiResponse{}, errors.New("GetRestaurants - could not get response from api.raflaamo.fi/query")
	}
	response, err := ReadRestaurantApiResponse(resp)

	if err != nil {
		return &RestaurantApiResponse{}, err
	}
	return response, nil
}

func getRestaurantsFromApi(req *http.Request) (*http.Response, error) {
	resp, err := sendRequest(req)
	return resp, err
}

var ReadRestaurantApiResponse = func(res *http.Response) (*RestaurantApiResponse, error) {
	readBuffer, err := ReadResponseBuffer(res)
	if err != nil {
		return &RestaurantApiResponse{}, err
	}
	deserializedResponse, err := deserializeRestaurantApiResponse(readBuffer)
	if err != nil {
		return &RestaurantApiResponse{}, err
	}
	return deserializedResponse, nil
}

func deserializeRestaurantApiResponse(response []byte) (*RestaurantApiResponse, error) {
	responseStructure := RestaurantApiResponse{}
	deserializedResponse, err := deserializeResponse(response, &responseStructure)
	if err != nil {
		return nil, err
	}
	result := deserializedResponse.(*RestaurantApiResponse)
	return result, nil
}
