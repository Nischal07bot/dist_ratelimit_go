package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nischal/rate-limiter/internal/handler"
	"github.com/nischal/rate-limiter/internal/middlewares"
	"github.com/nischal/rate-limiter/internal/service"
)

func RegisterRoutes(e *echo.Echo, rateLimitHandler *handler.RateLimitHandler, svc *service.RateLimitService) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	}, middlewares.RateLimitMiddleware(svc))
	e.POST("/rate-limit/check", rateLimitHandler.Check)
	e.GET("/test", func(c echo.Context) error {
		return c.String(200, "passed")
	}, middlewares.RateLimitMiddleware(svc))
}
