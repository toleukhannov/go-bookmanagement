package service

import "github.com/batyrbek/pkg/models"

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
}