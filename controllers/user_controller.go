package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/config"
	"backend/helpers"
	"backend/models"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers handles GET /api/v1/users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		helpers.JSONError(w, "Failed to fetch users", "error_sound", http.StatusInternalServerError)
		return
	}

	helpers.JSONSuccess(w, users, "Users fetched successfully", "success_sound")
}

// GetUser handles GET /api/v1/users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, "Invalid user ID", "error_sound", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		helpers.JSONError(w, "User not found", "error_sound", http.StatusNotFound)
		return
	}

	helpers.JSONSuccess(w, user, "User fetched successfully", "success_sound")
}

// CreateUser handles POST /api/v1/users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.JSONError(w, "Invalid request body", "error_sound", http.StatusBadRequest)
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), 12)
	if err != nil {
		helpers.JSONError(w, "Failed to hash password", "error_sound", http.StatusInternalServerError)
		return
	}
	req.PasswordHash = string(hash)

	if err := config.DB.Create(&req).Error; err != nil {
		helpers.JSONError(w, "Failed to create user: "+err.Error(), "error_sound", http.StatusInternalServerError)
		return
	}

	helpers.JSONSuccess(w, req, "User created successfully", "success_sound")
}

// UpdateUser handles PUT /api/v1/users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, "Invalid user ID", "error_sound", http.StatusBadRequest)
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helpers.JSONError(w, "Invalid request payload", "error_sound", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		helpers.JSONError(w, "User not found", "error_sound", http.StatusNotFound)
		return
	}

	// Optional: hash password if updated
	if input.PasswordHash != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), 12)
		if err != nil {
			helpers.JSONError(w, "Failed to hash password", "error_sound", http.StatusInternalServerError)
			return
		}
		input.PasswordHash = string(hash)
	}

	if err := config.DB.Model(&user).Updates(input).Error; err != nil {
		helpers.JSONError(w, "Failed to update user: "+err.Error(), "error_sound", http.StatusInternalServerError)
		return
	}

	helpers.JSONSuccess(w, user, "User updated successfully", "success_sound")
}

// DeleteUser handles DELETE /api/v1/users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JSONError(w, "Invalid user ID", "error_sound", http.StatusBadRequest)
		return
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		helpers.JSONError(w, "Failed to delete user", "error_sound", http.StatusInternalServerError)
		return
	}

	helpers.JSONSuccess(w, nil, "User deleted successfully", "success_sound")
}
