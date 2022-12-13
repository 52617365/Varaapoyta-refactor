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
	var allTimeSlots []string
	for _, url := range urls {
		timeSlots, err := GetTimeSlotFrom(url)
		if err != nil {
			if urlShouldBeSkipped(err) {
				continue
			}
			return nil, fmt.Errorf("GetTimeSlotFrom - Error getting time slot from graph api. - %w", err)
		}
		allTimeSlots = append(allTimeSlots, timeSlots...)
	}
	return allTimeSlots, nil
}

var GetTimeSlotFrom = func(requestUrl string) ([]string, error) { // this will be returning a string slice later on.
	response, err := getResponseFromGraphApi(requestUrl)
	if err != nil {
		return []string{}, err
	}
	deserializedResponse, err := deserializeGraphApiResponse(response)
	if err != nil {
		return []string{}, fmt.Errorf("deserializeGraphApiResponse - Error deserializing response. - %w", err)
	}

	err = errOnInvalidData(deserializedResponse)
	if err != nil {
		return []string{}, err
	}

	timeSlots := time.GetUnixStampsInBetweenTimesAsString(deserializedResponse.Intervals[0].From, deserializedResponse.Intervals[0].To)
	return timeSlots, nil
}

func errOnInvalidData(deserializedResponse *responseStructures.RelevantIndex) error {
	if !graphIsVisible(deserializedResponse) {
		return &GraphNotVisible{}
	}

	if timeIntervalsAreIdentical(deserializedResponse) {
		return &InvalidGraphApiIntervals{}
	}
	return nil
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

func sendRequestToGraphApi(requestHandler *http.Request) (*http.Response, error) {
	response, err := sendRequest(requestHandler)
	return response, err
}

var deserializeGraphApiResponse = func(responseBuffer []byte) (*responseStructures.RelevantIndex, error) {
	deserializedResponse, err := deserializeResponse(responseBuffer, &responseStructures.GraphApiResponse{})
	if err != nil {
		return &responseStructures.RelevantIndex{}, err
	}
	result := deserializedResponse.(*responseStructures.GraphApiResponse)
	return &(*result)[0], nil
}

func urlShouldBeSkipped(err error) bool {
	graphIsMissing := &GraphNotVisible{}
	invalidGraphIntervals := &InvalidGraphApiIntervals{}
	return errors.As(err, &graphIsMissing) || errors.As(err, &invalidGraphIntervals)
}
