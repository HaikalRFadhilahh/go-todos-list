package router

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/go-todos-list/internal/handler"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// Init Router Routing
	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.LoggingMiddleware)

	// Mux Route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&handler.ErrorResponse{
			StatusCode: http.StatusOK,
			Status:     "success",
			Message:    "Todos List APP Go With Mux Library",
		})
	}).Methods("GET")

	// Handler Error Not Found
	r.NotFoundHandler = handler.NotFound()

	// Return Router
	return r
}
