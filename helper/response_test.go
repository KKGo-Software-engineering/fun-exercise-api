package helper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func TestResponseHandler(t *testing.T) {
	t.Run("Response: Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		result := map[string]interface{}{"message": "success"}

		err := SuccessHandler(c, result)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response SuccessResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "SUCCESS", response.Status)
		assert.Equal(t, result, response.Result)
	})

	t.Run("Response: Success with custom status code", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		result := map[string]interface{}{"message": "success"}

		err := SuccessHandler(c, result, http.StatusCreated)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response SuccessResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "SUCCESS", response.Status)
		assert.Equal(t, result, response.Result)
	})

	t.Run("Response: Failed", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		result := map[string]interface{}{"message": "error"}

		err := FailedHandler(c, result, http.StatusBadRequest)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response ErrorResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ERROR", response.Status)
		assert.Equal(t, result, response.Result)
	})

	t.Run("Response: Failed with default status code", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		result := map[string]interface{}{"message": "error"}

		err := FailedHandler(c, result)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response ErrorResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ERROR", response.Status)
		assert.Equal(t, result, response.Result)
	})
}
