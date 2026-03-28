package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nischal/rate-limiter/internal/config"
	"github.com/nischal/rate-limiter/internal/repositories"
)

type RateLimitService struct {
	repo *repositories.RateLimitRepository
	cfg  config.RateLimiterConfig
}

func NewRateLimitService(repo *repositories.RateLimitRepository, cfg config.RateLimiterConfig) *RateLimitService {
	return &RateLimitService{
		repo: repo,
		cfg:  cfg,
	}
}//new rate limit service with rate limit script and client set client should be passed on to repo

func (s *RateLimitService) IsAllowed(ctx context.Context, identifier string) (bool, int64, error) {
	ctx, cancel:= context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	defer func() {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Rate Limiter check timed out for identifier:", identifier)
		}
	}()
	key := fmt.Sprintf("rate_limit:%s", identifier);
	allowed, remaining, err := s.repo.CheckLimit(
		ctx,
		key,
		int64(s.cfg.Limit),
		int64(s.cfg.Limit)/int64(s.cfg.Window.Seconds()), // refill rate
	)
	if err != nil {
		fmt.Println("RateLimiter ERROR:", err)

	if s.cfg.FailOpen {
		// allow
		return true, int64(s.cfg.Limit), nil
	}

	// fail closed
	return false, 0, nil

	}
	return allowed, remaining, nil
}