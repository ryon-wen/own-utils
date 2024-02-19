package config

import (
	"github.com/ryon-wen/own-utils/own"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Nacos own.Nacos
}

type NacosConfig struct {
	MySQL         own.MySQL         `json:"mysql"`
	Redis         own.Redis         `json:"redis"`
	ElasticSearch own.ElasticSearch `json:"es"`
}
