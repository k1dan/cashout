package main

import (
	"github.com/gorilla/mux"
	"github.com/k1dan/cashout/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Payment).Methods("POST")
	router.HandleFunc("/callback", handler.GetCallback).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}