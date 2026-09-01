package main

import (
	"fmt"
	"log"
	"os"

	"fizz-buzz.com/internal/handler"
	"fizz-buzz.com/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on existing environment variables")
	}

	e := echo.New()

	e.Validator = &CustomValidator{
		validator: validator.New(),
	}
	e.Use(middleware.StatsMiddleware)

	e.GET("/fuzzbuzz", handler.FuzzBuzzHandler)
	e.GET("/stats", handler.StatsHandler)
	e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
