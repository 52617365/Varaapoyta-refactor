package time

import (
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

func TestGetSlotsFromTheFuture(t *testing.T) {
	currentHour := GetCurrentHour()
	slotsInTheFuture := GetSlotsFromTheFuture()
	for _, slot := range slotsInTheFuture {
		if slot < currentHour {
			t.Errorf("slot %d is in the past", slot)
		}
	}
}

func TestSlotIsInFuture(t *testing.T) {
	currentHour := GetCurrentHour()
	slot := currentHour - 1
	if slotIsInFuture(slot, currentHour) {
		t.Errorf("slot: %d is in the future when it should be in the past", slot)
	}
}

func TestConvertUnixToTime(t *testing.T) {
	unix := int64(1625744000000)
	actualTime := convertUnixToFinnishTime(unix)

	expectedTime := time.Date(2021, time.July, 8, 0, 0, 0, 0, time.UTC)
	if actualTime.Day() != expectedTime.Day() {
		t.Errorf("expected %d, got %d", expectedTime.Day(), actualTime.Day())
	}
	if actualTime.Month() != expectedTime.Month() {
		t.Errorf("expected %d, got %d", expectedTime.Month(), actualTime.Month())
	}
	if actualTime.Year() != expectedTime.Year() {
		t.Errorf("expected %d, got %d", expectedTime.Year(), actualTime.Year())
	}
}

func TestGetTimeDifferenceBetweenTwoTimes(t *testing.T) {
	startTime := time.Date(2021, time.July, 8, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2021, time.July, 8, 1, 0, 0, 0, time.UTC)
	actualTimeDifference := getTimeDifferenceBetweenTwoTimes(startTime, endTime)
	expectedTimeDifference := time.Hour
	if actualTimeDifference != expectedTimeDifference {
		t.Errorf("expected %d, got %d", expectedTimeDifference, actualTimeDifference)
	}
}

func TestSpawnUnixTimeIntervals(t *testing.T) {
	actualUnixTimeIntervals := spawnUnixTimeIntervals()
	if len(actualUnixTimeIntervals) != 96 {
		t.Errorf("expected %d, got %d", 96, len(actualUnixTimeIntervals))
	}
}

func TestGetUnixStampsInbetweenTimes(t *testing.T) {
	spawnUnixTimeIntervals = func() []int64 {
		return []int64{1625743000000, 1625744000001, 1625744000002, 1625744000003}
	}
	from := int64(1625744000000)
	to := int64(1625747600000)
	actualUnixStampsInbetweenTimes := getUnixStampsInbetweenTimes(from, to)
	for _, time := range actualUnixStampsInbetweenTimes {
		if time < from || time > to {
			t.Errorf("time %d is not inbetween %d and %d", time, from, to)
		}
	}
	if len(actualUnixStampsInbetweenTimes) != 3 {
		t.Errorf("expected %d, got %d", 3, len(actualUnixStampsInbetweenTimes))
	}
}

func TestGetUnixStampsInBetweenTimesAsString(t *testing.T) {
	spawnUnixTimeIntervals = func() []int64 {
		return []int64{1625740000000, 1625742000000, 1625742000000 + 1800000, 1625742000000 + 3600000}
	}
	from := int64(1625741000000)
	to := int64(9999999999999)

	actualUnixStampsInbetweenTimes := GetUnixStampsInBetweenTimesAsString(from, to)

	expectedTimes := []string{"1300", "1330", "1400"}

	for _, time := range actualUnixStampsInbetweenTimes {
		if !slices.Contains(expectedTimes, time) {
			t.Errorf("expectedTimes does not contain %s", time)
		}
	}

	if len(actualUnixStampsInbetweenTimes) != 3 {
		t.Errorf("expected %d, got %d", 3, len(actualUnixStampsInbetweenTimes))
	}

}

func TestExtractUnwantedTimeSlots(t *testing.T) {
	timeSlots := []string{"1300", "1315", "1330", "1345", "1400", "1415", "1430", "1445"}
	kitchenClosingTime := "14:45"

	wantedTimeSlots := ExtractUnwantedTimeSlots(timeSlots, kitchenClosingTime)
	expectedTimeSlots := []string{"1300", "1315", "1330", "1345"}

	if len(wantedTimeSlots) != len(expectedTimeSlots) {
		t.Errorf("expected %d, got %d", len(expectedTimeSlots), len(wantedTimeSlots))
	}
	for _, time := range expectedTimeSlots {
		if !slices.Contains(wantedTimeSlots, time) {
			t.Errorf("wantedTimeSlots does not contain %s", time)
		}
	}
}

func TestExtractUnwantedTimeSlotsReturnsEmpty(t *testing.T) {
	timeSlots := []string{"1300", "1315", "1330", "1345", "1400", "1415", "1430", "1445"}
	kitchenClosingTime := "13:45"

	wantedTimeSlots := ExtractUnwantedTimeSlots(timeSlots, kitchenClosingTime)

	if len(wantedTimeSlots) != 0 {
		t.Errorf("expected %d, got %d", 0, len(wantedTimeSlots))
	}
}
