// cmd/server/server.go
package main

import (
	"backend/config"
	"backend/migrations"
	"backend/routes"
	"backend/seeder"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <cmd>")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "server":
		config.Connect()
		migrations.RunMigrations()
		seeder.RunSeeders()

		r := routes.SetupRouter() // chi router
		fmt.Println("Starting HTTP server on :8080")
		log.Fatal(http.ListenAndServe(":8080", r)) // use chi router here

	case "migrate":
		config.Connect()
		migrations.RunMigrations()

	case "seed":
		config.Connect()
		seeder.RunSeeders()

	case "fresh":
		config.Connect()
		migrations.RunMigrations()
		seeder.RunSeeders()

	default:
		fmt.Println("Unknown command:", mode)
	}
}

