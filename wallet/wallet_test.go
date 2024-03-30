package wallet

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

type mockStorerFailure struct{}

func (m *mockStorerFailure) Wallets(walletType string) ([]Wallet, error) {
	return nil, errors.New("error fetching wallets")
}

func (m *mockStorerFailure) Wallet(id uint64) (Wallet, error) {
	return Wallet{}, errors.New("error fetching wallet")
}

func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := New(&mockStorerFailure{})

		// Assertions
		if assert.NoError(t, h.WalletHandler(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

			expectedResponse := `{"message":"error fetching wallets"}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := New(&mockStorerSuccess{})

		// Assertions
		if assert.NoError(t, h.WalletHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			expectedResponse := `[{"id":1,"balance":100,"created_at":"2024-02-03T00:00:00Z","user_id":1,"user_name":"John Doe","wallet_name":"John's Wallet","wallet_type":"Create Card"}]`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})

	t.Run("Get wallet by walletId: not found", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallet/122", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := New(&mockStorerFailure{})

		// Assertions
		if assert.NoError(t, h.WalletHandlerByID(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)

			expectedResponse := `{"message":"Invalid wallet id"}`
			assert.JSONEq(t, expectedResponse, rec.Body.String())
		}
	})
}
