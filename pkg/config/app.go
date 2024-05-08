package config

import (
	"fmt"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/bookmanager?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
		fmt.Println("Failed to connect")
	}

	db = d
	fmt.Println("Successful connection to db")

}

func GetDB() *gorm.DB {
	return db
}
