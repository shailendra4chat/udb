package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/user", CreateUser).Methods("POST")

	r.HandleFunc("/users", GetUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
