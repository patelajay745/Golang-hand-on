package routes

import (
	"github.com/gorilla/mux"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/controllers"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/middleware"
)

func SetUpRoutes() *mux.Router {
	r := mux.NewRouter()

	// another way
	// Use AuthMiddleware for authentication
	// r.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET").Middleware(AuthMiddleware)

	// Protected routes with JWT token middleware
	protectedRoutes := r.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// User routes
	protectedRoutes.HandleFunc("/users", middleware.CheckUserRole(controllers.GetAllUsers, "admin")).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", middleware.CheckUserRole(controllers.GetUserByID, "admin")).Methods("GET")
	protectedRoutes.HandleFunc("/users", middleware.CheckUserRole(controllers.CreateUser, "admin")).Methods("POST")
	protectedRoutes.HandleFunc("/users/{id}", middleware.CheckUserRole(controllers.UpdateUser, "admin")).Methods("PUT")
	protectedRoutes.HandleFunc("/users/{id}", middleware.CheckUserRole(controllers.DeleteUser, "admin")).Methods("DELETE")

	// Login route
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	// Category routes
	protectedRoutes.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	protectedRoutes.HandleFunc("/categories/{id}", controllers.GetCategoryByID).Methods("GET")
	protectedRoutes.HandleFunc("/categories", middleware.CheckUserRole(controllers.CreateCategory, "admin")).Methods("POST")
	protectedRoutes.HandleFunc("/categories/{id}", middleware.CheckUserRole(controllers.UpdateCategory, "admin")).Methods("PUT")
	protectedRoutes.HandleFunc("/categories/{id}", middleware.CheckUserRole(controllers.DeleteCategory, "admin")).Methods("DELETE")

	// Supplier routes
	protectedRoutes.HandleFunc("/suppliers", middleware.CheckUserRole(controllers.GetSupplier, "admin", "manager")).Methods("GET")
	protectedRoutes.HandleFunc("/suppliers/{id}", middleware.CheckUserRole(controllers.GetSupplierByID, "admin", "manager")).Methods("GET")
	protectedRoutes.HandleFunc("/suppliers", middleware.CheckUserRole(controllers.CreateSupplier, "admin", "manager")).Methods("POST")
	protectedRoutes.HandleFunc("/suppliers/{id}", middleware.CheckUserRole(controllers.UpdateSupplier, "admin", "manager")).Methods("PUT")
	protectedRoutes.HandleFunc("/suppliers/{id}", middleware.CheckUserRole(controllers.DeleteSupplier, "admin", "manager")).Methods("DELETE")

	// InventoryItem routes
	protectedRoutes.HandleFunc("/inventory-items", controllers.GetInventoryHandler).Methods("GET")
	protectedRoutes.HandleFunc("/inventory-items/{id}", controllers.GetInventoryByIdHandler).Methods("GET")
	protectedRoutes.HandleFunc("/inventory-items", controllers.CreateInventoryHandler).Methods("POST")
	protectedRoutes.HandleFunc("/inventory-items/{id}", controllers.UpdateInventoryHandler).Methods("PUT")
	protectedRoutes.HandleFunc("/inventory-items/{id}", controllers.DeleteInventoryHandler).Methods("DELETE")

	// Transaction routes
	protectedRoutes.HandleFunc("/transactions", controllers.GetTransactionHandler).Methods("GET")
	protectedRoutes.HandleFunc("/transactions/{id}", controllers.GetTransactionByIdHandler).Methods("GET")
	protectedRoutes.HandleFunc("/transactions", controllers.CreateTransactionHandler).Methods("POST")
	protectedRoutes.HandleFunc("/transactions/{id}", controllers.UpdateTransactionHandler).Methods("PUT")
	protectedRoutes.HandleFunc("/transactions/{id}", controllers.DeleteTransactionHandler).Methods("DELETE")

	return r
}
