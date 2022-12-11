package requests

import (
	"fmt"
	"varaapoyta-backend-refactor/date"
	"varaapoyta-backend-refactor/time"
)

func getUrls(restaurantId int) []string {
	var urls []string
	for _, timeSlot := range time.GetSlotsFromTheFuture() {
		url := getUrl(restaurantId, timeSlot)
		urls = append(urls, url)
	}
	return urls
}
func getUrl(restaurantId int, timeSlot int) string {
	currentDate := date.GetCurrentDate()
	formattedTimeSlot := formatTimeSlotFrom(timeSlot)

	url := fmt.Sprintf(`https://s-varaukset.fi/api/recommendations/slot/%d/%s/%s/1`, restaurantId, currentDate, formattedTimeSlot)
	return url
}

func formatTimeSlotFrom(timeSlot int) string {
	if timeSlotIsTwoDigits(timeSlot) {
		return fmt.Sprintf("%d00", timeSlot)
	} else {
		return fmt.Sprintf("0%d00", timeSlot)
	}
}

func timeSlotIsTwoDigits(timeSlot int) bool {
	return timeSlot > 9
}
