package main

import (
	"fmt"
	"varaapoyta-backend-refactor/requests"
)

func main() {
	restaurants, _ := requests.GetRestaurants("Rovaniemi")
	for _, restaurant := range restaurants {
		timeSlots, _ := requests.GetGraphApiTimeSlotsFrom(restaurant.ReservationPageID)
		fmt.Println(timeSlots)
	}
}
