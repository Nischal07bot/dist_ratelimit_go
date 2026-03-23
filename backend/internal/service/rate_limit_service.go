package service

import (
	"context"
	"fmt"
	"github.com/nischal/rate-limiter/internal/config"
	"github.com/nischal/rate-limiter/internal/repositories"
)

type RateLimitService struct {
	repo *repositories.RateLimitRepository
	cfg  *config.RateLimiterConfig
}

func NewRateLimitService(repo *repositories.RateLimitRepository, cfg *config.RateLimiterConfig) *RateLimitService {
	return &RateLimitService{
		repo: repo,
		cfg:  cfg,
	}
}//new rate limit service with rate limit script and client set client should be passed on to repo

func (s *RateLimitService) IsAllowed(ctx context.Context, identifier string) (bool, int64, error) {
	key := fmt.Sprintf("rate_limit:%s", identifier);
	allowed, remaining, err := s.repo.CheckLimit(
		ctx,
		key,
		int64(s.cfg.Limit),
		int64(s.cfg.Limit)/int64(s.cfg.Window.Seconds()), // refill rate
	)
	if err != nil {
		return false, 0, err
	}
	return allowed, remaining, nil
}