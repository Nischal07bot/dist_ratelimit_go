package models

type RateLimitRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type RateLimitResponse struct {
	Allowed bool `json:"allowed"`
	Remaining int64 `json:"remaining"`
}