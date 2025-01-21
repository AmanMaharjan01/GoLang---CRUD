package controllers

import (
	"encoding/json"
	"go-postgres/services"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	response := services.GetUser(w,r)

	// send the response
	json.NewEncoder(w).Encode(response)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {

	response := services.GetAllUser(w,r)

	// send the response
	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	response := services.CreateUser(w,r)

	// send the response
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	response := services.UpdateUser(w,r)

	// send the response
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	response := services.DeleteUser(w,r)

	// send the response
	json.NewEncoder(w).Encode(response)
}