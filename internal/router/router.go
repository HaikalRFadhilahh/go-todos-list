package router

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/go-todos-list/internal/controller"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/db"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/di"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/handler"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/middleware"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/pkg/helper"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// Database Connection
	db := db.InitDB()

	// Dependency Injection Struct
	di := di.InitDI(db)

	// Controller
	uc := controller.UserController{DI: di}

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
	r.HandleFunc("/v1/users/getall", helper.ServiceHandler(uc.GetUser)).Methods("GET")

	// Handler Error Not Found
	r.NotFoundHandler = handler.NotFound()

	// Return Router
	return r
}
