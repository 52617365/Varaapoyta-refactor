package main

import (
	"fmt"
	"varaapoyta-backend-refactor/requests"
)

func main() {
	restaurants, _ := requests.GetRestaurantsWithTimeSlots("Helsinki")
	for _, restaurant := range restaurants {
		fmt.Println(restaurant)
	}
}
