package requests

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"varaapoyta-backend-refactor/responseStructures"
	"varaapoyta-backend-refactor/time"
)

type GraphTimeSlots struct {
	timeSlots []string
	err       error
}

func GetGraphApiTimeSlotsFrom(restaurantId string) ([]string, error) {
	urls := GetGraphApiUrlsFrom(restaurantId)

	wg := sync.WaitGroup{}

	graphTimeSlots := make(chan GraphTimeSlots, len(urls))

	for _, url := range urls {
		wg.Add(1)
		getTimeSlotsWithGoRoutine(url, graphTimeSlots, &wg)
	}

	wg.Wait()
	close(graphTimeSlots)

	allTimeSlots, err := syncGraphTimeSlots(graphTimeSlots)
	if err != nil {
		return nil, err
	}
	return allTimeSlots, nil
}

func getTimeSlotsWithGoRoutine(url string, graphTimeSlots chan GraphTimeSlots, wg *sync.WaitGroup) {
	go func(url string) {
		defer wg.Done()
		timeSlots, err := GetTimeSlotsFrom(url)
		if err != nil {
			handleTimeSlotErr(err, graphTimeSlots)
			return
		}
		graphTimeSlots <- GraphTimeSlots{timeSlots: timeSlots, err: nil}
	}(url)
}

func handleTimeSlotErr(err error, graphTimeSlots chan GraphTimeSlots) {
	if urlShouldBeSkipped(err) {
		graphTimeSlots <- GraphTimeSlots{timeSlots: nil, err: &UrlShouldBeSkipped{}}
		return
	}
	graphTimeSlots <- GraphTimeSlots{timeSlots: nil, err: fmt.Errorf("GetTimeSlotsFrom - Raflaamo graph api might be down. - %w", err)}
}

func syncGraphTimeSlots(graphTimeSlots chan GraphTimeSlots) ([]string, error) {
	syncedTimeSlots := make([]string, 0, 96)
	for timeSlot := range graphTimeSlots {
		if timeSlot.err != nil {
			urlShouldBeSkipped := &UrlShouldBeSkipped{}
			if errors.As(timeSlot.err, &urlShouldBeSkipped) {
				continue
			}
			return nil, timeSlot.err
		}
		syncedTimeSlots = append(syncedTimeSlots, timeSlot.timeSlots...)
	}
	return syncedTimeSlots, nil
}

func urlShouldBeSkipped(err error) bool {
	graphIsMissing := &GraphNotVisible{}
	invalidGraphIntervals := &InvalidGraphApiIntervals{}
	return errors.As(err, &graphIsMissing) || errors.As(err, &invalidGraphIntervals)
}

var GetTimeSlotsFrom = func(requestUrl string) ([]string, error) {
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
	fromTime := getFromAsCurrentTimeIfItsSmallerThan(deserializedResponse.Intervals[0].From)
	timeSlots := time.GetUnixStampsInBetweenTimesAsString(fromTime, deserializedResponse.Intervals[0].To)
	return timeSlots, nil
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

func errOnInvalidData(deserializedResponse *responseStructures.RelevantIndex) error {
	if !graphIsVisible(deserializedResponse) {
		return &GraphNotVisible{}
	}

	if timeIntervalsAreIdentical(deserializedResponse) {
		return &InvalidGraphApiIntervals{}
	}
	return nil
}
func graphIsVisible(deserializedResponse *responseStructures.RelevantIndex) bool {
	return deserializedResponse.Intervals[0].Color != "transparent"
}

func timeIntervalsAreIdentical(deserializedResponse *responseStructures.RelevantIndex) bool {
	return deserializedResponse.Intervals[0].From == deserializedResponse.Intervals[0].To
}

func getFromAsCurrentTimeIfItsSmallerThan(from int64) int64 {
	var resultFrom int64
	if currentTimeIsSmallerThanFrom(from) {
		resultFrom = time.GetCurrentTimeInUnixMs()
	} else {
		resultFrom = from
	}
	return resultFrom
}
func currentTimeIsSmallerThanFrom(from int64) bool {
	currentTimeUnix := time.GetCurrentTimeInUnixMs()
	return from > currentTimeUnix
}

var deserializeGraphApiResponse = func(responseBuffer []byte) (*responseStructures.RelevantIndex, error) {
	deserializedResponse, err := deserializeResponse(responseBuffer, &responseStructures.GraphApiResponse{})
	if err != nil {
		return &responseStructures.RelevantIndex{}, err
	}
	result := deserializedResponse.(*responseStructures.GraphApiResponse)
	return &(*result)[0], nil
}
