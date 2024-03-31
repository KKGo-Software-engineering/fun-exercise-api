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
	CreateWallet(wallet WalletPayload) (int, error)
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

type WalletPayload struct {
	UserID     int     `json:"user_id" example:"1" validate:"required"`
	UserName   string  `json:"user_name" example:"John Doe" validate:"required"`
	WalletName string  `json:"wallet_name" example:"John's Wallet" validate:"required"`
	WalletType string  `json:"wallet_type" example:"Create Card" validate:"required"`
	Balance    float64 `json:"balance" example:"100.00" validate:"required"`
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
//		@Failure		400	{object}	Err
//		@Failure		500	{object}	Err
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
//		@Failure		400	{object}	Err
//		@Failure		404	{object}	Err
//		@Failure		500	{object}	Err
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

//	 CreateWalletHandler
//		@Summary		Create wallet
//		@Description	Create wallet
//		@Tags			wallet
//		@Accept			json
//		@Produce		json
//		@Success		201	{object}	string
//		@Router			/api/v1/wallets [post]
//		@Param			wallet body WalletPayload true "Wallet Payload" Format(WalletPayload)
//		@Failure		400	{object}	Err
//		@Failure		500	{object}	Err
func (h *Handler) CreateWalletHandler(c echo.Context) error {
	var wallet WalletPayload
	var err error
	if err = c.Bind(&wallet); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(&wallet); err != nil {
		return err
	}

	if _, err := h.store.CreateWallet(wallet); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	createdResponse := "the wallet was just created"
	return c.JSON(http.StatusCreated, createdResponse)
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/wallets", h.WalletHandler)
	e.GET("/api/v1/wallets/:walletId", h.WalletHandlerByID)
	e.POST("/api/v1/wallets", h.CreateWalletHandler)
}
