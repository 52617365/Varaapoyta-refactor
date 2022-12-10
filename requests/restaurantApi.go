package requests

import (
	"encoding/json"
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
	deserializedResponse, err := deserializeResponse(readBuffer)
	if err != nil {
		return &RestaurantApiResponse{}, err
	}
	return deserializedResponse, nil
}

func deserializeResponse(response []byte) (*RestaurantApiResponse, error) {
	responseStructure := RestaurantApiResponse{}
	err := json.Unmarshal(response, &responseStructure)
	if err != nil {
		return nil, err
	}
	return &responseStructure, nil
}
