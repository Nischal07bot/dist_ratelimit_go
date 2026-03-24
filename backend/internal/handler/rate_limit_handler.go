package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/nischal/rate-limiter/internal/models"
	"github.com/nischal/rate-limiter/internal/service"
)

type RateLimitHandler struct {
	service *service.RateLimitService
}

func NewRateLimitHandler(service *service.RateLimitService) *RateLimitHandler {
	return &RateLimitHandler{
		service: service,
	}
}

func (h *RateLimitHandler) Check(c echo.Context) error {
	var req models.RateLimitRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if req.UserID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "user_id is required",
		})
	}
	allowed, remaining, err := h.service.IsAllowed(c.Request().Context(), req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}
	return c.JSON(http.StatusOK, models.RateLimitResponse{
		Allowed: allowed,
		Remaining: remaining,
	})
}
