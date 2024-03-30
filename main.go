package main

import (
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"os"

	"github.com/joho/godotenv"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("please create .env file with the following file .env.example")
	}

	dbConfig := postgres.Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	p, err := postgres.New(dbConfig)

	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(p)
	e.GET("/api/v1/wallets", handler.WalletHandler)
	e.GET("/api/v1/wallets/:walletId", handler.WalletHandlerByID)
	e.Logger.Fatal(e.Start(":1323"))
}
