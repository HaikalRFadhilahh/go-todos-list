package helper

import (
	"net/http"

	"github.com/HaikalRFadhilahh/go-todos-list/internal/handler"
)

type s func(w http.ResponseWriter, r *http.Request) (int, error)

func ServiceHandler(f s) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sc, err := f(w, r)
		if err != nil {
			handler.ReturnErrorResponse(w, sc, "error", err.Error())
		}
	}
}
