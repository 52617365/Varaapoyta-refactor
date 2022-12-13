package requests

import (
	"errors"
	"fmt"
	"net/http"
	"varaapoyta-backend-refactor/responseStructures"
	"varaapoyta-backend-refactor/time"
)

func GetGraphApiTimeSlotsFrom(restaurantId string) ([]string, error) {
	urls := GetGraphApiUrls(restaurantId)
	var deDupTimeSlots []string
	for _, url := range urls {
		timeSlots, err := GetTimeSlotFrom(url)
		if err != nil {
			if urlShouldBeSkipped(err) {
				continue
			}
			return nil, fmt.Errorf("GetTimeSlotFrom - Error getting time slot from graph api. - %w", err)
		}
		deDupTimeSlots = append(deDupTimeSlots, timeSlots...)
	}
	return deDupTimeSlots, nil
}

func urlShouldBeSkipped(err error) bool {
	graphIsMissing := &GraphNotVisible{}
	invalidGraphIntervals := &InvalidGraphApiIntervals{}
	return errors.As(err, &graphIsMissing) || errors.As(err, &invalidGraphIntervals)
}

func GetTimeSlotFrom(requestUrl string) ([]string, error) { // this will be returning a string slice later on.
	response, err := getResponseFromGraphApi(requestUrl)
	if err != nil {
		return []string{}, err
	}
	deserializedResponse, err := deserializeGraphApiResponse(response)
	if err != nil {
		return []string{}, fmt.Errorf("deserializeGraphApiResponse - Error deserializing response. - %w", err)
	}

	if !graphIsVisible(deserializedResponse) {
		return []string{}, &GraphNotVisible{}
	}

	if timeIntervalsAreIdentical(deserializedResponse) {
		return []string{}, &InvalidGraphApiIntervals{}
	}

	timeSlots := time.GetUnixStampsInBetweenTimesAsString(deserializedResponse.Intervals[0].From, deserializedResponse.Intervals[0].To)
	return timeSlots, nil
}

func timeIntervalsAreIdentical(deserializedResponse *responseStructures.RelevantIndex) bool {
	return deserializedResponse.Intervals[0].From == deserializedResponse.Intervals[0].To
}

func graphIsVisible(deserializedResponse *responseStructures.RelevantIndex) bool {
	return deserializedResponse.Intervals[0].Color != "transparent"
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

var sendRequestToGraphApi = func(requestHandler *http.Request) (*http.Response, error) {
	response, err := sendRequest(requestHandler)
	return response, err
}

func deserializeGraphApiResponse(responseBuffer []byte) (*responseStructures.RelevantIndex, error) {
	deserializedType := responseStructures.GraphApiResponse{}
	deserializedResponse, err := deserializeResponse(responseBuffer, &deserializedType)
	if err != nil {
		return &responseStructures.RelevantIndex{}, err
	}
	result := deserializedResponse.(*responseStructures.GraphApiResponse)
	return &(*result)[0], nil
}
