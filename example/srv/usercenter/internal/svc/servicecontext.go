package svc

import (
	"github.com/go-redis/redis/v7"
	"github.com/ryon-wen/own-utils/own"

	"2106A-zg6/srv/usercenter/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Rdb    *redis.Client
}

func NewServiceContext(c config.Config, nc config.NacosConfig) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Rdb: own.InitCache(&own.Redis{
			Host:     nc.Redis.Host,
			Port:     nc.Redis.Port,
			DB:       nc.Redis.DB,
			Password: nc.Redis.Password,
		}),
	}
}
