package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/models"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/utils"
)

// GetTransactionHandler retrieves all transaction  from the database and returns them as JSON.
func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	transactions := models.GetAllTransaction()

	// Encode the transactions into JSON format and send the response
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		// If there's an error encoding the JSON, return an internal server error
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	// Respond with a success status code

}

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize an empty Transaction struct to hold the data from the request body
	newTransaction := &models.Transaction{}

	// Parse the request body and populate the newTransaction struct with the received data
	utils.ParseBody(r, newTransaction)

	// Print the parsed transaction data
	fmt.Println(newTransaction)

	createdTransaction := newTransaction.CreateTransaction()

	// Encode the created supplier into JSON format and send it in the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Set the status code to 201 (Created)
	json.NewEncoder(w).Encode(createdTransaction)
}

func DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the transation ID from the request parameters.
	params := mux.Vars(r)
	transactionID := params["id"]
	id, err := strconv.Atoi(transactionID)
	if err != nil {
		fmt.Println("Error while parsing:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = models.DeleteTransaction(id)
	if err != nil {
		fmt.Println("Error deleting Transaction:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Transaction  deleted successfully"))

}

func GetTransactionByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the request parameters
	params := mux.Vars(r)
	transactionID := params["id"]
	id, err := strconv.Atoi(transactionID)
	if err != nil {
		fmt.Println("Error while parsing the supplier ID:", err)
		http.Error(w, "Invalid Transaction ID", http.StatusBadRequest)
		return
	}

	// Fetch details of the Inventory from the database
	transactionDetails, err := models.GetTransactionByID(id)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	// Encode and send the details in the response
	json.NewEncoder(w).Encode(transactionDetails)

}

// UpdateTransactionHandler updates an existing Inventory item.
func UpdateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Parse incoming data
	var updatedtransaction = &models.Transaction{}
	utils.ParseBody(r, updatedtransaction)

	// Get the ID to update
	params := mux.Vars(r)
	transactionID := params["id"]
	id, err := strconv.Atoi(transactionID)
	if err != nil {
		fmt.Println("Error while parsing ID:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve transaction details from the database
	transactionDetails, err := models.GetTransactionByID(id)
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	// Update Inventory details if provided in the request
	if updatedtransaction.UserID != 0 {
		transactionDetails.UserID = updatedtransaction.UserID
	}
	if updatedtransaction.Action != "" {
		transactionDetails.Action = updatedtransaction.Action
	}
	if updatedtransaction.InventoryItemID != 0 {
		transactionDetails.InventoryItemID = updatedtransaction.InventoryItemID
	}
	if updatedtransaction.QuantityChanged != 0 {
		transactionDetails.QuantityChanged = updatedtransaction.QuantityChanged
	}
	currentTimestamp := time.Now()
	transactionDetails.UpdatedAt = currentTimestamp

	// Save the updated Transaction details
	err = transactionDetails.UpdateTransactionItem()
	if err != nil {
		fmt.Println("Error updating Transaction:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Transaction updated successfully")
}
