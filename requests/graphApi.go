package requests

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var timeSlots = [...]string{"0800", "1200", "1600", "2000"}

//func GetGraphApiTimeSlotsFor(restaurantId int) []string {
//
//}

func getUrls(restaurantId int) []string {
	var urls []string
	for _, timeSlot := range timeSlots {
		url := getUrl(restaurantId, timeSlot)
		urls = append(urls, url)
	}
	return urls
}
func getUrl(restaurantId int, timeSlot string) string {
	currentDate := getCurrentDate()

	url := fmt.Sprintf(`https://s-varaukset.fi/api/recommendations/slot/%d/%s/%s/1`, restaurantId, currentDate, timeSlot)
	return url
}
func getCurrentDate() string {
	current := time.Now().String()
	indexOfSpace := strings.Index(current, " ")
	dateFromCurrent := current[:indexOfSpace]
	return dateFromCurrent
}
func GetGraphApiTimeSlotsFrom(requestUrl string) (*GraphApiResponse, error) {
	requestHandler := getGraphApiRequestHandler(requestUrl)
	response, err := sendRequestToGraphApi(requestHandler)
	if err != nil {
		return nil, fmt.Errorf("sendRequestToGraphApi - Error sending request to Raflaamo graph api. - %w", err)
	}
	responseBuffer, err := ReadResponseBuffer(response)
	if err != nil {
		return nil, fmt.Errorf("ReadResponseBuffer - Error reading response buffer. - %w", err)
	}
	deserializedResponse, err := deserializeGraphApiResponse(responseBuffer)
	if err != nil {
		return nil, fmt.Errorf("deserializeGraphApiResponse - Error deserializing response. - %w", err)
	}
	return deserializedResponse, nil
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
