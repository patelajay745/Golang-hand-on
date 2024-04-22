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

// GetInventoryHandler retrieves all inventory items from the database and returns them as JSON.
func GetInventoryHandler(w http.ResponseWriter, r *http.Request) {
	inventoryItems := models.GetAllInventory()

	// Encode the inventory items into JSON format and send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(inventoryItems); err != nil {
		// If there's an error encoding the JSON, return an internal server error
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

}

func CreateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize an empty InventoryItem struct to hold the data from the request body
	newInventory := &models.InventoryItem{}

	// Parse the request body and populate the newInventory struct with the received data
	utils.ParseBody(r, newInventory)

	createdInventory := newInventory.CreateInventory()

	// Encode the created supplier into JSON format and send it in the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Set the status code to 201 (Created)
	json.NewEncoder(w).Encode(createdInventory)
}

func DeleteInventoryHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the inventory item ID from the request parameters.
	params := mux.Vars(r)
	inventoryID := params["id"]
	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		fmt.Println("Error while parsing:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = models.DeleteInventoryItem(id)
	if err != nil {
		fmt.Println("Error deleting Inventory:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Inventory item deleted successfully"))

}

func GetInventoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the request parameters
	params := mux.Vars(r)
	InventoryID := params["id"]
	id, err := strconv.Atoi(InventoryID)
	if err != nil {
		fmt.Println("Error while parsing the supplier ID:", err)
		http.Error(w, "Invalid Inventory ID", http.StatusBadRequest)
		return
	}

	// Fetch details of the Inventory from the database
	inventoryDetails, err := models.GetInventoryByID(id)
	if err != nil {
		http.Error(w, "Inventory not found", http.StatusNotFound)
		return
	}

	// Encode and send the details in the response
	json.NewEncoder(w).Encode(inventoryDetails)

}

// UpdateInventoryHandler updates an existing Inventory item.
func UpdateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	// Parse incoming data
	var updatedInventory = &models.InventoryItem{}
	utils.ParseBody(r, updatedInventory)

	// Get the ID to update
	params := mux.Vars(r)
	inventoryID := params["id"]

	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		fmt.Println("Error while parsing ID:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve inventory item details from the database
	inventoryDetails, err := models.GetInventoryByID(id)
	if err != nil {
		http.Error(w, "Inventory not found", http.StatusNotFound)
		return
	}

	// Update Inventory details if provided in the request
	if updatedInventory.Name != "" {
		inventoryDetails.Name = updatedInventory.Name
	}
	if updatedInventory.Description != "" {
		inventoryDetails.Description = updatedInventory.Description
	}
	if updatedInventory.CategoryID != 0 {
		inventoryDetails.CategoryID = updatedInventory.CategoryID
	}
	if updatedInventory.SupplierID != 0 {
		inventoryDetails.SupplierID = updatedInventory.SupplierID
	}
	if updatedInventory.Quantity != 0 {
		inventoryDetails.Quantity = updatedInventory.Quantity
	}
	if updatedInventory.Price != 0 {
		inventoryDetails.Price = updatedInventory.Price
	}
	if updatedInventory.Status != "" {
		inventoryDetails.Status = updatedInventory.Status
	}

	fmt.Println(inventoryDetails)

	// Save the updated Inventory details
	err = inventoryDetails.UpdateInventroryItem()
	if err != nil {
		fmt.Println("Error updating Inventory:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Inventory updated successfully")
}
