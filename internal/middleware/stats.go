package middleware

import (
	"fizz-buzz.com/internal/service"
	"github.com/labstack/echo/v5"
)

// StatsMiddleware records every incoming request's query parameters so the
// parameters of the most frequently used request can be tracked, purely in
// memory (no persistence).
func StatsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		service.RecordRequest(c.QueryParams())
		return next(c)
	}
}
