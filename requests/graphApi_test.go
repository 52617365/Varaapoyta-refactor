package requests

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
	"varaapoyta-backend-refactor/responseStructures"
	"varaapoyta-backend-refactor/time"
)

func TestGetResponseFromGraphApi(t *testing.T) {
	mockRequestResult(`{"name": "test"}`)

	json := `{"name": "test"}`
	expectedResponseBuffer := io.NopCloser(bytes.NewReader([]byte(json))) // making new buffer because a singular one can't be read twice.
	expectedResponse, _ := io.ReadAll(expectedResponseBuffer)

	response, err := getResponseFromGraphApi("test")

	if err != nil {
		t.Errorf("getResponseFromGraphApi - Error getting response from mock graph api.")
	}
	if string(response) != string(expectedResponse) {
		t.Errorf("getResponseFromGraphApi - Expected response to be %s but got %s", expectedResponse, response)
	}
}

func TestGetGraphApiTimeSlotsFrom(t *testing.T) {
	t.Skip("Not implemented yet")
}

func TestGetTimeSlotFromReturnsGraphNotVisible(t *testing.T) {
	mockRequestResult(`[{"name": "Stone's","intervals":[{"from":1660330800000,"to":1660330800000,"color":"transparent"}],"id":281}]`)

	_, err := GetTimeSlotsFrom("test")
	graphNotVisible := &GraphNotVisible{}
	if !errors.As(err, &graphNotVisible) {
		t.Errorf("getTimeSlotFrom - Expected error to be of type GraphNotVisible but it wasn't.")
	}
}
func TestGetTimeSlotFromReturnsInvalidGraphApiIntervals(t *testing.T) {
	mockRequestResult(`[{"name": "Stone's","intervals":[{"from":1660330800000,"to":1660330800000,"color":""}],"id":281}]`)

	_, err := GetTimeSlotsFrom("test")
	invalidGraphApiIntervals := &InvalidGraphApiIntervals{}
	if !errors.As(err, &invalidGraphApiIntervals) {
		t.Errorf("getTimeSlotFrom - Expected error to be of type InvalidGraphApiIntervals but it wasn't.")
	}
}
func TestGetTimeSlotFrom(t *testing.T) {
	mockRequestResult(`[{"name": "Stone's","intervals":[{"from":1660330800000,"to":1670875200000,"color":""}],"id":281}]`)
	mockUnixTimeStampsBetweenTimesAsString([]string{"1900", "1915", "1930", "1945", "2000"})

	expectedTimeSlots := []string{"1900", "1915", "1930", "1945", "2000"}
	timeSlots, err := GetTimeSlotsFrom("test")
	if err != nil {
		t.Errorf("getTimeSlotFrom - Error getting time slots from mock graph api.")
	}
	if len(timeSlots) != len(expectedTimeSlots) {
		t.Errorf("getTimeSlotFrom - Expected time slots length to be %d but got %d", len(expectedTimeSlots), len(timeSlots))
	}
	for i, timeSlot := range timeSlots {
		if timeSlot != expectedTimeSlots[i] {
			t.Errorf("getTimeSlotFrom - Expected time slot to be %s but got %s", expectedTimeSlots[i], timeSlot)
		}
	}
}

func TestUrlShouldBeSkipped(t *testing.T) {
	if !urlShouldBeSkipped(&GraphNotVisible{}) {
		t.Errorf("urlShouldBeSkipped - Expected url to be skipped but it wasn't.")
	}
	if !urlShouldBeSkipped(&InvalidGraphApiIntervals{}) {
		t.Errorf("urlShouldBeSkipped - Expected url to be skipped but it wasn't.")
	}
	if urlShouldBeSkipped(errors.New("test")) {
		t.Errorf("urlShouldBeSkipped - Expected url to not be skipped but it was.")
	}
}

func TestTimeIntervalsAreIdenticalReturnsTrue(t *testing.T) {
	r := &responseStructures.RelevantIndex{
		Name:      "test",
		Intervals: responseStructures.Intervals{{From: 22222, To: 22222}},
		ID:        2,
	}

	if !timeIntervalsAreIdentical(r) {
		t.Errorf("timeIntervalsAreIdentical - Expected time intervals to be identical but they weren't.")
	}
}

func TestTimeIntervalsAreIdenticalReturnsFalse(t *testing.T) {
	r := &responseStructures.RelevantIndex{
		Name:      "test",
		Intervals: responseStructures.Intervals{{From: 22222, To: 55555}},
		ID:        2,
	}

	if timeIntervalsAreIdentical(r) {
		t.Errorf("timeIntervalsAreIdentical - Expected time intervals to NOT be identical but they were.")
	}
}

func TestGraphIsVisibleReturnsTrue(t *testing.T) {
	r := &responseStructures.RelevantIndex{
		Name:      "test",
		Intervals: responseStructures.Intervals{{From: 22222, To: 55555, Color: ""}},
		ID:        2,
	}
	if !graphIsVisible(r) {
		t.Errorf("graphIsVisible - Expected graph to be visible but it wasn't.")
	}
}

func TestGraphIsVisibleReturnsFalse(t *testing.T) {
	r := &responseStructures.RelevantIndex{
		Name:      "test",
		Intervals: responseStructures.Intervals{{From: 22222, To: 55555, Color: "transparent"}},
		ID:        2,
	}
	if graphIsVisible(r) {
		t.Errorf("graphIsVisible - Expected graph to NOT be visible but it was.")
	}
}

func TestGetCurrentTimeInUnixMs(t *testing.T) {
	currentTime := time.GetCurrentTimeInUnixMs()
	if currentTime < 0 {
		t.Errorf("getCurrentTimeInUnixMs - Expected current time to be positive but it wasn't.")
	}
}

func mockRequestResult(returnValue string) {
	mockResponseBuffer := io.NopCloser(bytes.NewReader([]byte(returnValue)))

	sendRequest = func(requestHandler *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       mockResponseBuffer,
		}, nil
	}
}

func mockUnixTimeStampsBetweenTimesAsString(returnValue []string) {
	time.GetUnixStampsInBetweenTimesAsString = func(fromMs int64, toMs int64) []string {
		return returnValue
	}
}
