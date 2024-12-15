package middleware

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return Content-Type = application/json
		w.Header().Set("Content-Type", "application/json")

		// Print Route URL and Method
		fmt.Println(fmt.Sprintf("Path %s - %s", r.URL.Path, r.Method))

		next.ServeHTTP(w, r)
	})
}
