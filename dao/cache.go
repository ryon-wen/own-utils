package dao

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"
)

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DB       int    `json:"db"`
	Password string `json:"password"`
}

func InitCache(r *Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		DB:       r.DB,
		Password: r.Password,
	})
	str, err := rdb.Ping().Result()
	if err != nil {
		panic("failed to connect to redis: " + err.Error())
	}
	log.Printf("connect to redis success: %s", str)
	return rdb
}
