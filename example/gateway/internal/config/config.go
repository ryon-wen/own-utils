package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	AliPay struct {
		AppID        string
		PrivateKey   string
		AliPublicKey string
		NotifyURL    string
	}
}
