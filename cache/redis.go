package cache

import (
	"context"
	"defaultProject/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx context.Context

type RedisCache struct{}

func (r *RedisCache) Init() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.GetString("Redis.Host"),
		Password: config.Cfg.GetString("Redis.Passwd"), // no password set
		DB:       0,                                    // use default DB
		PoolSize: config.Cfg.GetInt("Redis.PoolSize"),  // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	client = rdb
	return nil
}

func GetCache() *redis.Client {
	return client
}

func Set(key, val string, expire ...time.Duration) error {
	var expiration time.Duration = 0
	if len(expire) > 0 {
		expiration = expire[0]
	}
	return client.Set(ctx, key, val, expiration).Err()
}

func Get(key string) string {
	r, err := client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return r
}
