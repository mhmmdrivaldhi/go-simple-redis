package db

import "github.com/redis/go-redis/v9"

var rdb *redis.Client

func RedisInit() error {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       2,
	})

	rdb = db

	return nil
}

func RedisConnect() *redis.Client{
	return rdb
}