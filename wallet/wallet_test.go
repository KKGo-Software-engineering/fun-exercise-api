package wallet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

type StubWallet struct {
	wallets         []Wallet
	walletsByUserId []Wallet
	err             error
}

func (s StubWallet) Wallets() ([]Wallet, error) {
	return s.wallets, s.err
}

func (s StubWallet) WalletsByUserId(id string) ([]Wallet, error) {
	return s.walletsByUserId, s.err
}

func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubError := StubWallet{err: echo.ErrInternalServerError}
		s := New(stubError)

		s.WalletHandler(c)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		stubUser := StubWallet{
			walletsByUserId: []Wallet{
				{UserName: "John Doe", WalletType: "Savings"},
				{UserName: "John Doe", WalletType: "Credit Card"},
				{UserName: "John Doe", WalletType: "Crypto Wallet"},
			},
		}
		p := New(stubUser)

		p.WalletHandlerByUserId(c)

		wantUserName := "John Doe"
		want := []Wallet{
			{UserName: wantUserName, WalletType: "Savings"},
			{UserName: wantUserName, WalletType: "Credit Card"},
			{UserName: wantUserName, WalletType: "Crypto Wallet"},
		}
		gotJson := rec.Body.Bytes()
		var got []Wallet
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
