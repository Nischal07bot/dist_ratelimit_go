package redis

import (
	"context"
	"time"

	"github.com/nischal/rate-limiter/internal/config"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	rdb redis.UniversalClient //encapsulate the redis client
}

func NewClient(cfg config.RedisConfig) (*Client, error) {
if(cfg.UseCluster) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        []string{cfg.Address},
		Password:     cfg.Password,
		
		PoolSize:     20,
		MinIdleConns: 5,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err:= rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Client{
		rdb: rdb,
	}, nil

}
rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Address,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     20,
		MinIdleConns: 5,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err:= rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Client{
		rdb: rdb,
	}, nil

}

func (c *Client) GetClient() redis.UniversalClient {
	return c.rdb
}
func (c *Client) Close() error {
	return c.rdb.Close()
}