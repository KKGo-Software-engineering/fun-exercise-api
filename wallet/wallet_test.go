package wallet

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/helper"
	"github.com/KKGo-Software-engineering/fun-exercise-api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockStorerSuccess struct{}

func (m *mockStorerSuccess) Wallets(walletType string) ([]Wallet, error) {
	mockDateString := "2024-02-03T00:00:00Z"
	layout := "2006-01-02T15:04:05Z"
	createdAt, err := time.Parse(layout, mockDateString)
	if err != nil {
		panic(err)
	}

	return []Wallet{{
		ID:         1,
		UserID:     1,
		UserName:   "John Doe",
		WalletName: "John's Wallet",
		WalletType: "Create Card",
		CreatedAt:  createdAt,
		Balance:    100,
	}}, nil
}

func (m *mockStorerSuccess) Wallet(id uint64) (Wallet, error) {
	mockDateString := "2024-02-03T00:00:00Z"
	layout := "2006-01-02T15:04:05Z"
	createdAt, err := time.Parse(layout, mockDateString)
	if err != nil {
		panic(err)
	}

	return Wallet{
		ID:         1,
		UserID:     1,
		UserName:   "John Doe",
		WalletName: "John's Wallet",
		WalletType: "Create Card",
		CreatedAt:  createdAt,
		Balance:    100,
	}, nil
}

func (m *mockStorerSuccess) CreateWallet(wallet WalletPayload) (int, error) {
	return 1, nil
}

func (m *mockStorerSuccess) UpdateWallet(id uint64, wallet WalletPayload) (int, error) {
	return 1, nil
}

func (m *mockStorerSuccess) DeleteWallet(id uint64) (int, error) {
	return 1, nil
}

type mockStorerFailure struct{}

func (m *mockStorerFailure) Wallets(walletType string) ([]Wallet, error) {
	return nil, errors.New("error forced")
}

func (m *mockStorerFailure) Wallet(id uint64) (Wallet, error) {
	return Wallet{}, errors.New("error fetching wallet")
}

func (m *mockStorerFailure) CreateWallet(wallet WalletPayload) (int, error) {
	return 0, errors.New("error creating wallet")
}

func (m *mockStorerFailure) UpdateWallet(id uint64, wallet WalletPayload) (int, error) {
	return 0, errors.New("error updating wallet")
}

func (m *mockStorerFailure) DeleteWallet(id uint64) (int, error) {
	return 0, errors.New("error deleting wallet")
}

