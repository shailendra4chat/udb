package main

import (
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	Database.Create(&user)

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var users []User

	Database.Find(&users)

	json.NewEncoder(w).Encode(users)
}
