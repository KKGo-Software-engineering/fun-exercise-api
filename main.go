package main

import (
	"github.com/KKGo-Software-engineering/fun-exercise-api/helper"
	"github.com/KKGo-Software-engineering/fun-exercise-api/middleware"
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"os"

	_ "github.com/joho/godotenv/autoload"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	dbConfig := postgres.Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	p, err := postgres.New(dbConfig)

	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.ErrorHandler)
	e.Validator = helper.NewValidator()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(p)
	handler.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
