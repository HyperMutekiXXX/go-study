package db

import "github.com/go-redis/redis/v8"

var (
	Rdb *redis.Client
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
}
