package seeder

import "log"

// RunSeeders runs all seeders
func RunSeeders() {
	log.Println("Starting database seeding...")

	// Run individual seeders (call directly)
	SeedUsers()
	SeedSchools()

	log.Println("All seeders completed successfully!")
}
