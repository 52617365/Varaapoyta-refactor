package requests

import (
	"encoding/json"
	"errors"
	"io"
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
	response, err := ReadResponseBuffer(resp)
	if err != nil {
		return &RestaurantApiResponse{}, err
	}
	return response, nil
}

func getRestaurantsFromApi(req *http.Request) (*http.Response, error) {
	resp, err := sendRequest(req)
	return resp, err
}

var ReadResponseBuffer = func(res *http.Response) (*RestaurantApiResponse, error) {
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return &RestaurantApiResponse{}, err
	}

	deserializedResponse, err := deserializeResponse(b)
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
