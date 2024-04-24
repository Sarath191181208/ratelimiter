package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func ratelimit(next func(writer http.ResponseWriter, request *http.Request)) http.Handler {
  ratelimiter := rate.NewLimiter(2, 4)
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if !ratelimiter.Allow(){
      msg := Message{
        Status: "Too many requests",
        Body: "The API got too many requests too frequently",
      }
      w.WriteHeader(http.StatusTooManyRequests)
      json.NewEncoder(w).Encode(msg)
      return
    } else{
      next(w, r)
      return 
    }
  })
}
