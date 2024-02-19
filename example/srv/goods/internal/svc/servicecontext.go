package svc

import "2106A-zg6/srv/goods/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config, nc config.NacosConfig) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
