package main

import (
	"flag"
	"fmt"

	"github.com/ryon-wen/own-utils/own"

	"2106A-zg6/gateway/internal/config"
	"2106A-zg6/gateway/internal/handler"
	"2106A-zg6/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	own.InitAliPay(&own.AliPay{
		AppID:        c.AliPay.AppID,
		PrivateKey:   c.AliPay.PrivateKey,
		AliPublicKey: c.AliPay.AliPublicKey,
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
