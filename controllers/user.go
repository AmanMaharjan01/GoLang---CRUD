package controllers

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	// models package where User schema is defined
	"go-postgres/services" // models package where User schema is defined
	"net/http"             // used to access the request and response object of the api
	// package used to covert string into int type
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