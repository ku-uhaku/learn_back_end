package seeder

import (
	"backend/config"
	"backend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SeedUsers inserts initial user data
func SeedUsers() {
	users := []models.User{
		{
			Username:     "admin",
			FirstName:    "Administrator",
			LastName:     "Administrator",
			Email:        "admin@gmail.com",
			PasswordHash: hashPassword("password"),
		},
		{
			Username:     "john_doe",
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "john@gmail.com",
			PasswordHash: hashPassword("password"),
		},
	}

	for _, u := range users {
		// Check if user already exists
		var existing models.User
		result := config.DB.Where("email = ?", u.Email).First(&existing)
		if result.Error != nil && result.Error.Error() != "record not found" {
			log.Fatal("Failed to query users:", result.Error)
		}

		if existing.ID == 0 {
			// User does not exist, create
			if err := config.DB.Create(&u).Error; err != nil {
				log.Fatal("Failed to seed user:", err)
			}
		} else {
			// Optionally update user info
			config.DB.Model(&existing).Updates(models.User{
				FirstName:    u.FirstName,
				LastName:     u.LastName,
				PasswordHash: u.PasswordHash,
			})
		}
	}

	log.Println("User seeding completed!")
}

// hashPassword hashes passwords using bcrypt
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}
	return string(hash)
}
