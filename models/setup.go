package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/aquarium"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Opened to Database")

	// Migrate the schema
	// database.AutoMigrate(&Unit{}) 

	DB = database
}

