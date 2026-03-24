package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nischal/rate-limiter/internal/config"
	"github.com/nischal/rate-limiter/internal/server"
)

func main() {
	// Load .env from workspace root (3 levels up from here: backend/cmd/server/)
	_ = godotenv.Load("../../../.env")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	srv.Start(cfg.Server.Port)
}
