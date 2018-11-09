package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type DefaultResponse struct {
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Status  int         `json:"-"`
}

func ShowErrorResponse(e echo.Context, err error) error {
	response := DefaultResponse{
		Message: err.Error(),
		Status:  http.StatusInternalServerError,
	}
	return ShowResponse(e, response)
}

func ShowSuccessResponse(e echo.Context, r interface{}) error {
	response := DefaultResponse{
		Message: "success",
		Result:  r,
		Status:  http.StatusOK,
	}
	return ShowResponse(e, response)
}

func ShowResponse(e echo.Context, r DefaultResponse) error {
	return e.JSON(r.Status, r)
}
