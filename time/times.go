package time

import (
	"time"
)

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

func getTimeDifferenceBetweenTwoTimes(startTime time.Time, endTime time.Time) time.Duration {
	return endTime.Sub(startTime)
}

var spawnUnixTimeIntervals = func() []int64 {
	currentTime := time.Now()

	timeIntervals := make([]int64, 0, 96)
	for i := 0; i < 24; i++ {
		timeIntervals = append(timeIntervals, time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0, 0, time.UTC).UnixMilli())
		timeIntervals = append(timeIntervals, time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 15, 0, 0, time.UTC).UnixMilli())
		timeIntervals = append(timeIntervals, time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 30, 0, 0, time.UTC).UnixMilli())
		timeIntervals = append(timeIntervals, time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 45, 0, 0, time.UTC).UnixMilli())
	}
	return timeIntervals
}

func getUnixStampsInbetweenTimes(from int64, to int64) []int64 {
	timeIntervals := spawnUnixTimeIntervals()
	var unixStampsInbetweenTimes []int64
	for _, unixStamp := range timeIntervals {
		if unixStamp >= from && unixStamp <= to {
			unixStampsInbetweenTimes = append(unixStampsInbetweenTimes, unixStamp)
		}
	}
	return unixStampsInbetweenTimes
}
