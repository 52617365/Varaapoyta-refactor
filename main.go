package main

import (
	"fmt"
	"varaapoyta-backend-refactor/requests"
)

func main() {
	restaurants, err := requests.GetRestaurants()
	if err != nil {
		panic(err)
	}
	fmt.Println(restaurants)
}
