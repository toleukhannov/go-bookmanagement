package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// package models

// import (
// 	"github.com/batyrbek/pkg/config"
// 	"github.com/jinzhu/gorm"
// )

// type User struct {
// 	gorm.Model
// 	Id       int    `gorm:""json:"-"`
// 	Name     string `gorm:""json:"name" binding: "required"`
// 	Email    string `gorm:""json:"email" binding: "required"`
// 	Password string `gorm:""json:"password" binding: "required"`
// }

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	db.AutoMigrate(&Book{})
// }

// func (u *User) CreateUser() *User {
// 	db.NewRecord(u)
// 	db.Create(&u)
// 	return u
// }
