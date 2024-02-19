package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/ryon-wen/own-utils/own"

	"2106A-zg6/srv/order/internal/config"
	"2106A-zg6/srv/order/internal/server"
	"2106A-zg6/srv/order/internal/svc"
	"2106A-zg6/srv/order/model"
	"2106A-zg6/srv/order/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	content, err := own.InitNacos(&c.Nacos)
	var nc config.NacosConfig
	err = json.Unmarshal([]byte(content), &nc)
	if err != nil {
		panic("unmarshalNacos failed" + err.Error())
	}

	ctx := svc.NewServiceContext(c, nc)

	model.RegisterModel(&nc.MySQL)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServiceServer(grpcServer, server.NewOrderServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
