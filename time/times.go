package time

import (
	"log"
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

var GetCurrentTimeInUnixMs = func() int64 {
	return time.Now().UnixMilli()
}

func slotIsInFuture(slot int, currentHour int) bool {
	return slot >= currentHour
}

func convertUnixToFinnishTime(unix int64) time.Time {
	t := time.UnixMilli(unix).UTC().Add(2 * time.Hour)
	return t
}

type RelativeTime struct{
	Hour int `json:"hour"`
	Minute int `json:"minute"`
}
func CalcRelativeTimeToFromCurrentTime(closingTime string) *RelativeTime {
	closingTimeToTimeType := FormatStringTimeToTimeType(closingTime)
	currentTimeToTimeType := getCurrentTimeMatchingClosingFormat()

	relative := getTimeDifferenceBetweenTwoTimes(currentTimeToTimeType, closingTimeToTimeType)

	result := RelativeTime{
		Hour: int(relative.Seconds()) / 3600,
		Minute: int(relative.Seconds() / 60) % 60,
	}
	return &result
}

func getCurrentTimeMatchingClosingFormat() time.Time {
	var currentTime string = getFormattedCurrentTime()
	currentTimeToTimeType := FormatStringTimeToTimeType(currentTime)
	return currentTimeToTimeType
}

func getFormattedCurrentTime() string {
	currentTime := time.Now().Format("15:04")
	return currentTime
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

var GetUnixStampsInBetweenTimesAsString = func(fromMs int64, toMs int64) []string {
	unixStampsInbetweenTimes := getUnixStampsInbetweenTimes(fromMs, toMs)

	var unixStampsInbetweenTimesAsString []string
	for _, unixStamp := range unixStampsInbetweenTimes {
		unixStampsInbetweenTimesAsString = append(unixStampsInbetweenTimesAsString, convertUnixToFinnishTime(unixStamp).Format("1504"))
	}
	return unixStampsInbetweenTimesAsString
}

func getUnixStampsInbetweenTimes(from int64, to int64) []int64 {
	timeIntervals := spawnUnixTimeIntervals()
	unixStampsInbetweenTimes := make([]int64, 0, 96)
	for _, unixStamp := range timeIntervals {
		if unixStamp >= from && unixStamp <= to {
			unixStampsInbetweenTimes = append(unixStampsInbetweenTimes, unixStamp)
		}
	}
	return unixStampsInbetweenTimes
}

func ExtractUnwantedTimeSlots(timeSlots []string, kitchenClosingTime string) []string {
	if timeSlots == nil || kitchenClosingTime == "" {
		// This hopefully never gets hit.
		log.Fatal("timeSlots or kitchenClosingTime is nil")
	}

	closingTime := FormatStringTimeToTimeType(kitchenClosingTime)

	timeSlotsNotInClosingRange := make([]string, 0, len(timeSlots))

	for _, timeSlot := range timeSlots {
		timeSlotTime := convertTimeSlotToTime(timeSlot)
		if !isInClosingRange(timeSlotTime, closingTime) {
			timeSlotsNotInClosingRange = append(timeSlotsNotInClosingRange, timeSlot)
		}
	}
	return timeSlotsNotInClosingRange
}

func convertTimeSlotToTime(timeSlot string) time.Time {
	// time slots are stored as "1400", so we need to format it to "14:00" before we can convert it to a time.Time.
	formattedTimeSlot := formatTimeWithColon(timeSlot)
	t := FormatStringTimeToTimeType(formattedTimeSlot)
	return t
}

func formatTimeWithColon(timeSlot string) string {
	formattedTimeSlot := timeSlot[:2] + ":" + timeSlot[2:4]
	return formattedTimeSlot
}

func FormatStringTimeToTimeType(timeString string) time.Time {
	const layout = "15:04"
	t, _ := time.Parse(layout, timeString)
	return t
}

func isInClosingRange(timeSlot time.Time, kitchenClosingTime time.Time) bool {
	kitchenClosingTime = kitchenClosingTime.Add(-1 * time.Hour) // minus 1h from the time bcuz restaurants don't take reservations for the last hour of the day.
	return timeSlot.After(kitchenClosingTime)
}
