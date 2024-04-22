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

// GetSupplier retrieves all Suppliers from the database and returns them as JSON.
func GetSupplier(w http.ResponseWriter, r *http.Request) {
	suppliers := models.GetAllInventory()

	// Encode the suppliers items into JSON format and send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(suppliers); err != nil {
		// If there's an error encoding the JSON, return an internal server error
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	// Respond with a success status code
	w.WriteHeader(http.StatusOK)
}

func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	createSupplier := &models.Supplier{}
	utils.ParseBody(r, createSupplier)
	supplier := createSupplier.CreateSupplier()
	// Encode the created supplier into JSON format and send it in the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Set the status code to 201 (Created)
	json.NewEncoder(w).Encode(supplier)

}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	supId := params["id"]
	id, err := strconv.Atoi(supId)
	if err != nil {
		fmt.Println("Error while parsing:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = models.DeleteSupplier(id)
	if err != nil {
		fmt.Println("Error deleting supplier:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// UpdateSupplier updates an existing supplier.
func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	// Parse incoming data
	var updatedSupplier = &models.Supplier{}
	utils.ParseBody(r, updatedSupplier)

	// Get the ID to update
	params := mux.Vars(r)
	supID := params["id"]
	id, err := strconv.Atoi(supID)
	if err != nil {
		fmt.Println("Error while parsing ID:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve supplier details from the database
	supplierDetails, err := models.GetSupplierByID(id)
	if err != nil {
		http.Error(w, "Supplier not found", http.StatusNotFound)
		return
	}

	// Update supplier details if provided in the request
	if updatedSupplier.Name != "" {
		supplierDetails.Name = updatedSupplier.Name
	}
	if updatedSupplier.ContactPerson != "" {
		supplierDetails.ContactPerson = updatedSupplier.ContactPerson
	}
	if updatedSupplier.Email != "" {
		supplierDetails.Email = updatedSupplier.Email
	}
	if updatedSupplier.Phone != "" {
		supplierDetails.Phone = updatedSupplier.Phone
	}
	if updatedSupplier.Address != "" {
		supplierDetails.Address = updatedSupplier.Address
	}

	// Save the updated supplier details
	err = supplierDetails.UpdateSupplier()
	if err != nil {
		fmt.Println("Error updating supplier:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Supplier updated successfully")
}

// GetSupplierByID retrieves details of a supplier by ID.
func GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the request parameters
	params := mux.Vars(r)
	supplierID := params["id"]
	id, err := strconv.Atoi(supplierID)
	if err != nil {
		fmt.Println("Error while parsing the supplier ID:", err)
		http.Error(w, "Invalid Supplier ID", http.StatusBadRequest)
		return
	}

	// Fetch details of the supplier from the database
	supplierDetails, err := models.GetSupplierByID(id)
	if err != nil {
		http.Error(w, "Supplier not found", http.StatusNotFound)
		return
	}

	// Encode and send the details in the response
	json.NewEncoder(w).Encode(supplierDetails)
}
