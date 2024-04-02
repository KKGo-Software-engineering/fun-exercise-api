package middleware

import (
	"net/http"

	"github.com/KKGo-Software-engineering/fun-exercise-api/errortype"
	"github.com/KKGo-Software-engineering/fun-exercise-api/helper"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {
			c.Logger().Error(err)

			switch e := err.(type) {
			case errortype.CustomError:
				return helper.FailedHandler(c, map[string]interface{}{
					"type":    e.Type(),
					"message": e.Error(),
				}, e.Status())
			default:
				return helper.FailedHandler(c, "Oops! something went wrong.", http.StatusInternalServerError)
			}
		}

		return nil
	}
}
