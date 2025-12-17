package middleware

import (
	"net/http"
	"runtime/debug"

	"backend/helpers"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				helpers.JSONError(w, "Internal server error", "error_sound", http.StatusInternalServerError)
				debug.PrintStack()
			}
		}()
		next.ServeHTTP(w, r)
	})
}
