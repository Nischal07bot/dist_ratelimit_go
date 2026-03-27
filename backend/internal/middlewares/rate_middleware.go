package middlewares 

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/nischal/rate-limiter/internal/service"
)

func RateLimitMiddleware(svc *service.RateLimitService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Request().Header.Get("X-User-ID")
			if userID == "" {
				userID = "anonymous"
			}
			allowed, _, err := svc.IsAllowed(c.Request().Context(), userID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "rate limiter failed",
				})
			}
			if !allowed {
			return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "rate limiter failed",
				})
			}
			return next(c)
		}
	}
}