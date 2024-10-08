package service

import (
	"fmt"
)

type Hub struct {
	message chan *string
	err     chan error
	apiUrl  string
}

func NewHub() *Hub {
	return &Hub{make(chan *string), nil, "https://dog-api.kinduff.com/api/facts"}
}

func (h *Hub) Listen() {
	for {
		select {
		case message := <-h.message:
			fmt.Println(*message)
		case err := <-h.err:
			fmt.Println(err.Error())
		}
	}
}
