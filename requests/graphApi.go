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
		return nil, fmt.Errorf("sendRequestToGraphApi - Error sending request to Raflaamo graph api. - %w", err)
	}
	return response, nil
}

func sendRequestToGraphApi(requestHandler *http.Request) (*http.Response, error) {
	response, err := sendRequest(requestHandler)
	return response, err
}
