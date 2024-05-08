package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/batyrbek/pkg/config"
	"github.com/batyrbek/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string
	Password string
}

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}


func SignUp(c *gin.Context) {
    var request models.User
    if c.BindJSON(&request) != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to read body",
        })
        return
    }

    // Hashing
    hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    request.Password = string(hashPassword)
    request.CreatedAt = time.Now()

    // New User GORM
    if err := config.GetDB().Create(&request).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"user": request})
}

func Login(c *gin.Context) {
    var loginRequest LoginRequest
    if c.BindJSON(&loginRequest) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
        return
    }

    // user findByID
    var user models.User
    if err := config.GetDB().Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // pass checking
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // If authentication is successful, we create a token and send it to the user
    token, err := createToken(user.Username)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
        return
    }

    // Setting cookie token
    c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
