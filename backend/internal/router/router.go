package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nischal/rate-limiter/internal/handler"
)

func RegisterRoutes(e *echo.Echo, rateLimitHandler *handler.RateLimitHandler) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	e.POST("/rate-limiter/check", rateLimitHandler.Check)
}