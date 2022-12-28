package main

import (
	"fmt"
	"time"
	"varaapoyta-backend-refactor/requests"
)

func main() {
	// benchmark this code
	t := time.Now()
	restaurants, _ := requests.GetRestaurantsWithTimeSlots("Helsinki")
	fmt.Println(time.Since(t))
	for _, restaurant := range restaurants {
		fmt.Println(restaurant)
	}
}
