package main

import (
	"fmt"
	"varaapoyta-backend-refactor/requests"
)

func main() {
	restaurants, _ := requests.GetRestaurants("Helsinki")
	for _, restaurant := range restaurants {
		// TODO: we want to use goroutines here.
		timeSlots, _ := requests.GetGraphApiTimeSlotsFrom(restaurant.ReservationPageID)
		fmt.Println(timeSlots)
	}
}
