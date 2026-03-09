package main

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/redis/go-redis/v9"
)

func main() {
    _ = godotenv.Load()

    redisURL := os.Getenv("REDIS_URL")
    if redisURL == "" {
        log.Fatal("REDIS_URL is not set")
    }

    opts, err := redis.ParseURL(redisURL)
    if err != nil {
        log.Fatalf("invalid REDIS_URL: %v", err)
    }

    rdb := redis.NewClient(opts)
    if err := rdb.Ping(context.Background()).Err(); err != nil {
        log.Fatalf("redis connection failed: %v", err)
    }

    log.Println("Redis connected successfully")
    log.Println("Starting Rate Limiter Service")
}