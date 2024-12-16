package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"string"`
	Message    string `json:"message"`
}

type SuccessResponse struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"string"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func ReturnErrorResponse(w http.ResponseWriter, sc int, s string, m string) (int, error) {
	w.WriteHeader(sc)
	return 200, json.NewEncoder(w).Encode(&ErrorResponse{
		StatusCode: sc,
		Status:     s,
		Message:    m,
	})
}

func ReturnSuccessResponse(w http.ResponseWriter, sc int, s string, m string, d interface{}) (int, error) {
	w.WriteHeader(sc)
	return 200, json.NewEncoder(w).Encode(&SuccessResponse{
		StatusCode: sc,
		Status:     s,
		Message:    m,
		Data:       d,
	})
}
