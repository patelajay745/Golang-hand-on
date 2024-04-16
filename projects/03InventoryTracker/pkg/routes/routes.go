package routes

import (
	"github.com/gorilla/mux"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/controllers"
)

func SetUpRoutes() *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// Login route
	//r.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	// // Category routes
	// r.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	// r.HandleFunc("/categories/{id}", controllers.GetCategory).Methods("GET")
	// r.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")
	// r.HandleFunc("/categories/{id}", controllers.UpdateCategory).Methods("PUT")
	// r.HandleFunc("/categories/{id}", controllers.DeleteCategory).Methods("DELETE")

	// // Supplier routes
	// r.HandleFunc("/suppliers", controllers.GetSuppliers).Methods("GET")
	// r.HandleFunc("/suppliers/{id}", controllers.GetSupplier).Methods("GET")
	// r.HandleFunc("/suppliers", controllers.CreateSupplier).Methods("POST")
	// r.HandleFunc("/suppliers/{id}", controllers.UpdateSupplier).Methods("PUT")
	// r.HandleFunc("/suppliers/{id}", controllers.DeleteSupplier).Methods("DELETE")

	// // InventoryItem routes
	// r.HandleFunc("/inventory-items", controllers.GetInventoryItems).Methods("GET")
	// r.HandleFunc("/inventory-items/{id}", controllers.GetInventoryItem).Methods("GET")
	// r.HandleFunc("/inventory-items", controllers.CreateInventoryItem).Methods("POST")
	// r.HandleFunc("/inventory-items/{id}", controllers.UpdateInventoryItem).Methods("PUT")
	// r.HandleFunc("/inventory-items/{id}", controllers.DeleteInventoryItem).Methods("DELETE")

	// // Transaction routes
	// r.HandleFunc("/transactions", controllers.GetTransactions).Methods("GET")
	// r.HandleFunc("/transactions/{id}", controllers.GetTransaction).Methods("GET")
	// r.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
	// r.HandleFunc("/transactions/{id}", controllers.UpdateTransaction).Methods("PUT")
	// r.HandleFunc("/transactions/{id}", controllers.DeleteTransaction).Methods("DELETE")

	return r
}
