package requests

import (
	"fmt"
	"testing"
)

func TestGetUrl(t *testing.T) {
	restaurantId := 1
	currentDate := getCurrentDate()
	timeSlot := "0800"
	expectedUrl := fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/0800/1", currentDate)
	actualUrl := getUrl(restaurantId, timeSlot)
	if actualUrl != expectedUrl {
		t.Errorf("getUrl - Expected %s, got %s", expectedUrl, actualUrl)
	}
}

func TestGetUrls(t *testing.T) {
	restaurantId := 1
	currentDate := getCurrentDate()
	expectedUrls := []string{
		fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/0800/1", currentDate),
		fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/1200/1", currentDate),
		fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/1600/1", currentDate),
		fmt.Sprintf("https://s-varaukset.fi/api/recommendations/slot/1/%s/2000/1", currentDate),
	}
	actualUrls := getUrls(restaurantId)

	if len(actualUrls) != len(expectedUrls) {
		t.Errorf("getUrls - Expected %d urls, got %d", len(expectedUrls), len(actualUrls))
	}

	for i, actualUrl := range actualUrls {
		expectedUrl := expectedUrls[i]
		if actualUrl != expectedUrl {
			t.Errorf("getUrls - Expected %s, got %s", expectedUrl, actualUrl)
		}
	}
}
