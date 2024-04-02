package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
} // @name ErrorResponse

type SuccessResponse struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
} // @name SuccessResponse

func SuccessHandler(c echo.Context, status string, result interface{}, statusCode ...int) error {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	response := SuccessResponse{
		Status: status,
		Result: result,
	}
	return c.JSON(code, response)
}

func FailedHandler(c echo.Context, status string, result interface{}, statusCode int) error {
	response := ErrorResponse{
		Status: status,
		Result: result,
	}
	return c.JSON(statusCode, response)
}
