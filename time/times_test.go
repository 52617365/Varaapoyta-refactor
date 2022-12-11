package time

import (
	"testing"
	"time"
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
	actualTime := convertUnixToTime(unix)

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
