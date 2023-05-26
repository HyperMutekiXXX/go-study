package db

import "github.com/go-redis/redis/v8"

var (
	Rdb *redis.Client
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.56.105:6379",
		Password: "",
		DB:       0,
	})
}
