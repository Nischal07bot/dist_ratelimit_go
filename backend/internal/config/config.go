package config

import (
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	Primary     PrimaryConfig     `koanf:"primary" validate:"required"`
	Server      ServerConfig      `koanf:"server" validate:"required"`
	Redis       RedisConfig       `koanf:"redis" validate:"required"`
	RateLimiter RateLimiterConfig `koanf:"rate_limiter" validate:"required"`
	
}

type PrimaryConfig struct {
	Env string `koanf:"env" validate:"required"`
}

type ServerConfig struct {
	Port         string `koanf:"port" validate:"required"`
	ReadTimeout  int    `koanf:"read_timeout" validate:"required"`
	WriteTimeout int    `koanf:"write_timeout" validate:"required"`
	IdleTimeout  int    `koanf:"idle_timeout" validate:"required"`
}

type RedisConfig struct {
	Address  string `koanf:"address" validate:"required"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
	UseCluster bool   `koanf:"use_cluster"`
}

type RateLimiterConfig struct {
	Limit  int           `koanf:"limit" validate:"required"`
	Window time.Duration `koanf:"window" validate:"required"`
	FailOpen    bool              `koanf:"fail_open"`
}

func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	k := koanf.New(".")
	if err := k.Load(env.Provider("RATE_LIMITER_", ".", func(s string) string {
		key := strings.TrimPrefix(s, "RATE_LIMITER_")
		key = strings.ToLower(key)
		// Use double underscores for nesting, e.g. RATE_LIMITER_SERVER__PORT -> server.port
		key = strings.ReplaceAll(key, "__", ".")
		return key
	}), nil); err != nil {
		logger.Fatal().Err(err).Msg("failed to load configuration")
	}
	cfg := &Config{}

	if err := k.Unmarshal("", cfg); err != nil {
		logger.Fatal().Err(err).Msg("failed to unmarshal configuration")
	}
	validate := validator.New()

	if err := validate.Struct(cfg); err != nil {
		logger.Fatal().Err(err).Msg("config validation failed")
	}

	return cfg, nil

}
