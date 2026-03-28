package server

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/nischal/rate-limiter/internal/config"
	"github.com/nischal/rate-limiter/internal/handler"
	"github.com/nischal/rate-limiter/internal/lib/redis"
	"github.com/nischal/rate-limiter/internal/repositories"
	"github.com/nischal/rate-limiter/internal/router"
	"github.com/nischal/rate-limiter/internal/service"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(cfg *config.Config) (*Server, error) {

	e := echo.New()

	// 1. Redis client
	redisClient, err := redis.NewClient(cfg.Redis)
	if err != nil {
		return nil, err
	}

	// 2. Repository
	repo := repositories.NewRateLimitRepository(redisClient.GetClient())

	// 3. Service
	svc := service.NewRateLimitService(repo, cfg.RateLimiter)

	// 4. Handler
	handler := handler.NewRateLimitHandler(svc)
	//e.Use(middlewares.RateLimitMiddleware(svc))
	// 5. Routes
	router.RegisterRoutes(e, handler, svc)

	return &Server{
		echo: e,
	}, nil
}

func (s *Server) Start(port string) {

	log.Println("Server running on port", port)
	if err := s.echo.Start(":" + port); err != nil {
		log.Println("Server error:", err)
	}
}
