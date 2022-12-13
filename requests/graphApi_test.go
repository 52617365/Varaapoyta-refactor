package requests

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
)

func TestGetResponseFromGraphApi(t *testing.T) {
	mockRequestResult()

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

func TestGetTimeSlotFromReturnsRightErrors(t *testing.T) {
	t.Skip("Not implemented yet")
}

func TestUrlShouldBeSkipped(t *testing.T) {
	err := &GraphNotVisible{}
	if !urlShouldBeSkipped(err) {
		t.Errorf("urlShouldBeSkipped - Expected url to be skipped but it wasn't.")
	}
	err2 := &InvalidGraphApiIntervals{}
	if !urlShouldBeSkipped(err2) {
		t.Errorf("urlShouldBeSkipped - Expected url to be skipped but it wasn't.")
	}

	err3 := errors.New("test")
	if urlShouldBeSkipped(err3) {
		t.Errorf("urlShouldBeSkipped - Expected url to not be skipped but it was.")
	}
}

func mockRequestResult() {
	json := `{"name": "test"}`
	mockResponseBuffer := io.NopCloser(bytes.NewReader([]byte(json)))

	sendRequest = func(requestHandler *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       mockResponseBuffer,
		}, nil
	}
}
