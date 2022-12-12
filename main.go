package main

import (
	"fmt"
	"varaapoyta-backend-refactor/requests"
)

// https://s-varaukset.fi/online/reservation/fi/72?_ga=2.161416895.382807502.1612853101-189045693.1611044564
// Regex:

func main() {
	restaurants, err := requests.GetRestaurants()
	if err != nil {
		panic(err)
	}
	for _, restaurant := range restaurants {
		restaurantId := restaurant.ReservationPageID
		urls := requests.GetGraphApiUrls(restaurantId)
		for _, url := range urls {
			timeSlots, err := requests.GetGraphApiTimeSlotsFrom(url)
			if err != nil {
				panic(err)
			}
			fmt.Println(timeSlots)
		}
	}
}
