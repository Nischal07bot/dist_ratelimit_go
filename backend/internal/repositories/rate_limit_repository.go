package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimitRepository struct {
	client redis.UniversalClient
	script *redis.Script
}

func NewRateLimitRepository(client redis.UniversalClient) *RateLimitRepository {

	luaScript := redis.NewScript(`
local key = KEYS[1]

local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

local data = redis.call("HMGET", key, "tokens", "last_refill")

local tokens = tonumber(data[1])
local last_refill = tonumber(data[2])

-- Proper initialization check
if data[1] == false or data[2] == false then
    tokens = capacity
    last_refill = now
end

local delta = math.max(0, now - last_refill)
local refill = math.floor((delta / 1000) * refill_rate)

tokens = math.min(capacity, tokens + refill)

if tokens < 1 then
    return {0, tokens}
end

tokens = tokens - 1

redis.call("HMSET", key, "tokens", tokens, "last_refill", now)
redis.call("EXPIRE", key, 3600)

return {1, tokens}
`)

	return &RateLimitRepository{
		client: client,
		script: luaScript,
	}
}

func (r *RateLimitRepository) CheckLimit(
	ctx context.Context,
	key string,
	capacity int64,
	refillRate int64,
) (bool, int64, error) {

	now := time.Now().UnixMilli()

	result, err := r.script.Run(
		ctx,
		r.client,
		[]string{key},
		capacity,
		refillRate,
		now,
	).Result()
	fmt.Println("Running script for key:", key)
	fmt.Println("Result:", result)
	if err != nil {
		return false, 0, err
	}

	values := result.([]interface{})

	allowed := values[0].(int64)
	remaining := values[1].(int64)

	return allowed == 1, remaining, nil
}
