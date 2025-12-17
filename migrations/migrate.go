package migrations

import (
	"log"

	"backend/config"
	"backend/models"
)

// ResetAndMigrate drops all tables and recreates them
func RunMigrations() {
	// Drop tables
	err := config.DB.Migrator().DropTable(
		&models.User{},
		&models.School{},
	)
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}

	log.Println("All tables dropped successfully!")

	// Recreate tables
	err = config.DB.AutoMigrate(
		&models.User{},
		&models.School{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("All migrations ran successfully!")
}
