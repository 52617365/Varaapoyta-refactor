package time

import "testing"

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
