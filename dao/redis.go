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
