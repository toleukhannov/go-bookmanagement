package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"fmt"

	"github.com/batyrbek/pkg/config"
	"github.com/batyrbek/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)
type LoginRequest struct {
	Username string
	Password string
}
//Token interactions
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
//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	var request models.User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	request.Password = string(hashPassword)
	request.CreatedAt = time.Now()

	// Create new user using GORM
	if err := config.GetDB().Create(&request).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Respond with created user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(request)
}
//Login Handler
func Login(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
 Username string
 Password string
}
    var loginRequest LoginRequest
    err := json.NewDecoder(r.Body).Decode(&loginRequest)
    if err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }

    // Find user by username
    var user models.User
    if err := config.GetDB().Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // If authentication is successful, create a token and send it to the user
    token, err := createToken(user.Username)
    if err != nil {
        http.Error(w, "Failed to create token", http.StatusInternalServerError)
        return
    }

    // Set the token in a cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "Authorization",
        Value:    token,
        MaxAge:   3600 * 24 * 30,
        Path:     "/",
        SameSite: http.SameSiteStrictMode,
        Secure:   true,
        HttpOnly: true,
    })

    // Respond with successful login message
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func Logout(w http.ResponseWriter, r *http.Request) {
    // Remove the Authorization cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "Authorization",
        Value:    "",
        MaxAge:   -1,
        Path:     "/",
        SameSite: http.SameSiteStrictMode,
        Secure:   true,
        HttpOnly: true,
    })

    // Respond with a message indicating successful logout
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Logout successful"}`))
}
