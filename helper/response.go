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

func SuccessHandler(c echo.Context, result interface{}, statusCode ...int) error {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	response := SuccessResponse{
		Status: "SUCCESS",
		Result: result,
	}
	return c.JSON(code, response)
}

func FailedHandler(c echo.Context, result interface{}, statusCode ...int) error {
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	response := ErrorResponse{
		Status: "ERROR",
		Result: result,
	}
	return c.JSON(code, response)
}
