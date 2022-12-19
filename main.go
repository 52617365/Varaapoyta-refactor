package main

import (
	"fmt"
	"varaapoyta-backend-refactor/requests"
)

func main() {
	timeSlots, _ := requests.GetGraphApiTimeSlotsFrom("1679")
	fmt.Println(timeSlots)
}
