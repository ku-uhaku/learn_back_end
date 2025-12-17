package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"backend/controllers"
	"backend/middleware"
)

// UserRoutes returns a router for user-related endpoints (protected)
func UserRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Apply JWT middleware to all /users routes
	r.Use(middleware.JWTAuth)

	// Example endpoints
	r.Get("/", controllers.GetUsers)       // GET /users
	r.Get("/{id}", controllers.GetUser)    // GET /users/:id
	r.Post("/", controllers.CreateUser)    // POST /users

	// Example test route
	r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have access to a protected route!"))
	})

	return r
}
