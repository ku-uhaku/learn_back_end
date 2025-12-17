package routes

import (
	"github.com/go-chi/chi/v5"
	"backend/controllers"
	"backend/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RecoverMiddleware)

	// Auth routes
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login", controllers.LoginHandler)
	})

	// Protected routes
	r.Mount("/api/v1/users", UserRoutes())

	return r
}
