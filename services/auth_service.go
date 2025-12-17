package services

import (
	"errors"
	"os"
	"time"

	"backend/config"
	"backend/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// LoginService validates user credentials and returns user + JWT token
func LoginService(email, password string) (models.User, string, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return models.User{}, "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return models.User{}, "", err
	}

	return user, tokenString, nil
}
