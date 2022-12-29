package main

import (
	"varaapoyta-backend-refactor/endpoints"
)

// TODO: the program terminates if it's used with "Helsinki", we are for sure missing some potential check that can only happen with restaurants from Helsinki.
func main() {
    endpoints.InitApi()
}
