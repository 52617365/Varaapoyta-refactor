package requests

import (
	"fmt"
	"net/http"
)

func GetGraphApiTimeSlotsFrom(requestUrl string) (*GraphApiResponse, error) {
	response, err := getResponseFromGraphApi(requestUrl)
	if err != nil {
		return nil, err
	}
	deserializedResponse, err := deserializeGraphApiResponse(response)
	if err != nil {
		return nil, fmt.Errorf("deserializeGraphApiResponse - Error deserializing response. - %w", err)
	}
	return deserializedResponse, nil
}

func getResponseFromGraphApi(requestUrl string) ([]byte, error) {
	requestHandler := getGraphApiRequestHandler(requestUrl)
	response, err := sendRequestToGraphApi(requestHandler)
	if err != nil {
		return nil, fmt.Errorf("sendRequestToGraphApi - Error sending request to Raflaamo graph api. - %w", err)
	}
	responseBuffer, err := ReadResponseBuffer(response)
	if err != nil {
		return nil, fmt.Errorf("ReadResponseBuffer - Error reading response buffer. - %w", err)
	}
	return responseBuffer, nil
}

func getGraphApiRequestHandler(requestUrl string) *http.Request {
	graphApi := Api{
		Name: "graph",
		Url:  requestUrl,
	}
	requestHandler := GetRequestHandlerFor(&graphApi)
	return requestHandler
}

func deserializeGraphApiResponse(responseBuffer []byte) (*GraphApiResponse, error) {
	deserializedType := GraphApiResponse{}
	deserializedResponse, err := deserializeResponse(responseBuffer, &deserializedType)
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
