package handler

import (
	"fizz-buzz.com/internal/service"
	"github.com/labstack/echo/v5"
)

// StatsHandler exposes the parameters of the most frequently made request,
// tracked in memory by the stats middleware.
func StatsHandler(ctx *echo.Context) error {
	stats := service.MostUsedRequest()
	return ctx.JSON(200, stats.Params)
}
