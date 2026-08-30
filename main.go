package main

import (
	"fmt"
	"log"
	"os"

	"fizz-buzz.com/internal/handler"
	"fizz-buzz.com/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on existing environment variables")
	}

	e := echo.New()

	e.Use(middleware.StatsMiddleware)

	e.GET("/fuzzbuzz", handler.FuzzBuzzHandler)
	e.GET("/stats", handler.StatsHandler)
	e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
