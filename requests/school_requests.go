package requests

import "time"

type SchoolCreateRequest struct {
	Code           string     `json:"code" validate:"required"`
	Name           string     `json:"name" validate:"required"`
	MassarID       string     `json:"massar_id"`
	Email          string     `json:"email" validate:"omitempty,email"`
	Phone          string     `json:"phone"`
	City           string     `json:"city"`
	StartDate      *time.Time `json:"start_date"`
	Status         *bool      `json:"status"`
	RegistrationNum string    `json:"registration_number"`
}

type SchoolUpdateRequest struct {
	Name     *string `json:"name" validate:"omitempty"`
	Email    *string `json:"email" validate:"omitempty,email"`
	City     *string `json:"city"`
	Status   *bool   `json:"status"`
}
