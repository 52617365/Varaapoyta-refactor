package requests

import (
	"errors"
	"net/http"
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
	deserializedResponse, err := deserializeResponse(response, &responseStructure)
	if err != nil {
		return nil, err
	}
	result := deserializedResponse.(*responseStructures.RestaurantApiResponse)
	return result, nil
}
