package main

import (
	"dogapi/service"
	"fmt"
	"os"
	"strconv"
)

func main() {
	hub := service.NewHub()
	go hub.Listen()

	var delay *int
	if len(os.Args) > 1 {
		if val, err := strconv.Atoi(os.Args[1]); err != nil {
			fmt.Println("Delay is incorrect", val)
		} else {
			delay = &val
		}
	}
	service.ServeDogApi(hub, delay)
}
