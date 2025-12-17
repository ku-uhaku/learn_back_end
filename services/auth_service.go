package services

import (
	"backend/config"
	"backend/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("YOUR_SECRET_KEY") // Replace with env variable

// LoginService checks user credentials and returns JWT token
func LoginService(email, password string) (string, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", err
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24h expiration
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Failed to sign token:", err)
		return "", err
	}

	return tokenString, nil
}
