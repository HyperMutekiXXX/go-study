package dao

import (
	"context"
	"github.com/go-study/db"
	"time"
)

var (
	ctx context.Context
)

func init() {
	ctx = context.Background()
}

func GetByKey(key string) interface{} {
	result, err := db.Rdb.Do(ctx, "get", key).Result()
	if err != nil {
		panic(err)
	}
	return result
}

func SetKeyValue(key string, value interface{}, ttl int64) {
	err := db.Rdb.Set(ctx, key, value, time.Duration(ttl)).Err()
	if err != nil {
		panic(err)
	}
}

func DelByKey(key string) {
	err := db.Rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
}

func GetHash(key string, filed string) interface{} {
	result, err := db.Rdb.HGet(ctx, key, filed).Result()
	if err != nil {
		panic(err)
	}
	return result
}

func SetHash(key, filed string, value interface{}) {
	err := db.Rdb.HSet(ctx, key, filed, value).Err()
	if err != nil {
		panic(err)
	}
}

func GetAllHash(key string) map[string]string {
	result, err := db.Rdb.HGetAll(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return result
}
