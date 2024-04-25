package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/patelajay745/projects/03InventoryTracker/pkg/models"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/utils"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	credentials := &models.Credentials{}
	utils.ParseBody(r, credentials)

	// Encode the inventory items into JSON format and send the response
	w.Header().Set("Content-Type", "application/json")

	user, err := models.LoginUser(credentials.Username, credentials.Password)
	if err != nil {

		// If there's an error encoding the JSON, return an internal server error
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(int(user.ID), user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the created supplier into JSON format and send it in the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Set the status code to 201 (Created)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}
