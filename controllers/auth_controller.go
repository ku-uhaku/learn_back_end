package controllers

import (
	"encoding/json"
	"net/http"

	"backend/helpers"
	"backend/models"
	"backend/services"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.JSONError(w, "Invalid request body", "error_sound", http.StatusBadRequest)
		return
	}

	user, token, err := services.LoginService(req.Email, req.Password)
	if err != nil {
		helpers.JSONError(w, "Invalid credentials", "error_sound", http.StatusUnauthorized)
		return
	}

	res := LoginResponse{
		User:  user,
		Token: token,
	}

	helpers.JSONSuccess(w, res, "Login successful", "success_sound")
}
