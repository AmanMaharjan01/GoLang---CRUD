package services

import (
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"go-postgres/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) response {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}


	insertID := repositories.InsertUser(user)

	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	return res
}

func GetUser(w http.ResponseWriter, r *http.Request) models.User{
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	user, err := repositories.GetUser(int64(id))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	return user
}

func GetAllUser(w http.ResponseWriter, r *http.Request) []models.User{

	users, err := repositories.GetAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	return users
}

func UpdateUser(w http.ResponseWriter, r *http.Request) response {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := repositories.UpdateUser(int64(id), user)

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	return res;
}

func DeleteUser(w http.ResponseWriter, r *http.Request)  response{

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := repositories.DeleteUser(int64(id))

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	return res;
}