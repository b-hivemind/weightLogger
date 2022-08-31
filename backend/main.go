package main

import (
	"fmt"

	"bhavdeep.me/weight_logger/pkg/api"
)

func main() {
	api.HandleRequests()
	fmt.Println("Now listening on 10000")
}
