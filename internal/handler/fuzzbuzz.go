package handler

import (
	"fizz-buzz.com/internal/model"
	"fizz-buzz.com/internal/service"
	"github.com/labstack/echo/v5"
)

func FuzzBuzzHandler(ctx *echo.Context) error {
	var req model.FuzzBuzzRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, map[string]string{"error": "Invalid request"})
	}

	result := service.FuzzBuzz(req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)
	return ctx.JSON(200, result)
}
