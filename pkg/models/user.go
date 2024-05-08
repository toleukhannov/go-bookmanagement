package models

import (
	"github.com/batyrbek/pkg/config"
	"github.com/jinzhu/gorm"
)


type User struct {
	gorm.Model
	ID       int    `gorm:"" json:"id"`
	Username string `gorm:"" json:"username"`
	Password string `gorm:"" json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) newUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func DeleteUser(ID int64) User{
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}