package time

import "time"

func GetSlotsFromTheFuture() []int {
	var GraphSlotHoursForTheDay = [...]int{2, 8, 14, 20}
	slotsInTheFuture := make([]int, 0, len(GraphSlotHoursForTheDay))

	var currentHour = GetCurrentHour()
	for _, slot := range GraphSlotHoursForTheDay {
		if slotIsInFuture(slot, currentHour) {
			slotsInTheFuture = append(slotsInTheFuture, slot)
		}
	}
	return slotsInTheFuture
}

func GetCurrentHour() int {
	currentHour := time.Now().Hour()
	return currentHour
}

func slotIsInFuture(slot int, currentHour int) bool {
	return slot >= currentHour
}

func convertUnixToTime(unix int64) time.Time {
	return time.UnixMilli(unix)
}
