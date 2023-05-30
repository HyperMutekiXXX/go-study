package db

import "github.com/go-redis/redis/v8"

var (
	Rdb *redis.Client
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "120.79.13.57:6379",
		Password: "",
		DB:       0,
	})
}
