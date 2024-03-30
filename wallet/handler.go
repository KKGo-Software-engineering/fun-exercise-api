package wallet

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets(walletType string) ([]Wallet, error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

func (h *Handler) isValidWalletType(walletType string) bool {
	validWalletTypes := map[string]bool{
			"Savings":     true,
			"Credit Card": true,
			"Crypto Wallet": true,
	}

	_, ok := validWalletTypes[walletType]
	return ok
}

// WalletHandler
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletHandler(c echo.Context) error {
	walletType := c.QueryParam("wallet_type")

	if walletType == "" {
		walletType = ""
	}

	if walletType != "" && !h.isValidWalletType(walletType) {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid wallet type"})
	}

	wallets, err := h.store.Wallets(walletType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}
