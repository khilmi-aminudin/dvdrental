package redisdata

import (
	"dvdrental/helper"

	"github.com/go-redis/redis/v8"
)

const NoData = redis.Nil

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	result, err := rdb.Ping(rdb.Context()).Result()
	helper.LogErrorWithFields(err, "ping_result", result)
	helper.Logger().Println(result)

	return rdb
}
