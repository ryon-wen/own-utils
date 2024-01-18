package dao

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

func InitCache(host string, port int, DB int, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		DB:       DB,
		Password: password,
	})
}
