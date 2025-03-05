package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api/pkg"
)

func main() {
	store := pkg.NewInMemoryStore()
	handler := pkg.NewHandler(store)
	r := mux.NewRouter()
	r.HandleFunc("/items", handler.CreateItem).Methods("POST")
	r.HandleFunc("/items", handler.GetAllItems).Methods("GET")
	r.HandleFunc("/items/{id}", handler.GetItemByID).Methods("GET")
	r.HandleFunc("/items/{id}", handler.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", handler.DeleteItem).Methods("DELETE")
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
