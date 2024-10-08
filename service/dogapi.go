package service

import (
	"encoding/json"
	"net/http"
	"time"
)

type APIResponse struct {
	Facts   []string `json:"facts"`
	Success bool     `json:"success"`
}

func ServeDogApi(h *Hub, delay *int) {
	if delay == nil {
		d := DEFAULT_REQUEST_DELAY
		delay = &d
	}
	ticker := time.NewTicker(time.Duration(*delay) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			res, err := http.Get(h.apiUrl)
			if err != nil {
				h.err <- err
				continue
			}

			defer res.Body.Close()
			var message *APIResponse

			if err := json.NewDecoder(res.Body).Decode(&message); err != nil {
				h.err <- err
				continue
			}

			for _, v := range message.Facts {
				h.message <- &v
			}
		}
	}
}
