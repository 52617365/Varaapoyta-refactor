package main

import (
	"fmt"
	"varaapoyta-backend-refactor/requests"
)

// https://s-varaukset.fi/online/reservation/fi/72?_ga=2.161416895.382807502.1612853101-189045693.1611044564
// Regex:

func main() {
	graphApiTimeSlot, _ := requests.GetTimeSlotFrom("https://s-varaukset.fi/online/reservation/fi/72?_ga=2.161416895.382807502.1612853101-189045693.1611044564")
	fmt.Println(graphApiTimeSlot)
}
