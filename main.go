package main

import (
	"fmt"
	"os"

	"fizz-buzz.com/internal/handler"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	e.GET("/fuzzbuzz", handler.FuzzBuzzHandler)
	e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
