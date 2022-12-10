package requests

import (
	"fmt"
	"net/http"
)

func GetGraphApiTimeSlotsFrom(requestUrl string) (*http.Response, error) {
	graphApi := Api{
		Name: "graph",
		Url:  requestUrl,
	}
	requestHandler := GetRequestHandlerFor(&graphApi)
	response, err := sendRequestToGraphApi(requestHandler)
	if err != nil {
		return &GraphApiResponse{}, nil
	}
	result := deserializedResponse.(*GraphApiResponse)
	return result, nil
}

func sendRequestToGraphApi(requestHandler *http.Request) (*http.Response, error) {
	response, err := sendRequest(requestHandler)
	return response, err
}
