package wallet

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets(walletType string) ([]Wallet, error)
	Wallet(id uint64) (Wallet, error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

func (h *Handler) isValidWalletType(walletType string) bool {
	validWalletTypes := map[string]bool{
		"Savings":       true,
		"Credit Card":   true,
		"Crypto Wallet": true,
	}

	_, ok := validWalletTypes[walletType]
	return ok
}

//	 WalletHandler
//		@Summary		Get all wallets and filter by wallet type
//		@Description	Get all wallets and filter by wallet type
//		@Tags			wallet
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	Wallet
//		@Router			/api/v1/wallets [get]
//		@Param			wallet_type query string false "Filter by wallet type" Enums(Savings, Credit Card, Crypto Wallet)
//		@Failure		400	{object}	ErrorResponse
//		@Failure		500	{object}	ErrorResponse
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

//	 WalletHandlerByID
//		@Summary		Get wallet by id
//		@Description	Get wallet by id
//		@Tags			wallet
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	Wallet
//		@Router			/api/v1/wallets/{id} [get]
//		@Param			id path string false "Wallet Id" Format(uint64)
//		@Failure		400	{object}	ErrorResponse
//		@Failure		404	{object}	ErrorResponse
//		@Failure		500	{object}	ErrorResponse
func (h *Handler) WalletHandlerByID(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("walletId"), 0, 64)
	if id == 0 {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid wallet id"})
	}

	wallet, err := h.store.Wallet(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallet)
}
