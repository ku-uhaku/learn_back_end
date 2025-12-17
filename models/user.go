package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"size:50;unique;not null" json:"username"`
	LastName     string         `gorm:"size:50;unique;not null" json:"lastname"`
	FirstName    string         `gorm:"size:50;unique;not null" json:"firstname"`
	Email        string         `gorm:"size:255;unique;not null" json:"email"`
	PasswordHash string         `gorm:"size:60;not null" json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