func TestWallet(t *testing.T) {
	t.Run("Get all wallets: Internal server error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := middleware.ErrorHandler(New(&mockStorerFailure{}).WalletHandler)(c)

		// Assertions
		if assert.NoError(t, h) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

			expectedResponse := `{"status":"ERROR","result":"Oops! something went wrong."}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Get all wallets", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := New(&mockStorerSuccess{})

		// Assertions
		if assert.NoError(t, h.WalletHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			expectedResponse := `{"status":"SUCCESS","result": [{"id":1,"balance":100,"created_at":"2024-02-03T00:00:00Z","user_id":1,"user_name":"John Doe","wallet_name":"John's Wallet","wallet_type":"Create Card"}]}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Get wallet by walletId: Wallet id is required", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/122", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := middleware.ErrorHandler(New(&mockStorerFailure{}).WalletHandlerByID)(c)

		// Assertions
		if assert.NoError(t, h) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)

			expectedResponse := `{"status":"ERROR","result":{"message":"Wallet id is required","type":"ValidationError"}}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Get wallet by walletId: success", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("walletId")
		c.SetParamValues("1")

		h := New(&mockStorerSuccess{})

		// Assertions
		if assert.NoError(t, h.WalletHandlerByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			expectedResponse := `{"status":"SUCCESS","result": {"id":1,"balance":100,"created_at":"2024-02-03T00:00:00Z","user_id":1,"user_name":"John Doe","wallet_name":"John's Wallet","wallet_type":"Create Card"}}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Create wallet: Request body with missing fields", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = helper.NewValidator()

		// Prepare request body with missing fields
		whenBody := `{"user_id": null, "user_name": null, "wallet_name": null, "wallet_type": null, "balance": null}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", strings.NewReader(whenBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := middleware.ErrorHandler(New(&mockStorerFailure{}).CreateWalletHandler)(c)

		// Assertions
		if assert.NoError(t, h) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)

			expectedResponse := `{"status":"ERROR","result":{"message":"Field UserID is required, Field UserName is required, Field WalletName is required, Field WalletType is required, Field Balance is required","type":"ValidationError"}}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Create wallet: success", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = helper.NewValidator()

		// Prepare request body
		whenBody := `{"user_id": 1, "user_name": "John Doe", "wallet_name": "John's Wallet", "wallet_type": "Create Card", "balance": 100}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", strings.NewReader(whenBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := New(&mockStorerSuccess{})

		// Assertions
		if assert.NoError(t, h.CreateWalletHandler(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			expectedResponse := `{"status":"SUCCESS","result":"the wallet was just created"}`
			assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))
		}
	})

	t.Run("Update wallet: Request body with missing fields", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = helper.NewValidator()

		// Prepare request body with missing fields
		whenBody := `{"user_id": null, "user_name": null, "wallet_name": null, "wallet_type": null, "balance": null}`

		req := httptest.NewRequest(http.MethodPut, "/api/v1/wallets/1", strings.NewReader(whenBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("walletId")
		c.SetParamValues("1")

		h := middleware.ErrorHandler(New(&mockStorerFailure{}).UpdateWalletHandler)(c)

		// Assertions
		if assert.NoError(t, h) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)

			expectedResponse := `{"status":"ERROR","result":{"message":"Field UserID is required, Field UserName is required, Field WalletName is required, Field WalletType is required, Field Balance is required","type":"ValidationError"}}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Update wallet: success", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = helper.NewValidator()

		// Prepare request body
		whenBody := `{"user_id": 1, "user_name": "John Doe", "wallet_name": "John's Wallet", "wallet_type": "Create Card", "balance": 100}`

		req := httptest.NewRequest(http.MethodPut, "/api/v1/wallets/1", strings.NewReader(whenBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("walletId")
		c.SetParamValues("1")

		h := New(&mockStorerSuccess{})

		// Assertions
		if assert.NoError(t, h.UpdateWalletHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			expectedResponse := `{"status":"SUCCESS","result":"the wallet was just updated"}`
			assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))
		}
	})

	t.Run("Update wallet: wallet not found", func(t *testing.T) {
		// Setup
		e := echo.New()
		e.Validator = helper.NewValidator()

		// Prepare request body
		whenBody := `{"user_id": 1, "user_name": "John Doe", "wallet_name": "John's Wallet", "wallet_type": "Create Card", "balance": 100}`

		req := httptest.NewRequest(http.MethodPut, "/api/v1/wallets/122", strings.NewReader(whenBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("walletId")
		c.SetParamValues("122")

		h := middleware.ErrorHandler(New(&mockStorerFailure{}).UpdateWalletHandler)(c)

		// Assertions
		if assert.NoError(t, h) {
			assert.Equal(t, http.StatusNotFound, rec.Code)

			expectedResponse := `{"status":"ERROR","result":{"message":"Wallet not found","type":"NotFoundError"}}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Delete wallet: wallet not found", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/wallets/122", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("walletId")
		c.SetParamValues("122")

		h := middleware.ErrorHandler(New(&mockStorerFailure{}).DeleteWalletHandler)(c)

		// Assertions
		if assert.NoError(t, h) {
			assert.Equal(t, http.StatusNotFound, rec.Code)

			expectedResponse := `{"status":"ERROR","result":{"message":"Wallet not found","type":"NotFoundError"}}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Delete wallet: success", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/wallets/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("walletId")
		c.SetParamValues("1")

		h := New(&mockStorerSuccess{})

		// Assertions
		if assert.NoError(t, h.DeleteWalletHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			expectedResponse := `{"status":"SUCCESS","result":"the wallet was just deleted"}`
			assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))
		}
	})

	t.Run("Delete wallet: walletId is required", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/wallets", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := middleware.ErrorHandler(New(&mockStorerFailure{}).DeleteWalletHandler)(c)

		// Assertions
		if assert.NoError(t, h) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)

			expectedResponse := `{"status":"ERROR","result":{"message":"Wallet id is required","type":"ValidationError"}}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})
}
