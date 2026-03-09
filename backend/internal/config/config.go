package config

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/koanf/providers/env"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

type Config struct {
	Primary       PrimaryConfig       `koanf:"primary" validate:"required"`
	Server        ServerConfig        `koanf:"server" validate:"required"`
	Redis         RedisConfig         `koanf:"redis" validate:"required"`
	RateLimiter   RateLimiterConfig   `koanf:"rate_limiter" validate:"required"`
}

type PrimaryConfig struct {
	Env string `koanf:"env" validate:"required"`
}

type ServerConfig struct {
	Port        string `koanf:"port" validate:"required"`
	ReadTimeout int    `koanf:"read_timeout" validate:"required"`
	WriteTimeout int   `koanf:"write_timeout" validate:"required"`
	IdleTimeout int    `koanf:"idle_timeout" validate:"required"`
}

type RedisConfig struct {
	Address string `koanf:"address" validate:"required"`
	Password string `koanf:"password"`
	DB int `koanf:"db"`
}

type RateLimiterConfig struct {
	Limit  int           `koanf:"limit" validate:"required"`
	Window time.Duration `koanf:"window" validate:"required"`
}
func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp()
	k := koanf.New(".")
	
}