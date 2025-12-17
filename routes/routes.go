package routes

import (
	"github.com/go-chi/chi/v5"
	"backend/controllers"
	"backend/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Auth routes
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login", controllers.LoginHandler)
	})

	// Protected routes
	r.Route("/api/v1/users", func(r chi.Router) {
		r.Use(middleware.JWTAuth)

		r.Get("/", controllers.GetUsers)       // GET all users
		r.Get("/{id}", controllers.GetUser)    // GET user by id
		r.Post("/", controllers.CreateUser)    // create user
	})

	return r
}
