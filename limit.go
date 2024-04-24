package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func ratelimit(next func(writer http.ResponseWriter, request *http.Request)) http.Handler {
	// creating a rate limiter
	ratelimiter := rate.NewLimiter(2, 4)
	// handler func
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// checking if raterlimiter is bounced
		if !ratelimiter.Allow() {
			// Creating a message
			msg := Message{
				Status: "Too many requests",
				Body:   "The API got too many requests too frequently",
			}
			// sending the response
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(msg)
			return
		} else {
			// running the next func if rate limiter allows
			next(w, r)
			return
		}
	})
}
