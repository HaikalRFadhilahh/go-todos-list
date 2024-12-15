package handler

import (
	"encoding/json"
	"net/http"
)

func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&ErrorResponse{
			StatusCode: http.StatusNotFound,
			Status:     "error",
			Message:    "Not Found",
		})
	})
}
