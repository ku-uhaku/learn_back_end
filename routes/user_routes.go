package routes

import (
	"net/http"

	"backend/controllers"
	"backend/middleware"

	"github.com/go-chi/chi/v5"
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
	r.Put("/{id}", controllers.UpdateUser) // PUT /schools/:id
	r.Delete("/{id}", controllers.DeleteUser)
	// Example test route
	r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have access to a protected route!"))
	})

	return r
}
