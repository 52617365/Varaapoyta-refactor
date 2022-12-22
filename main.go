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
	for _, restaurant := range restaurants {
		fmt.Println(restaurant)
	}
	fmt.Println(time.Since(t))
}
