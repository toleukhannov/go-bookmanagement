package service

import (
    "errors"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// GenerateToken генерирует JWT токен на основе имени пользователя
func GenerateToken(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// VerifyToken проверяет валидность JWT токена
func VerifyToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}




// package service

// import (
// 	"crypto/sha1"
// 	"errors"
// 	"fmt"
// 	"time"

// 	"github.com/batyrbek/pkg/models"
// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/zhashkevych/todo-app/pkg/repository"
// )

// const (
// 	salt       = "hjqrhjqw124617ajfhajs"
// 	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
// 	tokenTTL   = 12 * time.Hour
// )

// type tokenClaims struct {
// 	jwt.StandardClaims
// 	UserId int `json:"user_id"`
// }

// type AuthService struct {
// 	repo repository.Authorization
// }

// func NewAuthService(repo repository.Authorization) *AuthService {
// 	return &AuthService{repo: repo}
// }

// func (s *AuthService) CreateUser(user models.User) (int, error) {
// 	user.Password = generatePasswordHash(user.Password)
// 	return s.repo.CreateUser(user)
// }

// func (s *AuthService) GenerateToken(username, password string) (string, error) {
// 	user, err := s.repo.GetUser(username, generatePasswordHash(password))
// 	if err != nil {
// 		return "", err
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 		user.Id,
// 	})

// 	return token.SignedString([]byte(signingKey))
// }

// func (s *AuthService) ParseToken(accessToken string) (int, error) {
// 	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("invalid signing method")
// 		}

// 		return []byte(signingKey), nil
// 	})
// 	if err != nil {
// 		return 0, err
// 	}

// 	claims, ok := token.Claims.(*tokenClaims)
// 	if !ok {
// 		return 0, errors.New("token claims are not of type *tokenClaims")
// 	}

// 	return claims.UserId, nil
// }

// func generatePasswordHash(password string) string {
// 	hash := sha1.New()
// 	hash.Write([]byte(password))

// 	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }
