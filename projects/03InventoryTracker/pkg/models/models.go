package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/patelajay745/projects/03InventoryTracker/pkg/config"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Supplier{})
	db.AutoMigrate(&InventoryItem{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&User{})

}

type CustomModel struct {
	ID uint `gorm:"primaryKey"`
}

type User struct {
	CustomModel
	Username string `json:"username" validate:"required, min=2,max=100"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" validate:"required, eq=admin|eq=manager|eq=cashier|eq=cook|eq=Admin|eq=Manager|eq=Cashier|eq=Cook"`
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

// GetUserByID retrieves a user by ID.
func GetUserByID(id int) (*User, *gorm.DB) {
	var user User
	db := db.Where("ID=?", id).Find(&user)

	return &user, db
}

// GetAllUsers
func GetAllUsers() []User {
	var users []User
	db.Find(&users)

	return users
}

// UpdateUser updates an existing user.
func (u *User) UpdateUser() error {
	if err := db.Save(&u).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes an existing user.
func DeleteUser(id int) error {
	if err := db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// LoginUser authenticates a user based on email and password.
func LoginUser(email, password string) (*User, error) {
	var user User
	if err := db.Where("username = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Category represents a category of inventory items.
type Category struct {
	CustomModel
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GetUserByID retrieves a user by ID.
func GetCategoryByID(id int) *Category {
	var category Category
	db.Where("ID=?", id).Find(&category)

	return &category
}

func GetAllCategories() []Category {

	var categories []Category
	db.Find(&categories)

	return categories
}

func (cat *Category) CreateCategory() *Category {
	db.NewRecord(cat)
	db.Create(&cat)
	return cat
}

// Delete Category
func DeleteCategory(id int) error {
	if err := db.Delete(&Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Update Category
func (cat *Category) UpdateCategory() error {
	if err := db.Save(&cat).Error; err != nil {
		return err
	}
	return nil
}

// Supplier represents a supplier of inventory items.
type Supplier struct {
	CustomModel
	Name          string    `json:"name"`
	ContactPerson string    `json:"contact_person"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Get All Suppliers
func GetAllSuppliers() []Supplier {
	var suppliers []Supplier
	db.Find(&suppliers)
	return suppliers
}

// Create Supplier
func (supplier *Supplier) CreateSupplier() *Supplier {
	db.NewRecord(supplier)
	db.Create(&supplier)
	return supplier
}

// Delete Supplier
func DeleteSupplier(id int) error {
	if err := db.Delete(&Supplier{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetSupplierByID retrieves details of a supplier by ID.
func GetSupplierByID(id int) (*Supplier, error) {
	var supplier Supplier
	err := db.Where("ID=?", id).First(&supplier).Error
	if err != nil {
		// Check if the error is due to record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("supplier with ID %d not found", id)
		}
		// Return other database-related errors
		return nil, fmt.Errorf("error fetching supplier details: %v", err)
	}
	return &supplier, nil
}

// update Supplier
func (supplier *Supplier) UpdateSupplier() error {
	if err := db.Save(&supplier).Error; err != nil {
		return err
	}
	return nil
}

// InventoryItem represents an individual inventory item.
type InventoryItem struct {
	CustomModel
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  int       `json:"category_id"`
	SupplierID  int       `json:"supplier_id"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Get All InventoryItem
func GetAllInventory() []InventoryItem {
	var inventories []InventoryItem
	db.Find(&inventories)
	return inventories
}

// Get INventoryByID
func GetInventoryByID(id int) (*InventoryItem, error) {
	var inventory InventoryItem
	err := db.Where("ID=?", id).First(&inventory).Error
	if err != nil {
		// Check if the error is due to record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("supplier with ID %d not found", id)
		}
		// Return other database-related errors
		return nil, fmt.Errorf("error fetching supplier details: %v", err)
	}
	fmt.Println(&inventory)
	return &inventory, nil

}

// Create Inventory Item
func (inventory *InventoryItem) CreateInventory() *InventoryItem {
	db.NewRecord(inventory)
	db.Create(&inventory)
	return inventory
}

// Delete Inventory Record
func DeleteInventoryItem(id int) error {
	if err := db.Delete(&InventoryItem{}, id).Error; err != nil {
		return err
	}
	return nil
}

// update inventory Record
func (inventory *InventoryItem) UpdateInventroryItem() error {
	if err := db.Save(&inventory).Error; err != nil {
		return err
	}
	return nil
}

// Transaction represents a transaction performed on an inventory item.
type Transaction struct {
	CustomModel
	UserID          int       `json:"user_id"`
	Action          string    `json:"action"`
	InventoryItemID int       `json:"inventory_item_id"`
	QuantityChanged int       `json:"quantity_changed"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Get All Transaction
func GetAllTransaction() []Transaction {
	var transaction []Transaction
	db.Find(&transaction)
	return transaction
}

// Get TransactionByID
func GetTransactionByID(id int) (*Transaction, error) {
	var trans Transaction
	err := db.Where("ID=?", id).First(&trans).Error
	if err != nil {
		// Check if the error is due to record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("supplier with ID %d not found", id)
		}
		// Return other database-related errors
		return nil, fmt.Errorf("error fetching supplier details: %v", err)
	}
	return &trans, nil

}

// Create Transaction
func (trans *Transaction) CreateTransaction() *Transaction {
	db.NewRecord(trans)
	db.Create(&trans)
	return trans
}

// Delete Transaction Record
func DeleteTransaction(id int) error {
	if err := db.Delete(&Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}

// update inventory Record
func (transaction *Transaction) UpdateTransactionItem() error {
	if err := db.Save(&transaction).Error; err != nil {
		return err
	}
	return nil
}

// for login
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
