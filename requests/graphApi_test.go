package requests

import "testing"

func TestGetUrl(t *testing.T) {
	restaurantId := 1
	timeSlot := "0800"
	expectedUrl := "https://s-varaukset.fi/api/recommendations/slot/1/TODO/0800/1"
	actualUrl := getUrl(restaurantId, timeSlot)
	if actualUrl != expectedUrl {
		t.Errorf("getUrl - Expected %s, got %s", expectedUrl, actualUrl)
	}
}

func TestGetUrls(t *testing.T) {
	restaurantId := 1
	expectedUrls := []string{
		"https://s-varaukset.fi/api/recommendations/slot/1/TODO/0800/1",
		"https://s-varaukset.fi/api/recommendations/slot/1/TODO/1200/1",
		"https://s-varaukset.fi/api/recommendations/slot/1/TODO/1600/1",
		"https://s-varaukset.fi/api/recommendations/slot/1/TODO/2000/1",
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
