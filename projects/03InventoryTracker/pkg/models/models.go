package models

import (
	"time"

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

}

type CustomModel struct {
	ID uint `gorm:"primaryKey"`
}

// User represents a user of the system.
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
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
	if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
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

// Transaction represents a transaction performed on an inventory item.
type Transaction struct {
	CustomModel
	UserID          int       `json:"user_id"`
	Action          string    `json:"action"`
	InventoryItemID int       `json:"inventory_item_id"`
	QuantityChanged int       `json:"quantity_changed"`
	Timestamp       time.Time `json:"timestamp"`
}
