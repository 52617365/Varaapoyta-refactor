package main

import (
	"fmt"
	"varaapoyta-backend-refactor/requests"
)

func main() {
	timeSlots, _ := requests.GetGraphApiTimeSlotsFrom("1685")
	fmt.Println(timeSlots)
}
