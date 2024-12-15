package handler

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"string"`
	Message    string `json:"message"`
}

type SuccessResponse struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"string"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}
