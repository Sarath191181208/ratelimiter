package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func Ping(writer http.ResponseWriter, request *http.Request) {
	// Writing headers
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// creating the message
	msg := Message{
		Status: "Success",
		Body:   "This is something I am returing man!",
	}

	// to json
	err := json.NewEncoder(writer).Encode(msg)
	// handle write err
	if err != nil {
		return
	}
}

func main() {
	// defining port
	const portNum = 8080
	port := ":" + strconv.Itoa(portNum)

	// adding the handlers
	http.Handle("/ping", ratelimit(Ping))

	// starting the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Can't start server on port: " + strconv.Itoa(portNum))
	}
}
