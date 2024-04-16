package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {

	dbDriver := "mysql"
	dbName := "InventoryTracking"
	dbUser := "root"
	dbPassword := ""
	dbTcp := "@tcp(127.0.0.1:3306)/"
	d, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+dbTcp+dbName+
		"?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("gorm Db connection ", err)
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
