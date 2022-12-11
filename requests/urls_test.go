package requests

import (
	"fmt"
	"golang.org/x/exp/slices"
	"testing"
	"varaapoyta-backend-refactor/date"
	"varaapoyta-backend-refactor/time"
)

func TestGetUrl(t *testing.T) {
	restaurantId := "1"
	currentDate := date.GetCurrentDate()
	timeSlotHour := 8
	expectedUrl := fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/0800/1", currentDate)
	actualUrl := getUrl(restaurantId, timeSlotHour)
	if actualUrl != expectedUrl {
		t.Errorf("getUrl - Expected %s, got %s", expectedUrl, actualUrl)
	}
}

func TestGetUrls(t *testing.T) {
	restaurantId := "1"
	expectedUrls := getExpectedUrls()
	actualUrls := GetUrls(restaurantId)

	if len(actualUrls) != len(expectedUrls) {
		t.Errorf("GetUrls - Expected %d urls, got %d", len(expectedUrls), len(actualUrls))
	}

	for _, expectedUrl := range expectedUrls {
		if !slices.Contains(actualUrls, expectedUrl) {
			t.Errorf("GetUrls - Expected returned urls to contain %s but it did not", expectedUrl)
		}
	}
}

func getExpectedUrls() []string {
	currentDate := date.GetCurrentDate()
	currentHour := time.GetCurrentHour()

	expectedUrls := make([]string, 0, 4)

	if currentHour <= 20 {
		expectedUrls = append(expectedUrls, fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/2000/1", currentDate))
	}
	if currentHour <= 14 {
		expectedUrls = append(expectedUrls, fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/1400/1", currentDate))
	}
	if currentHour <= 8 {
		expectedUrls = append(expectedUrls, fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/0800/1", currentDate))
	}
	if currentHour <= 2 {
		expectedUrls = append(expectedUrls, fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/0200/1", currentDate))
	}
	return expectedUrls
}
