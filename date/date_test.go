package date

import (
	"strings"
	"testing"
)

func TestGetCurrentDate(t *testing.T) {
	date := GetCurrentDate()
	if strings.Count(date, "-") != 2 {
		t.Errorf("GetCurrentDate - Expected date to have 2 dashes but it had %d.", strings.Count(date, "-"))
	}
	if len(date) != 10 {
		t.Errorf("GetCurrentDate - Expected date to be 4 characters long but it was %d", len(date))
	}
}
