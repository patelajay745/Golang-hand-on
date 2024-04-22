package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/models"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/utils"
)

var NewUser models.User

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userFromDB := models.GetAllUsers()
	res, _ := json.Marshal(userFromDB)
	w.Header().Set("Content-Type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	Id, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, _ := models.GetUserByID(Id)

	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	Id, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	models.DeleteUser(Id)

	w.Header().Set("Content-Type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	u := CreateUser.CreateUser()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var detailsFromUser = &models.User{}
	utils.ParseBody(r, detailsFromUser)

	vars := mux.Vars(r)

	userID := vars["id"]
	ID, err := strconv.Atoi(userID)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	detailsFromDB, db := models.GetUserByID(ID)

	if detailsFromUser.Username != "" {
		detailsFromDB.Username = detailsFromUser.Username
	}

	if detailsFromUser.Email != "" {
		detailsFromDB.Email = detailsFromUser.Email
	}

	if detailsFromUser.Password != "" {
		detailsFromDB.Password = detailsFromUser.Password
	}

	if detailsFromUser.Role != "" {
		detailsFromDB.Role = detailsFromUser.Role
	}

	db.Save(&detailsFromDB)

	res, _ := json.Marshal(detailsFromDB)

	w.Header().Set("Content-Type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
