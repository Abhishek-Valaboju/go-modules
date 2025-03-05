package main

import (
	"log"
	"net/http"
	"task-manager/controller/handler"
	"task-manager/controller/middleware"
)

func main() {
	mux := http.NewServeMux()
	port := ":8080"
	mux.HandleFunc("/", handler.GetHandler)
	mux.HandleFunc("/add", handler.GetHandler)
	log.Printf("Server running in %s", port)

	err := http.ListenAndServe(port, middleware.LoggingMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
