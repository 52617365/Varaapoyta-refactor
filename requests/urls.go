package requests

import (
	"fmt"
	"varaapoyta-backend-refactor/date"
	"varaapoyta-backend-refactor/time"
)

func getUrls(restaurantId int) []string {
	var urls []string
	for _, timeSlot := range time.GraphApiTimeslots {
		url := getUrl(restaurantId, timeSlot)
		urls = append(urls, url)
	}
	return urls
}
func getUrl(restaurantId int, timeSlot string) string {
	currentDate := date.GetCurrentDate()

	url := fmt.Sprintf(`https://s-varaukset.fi/api/recommendations/slot/%d/%s/%s/1`, restaurantId, currentDate, timeSlot)
	return url
}
