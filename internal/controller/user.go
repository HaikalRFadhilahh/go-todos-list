package controller

import (
	"net/http"

	"github.com/HaikalRFadhilahh/go-todos-list/internal/di"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/handler"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/service/user"
)

type UserController struct {
	DI *di.DependencyInjection
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) (int, error) {
	res, err := user.GetAll(c.DI.DB)
	if err != nil {
		return 500, err
	}
	return handler.ReturnSuccessResponse(w, http.StatusOK, "success", "Users Data", res)
}
