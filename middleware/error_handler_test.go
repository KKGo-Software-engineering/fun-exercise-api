package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KKGo-Software-engineering/fun-exercise-api/errortype"
	"github.com/labstack/echo/v4"
)

func TestErrorHandler(t *testing.T) {
	t.Run("No error", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return nil
		}
		handler := ErrorHandler(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("CustomError: ValidationError", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return errortype.ValidationError{Message: "validation error"}
		}
		handler := ErrorHandler(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Response().Writer = rec

		expectedResponse := `{"status":"ERROR","result":{"message":"validation error","type":"ValidationError"}}`
		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})

	t.Run("CustomError: NotFoundError", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return errortype.NotFoundError{Message: "not found error"}
		}
		handler := ErrorHandler(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Response().Writer = rec

		expectedResponse := `{"status":"ERROR","result":{"message":"not found error","type":"NotFoundError"}}`
		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})

	t.Run("CustomError: AuthenticationError", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return errortype.AuthenticationError{Message: "authentication error"}
		}
		handler := ErrorHandler(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Response().Writer = rec

		expectedResponse := `{"status":"ERROR","result":{"message":"authentication error","type":"AuthenticationError"}}`
		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})

	t.Run("CustomError: ForbiddenError", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return errortype.ForbiddenError{Message: "forbidden error"}
		}
		handler := ErrorHandler(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Response().Writer = rec

		expectedResponse := `{"status":"ERROR","result":{"message":"forbidden error","type":"ForbiddenError"}}`
		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusForbidden, rec.Code)
		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})

	// Case 3: Default error
	t.Run("Default error", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return errors.New("error")
		}
		handler := ErrorHandler(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Response().Writer = rec

		expectedResponse := `{"status":"ERROR","result": "Oops! something went wrong."}`
		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})
}
