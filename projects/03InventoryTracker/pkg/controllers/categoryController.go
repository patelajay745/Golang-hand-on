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

func GetCategories(w http.ResponseWriter, r *http.Request) {

	categoryFromDB := models.GetAllCategories()
	json.NewEncoder(w).Encode(categoryFromDB)

}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	CreateCategory := &models.Category{}
	utils.ParseBody(r, CreateCategory)
	cat := CreateCategory.CreateCategory()
	json.NewEncoder(w).Encode(cat)

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	catId := params["id"]
	id, err := strconv.Atoi((catId))
	if err != nil {
		fmt.Println("Error while parsing ")
	}
	models.DeleteCategory(id)
	json.NewEncoder(w).Encode("Category is deleted")

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	var detailsFromUser = &models.Category{}
	utils.ParseBody(r, detailsFromUser)

	params := mux.Vars(r)
	catId := params["id"]
	id, err := strconv.Atoi((catId))
	if err != nil {
		fmt.Println("Error while parsing ")
	}

	detailsFromDB := models.GetCategoryByID(id)

	if detailsFromUser.Name != "" {
		detailsFromDB.Name = detailsFromUser.Name
	}

	if detailsFromUser.Description != "" {
		detailsFromDB.Description = detailsFromUser.Description
	}

	detailsFromDB.UpdateCategory()

	json.NewEncoder(w).Encode("Category is Updated")

}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	catId := params["id"]
	id, err := strconv.Atoi((catId))
	if err != nil {
		fmt.Println("Error while parsing ")
	}

	detailsFromDB := models.GetCategoryByID(id)
	json.NewEncoder(w).Encode(detailsFromDB)

}
