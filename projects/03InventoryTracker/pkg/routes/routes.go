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

	// Category routes
	r.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", controllers.GetCategoryByID).Methods("GET")
	r.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", controllers.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", controllers.DeleteCategory).Methods("DELETE")

	// Supplier routes
	r.HandleFunc("/suppliers", controllers.GetSupplier).Methods("GET")
	r.HandleFunc("/suppliers/{id}", controllers.GetSupplierByID).Methods("GET")
	r.HandleFunc("/suppliers", controllers.CreateSupplier).Methods("POST")
	r.HandleFunc("/suppliers/{id}", controllers.UpdateSupplier).Methods("PUT")
	r.HandleFunc("/suppliers/{id}", controllers.DeleteSupplier).Methods("DELETE")

	// InventoryItem routes
	r.HandleFunc("/inventory-items", controllers.GetInventoryHandler).Methods("GET")
	r.HandleFunc("/inventory-items/{id}", controllers.GetInventoryByIdHandler).Methods("GET")
	r.HandleFunc("/inventory-items", controllers.CreateInventoryHandler).Methods("POST")
	r.HandleFunc("/inventory-items/{id}", controllers.UpdateInventoryHandler).Methods("PUT")
	r.HandleFunc("/inventory-items/{id}", controllers.DeleteInventoryHandler).Methods("DELETE")

	// Transaction routes
	r.HandleFunc("/transactions", controllers.GetTransactionHandler).Methods("GET")
	r.HandleFunc("/transactions/{id}", controllers.GetTransactionByIdHandler).Methods("GET")
	r.HandleFunc("/transactions", controllers.CreateTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions/{id}", controllers.UpdateInventoryHandler).Methods("PUT")
	r.HandleFunc("/transactions/{id}", controllers.DeleteTransactionHandler).Methods("DELETE")

	return r
}
