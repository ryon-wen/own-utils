package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"2106A-zg6/srv/goods/internal/config"
	"2106A-zg6/srv/goods/internal/server"
	"2106A-zg6/srv/goods/internal/svc"
	"2106A-zg6/srv/goods/model"
	"2106A-zg6/srv/goods/pb"

	"github.com/ryon-wen/own-utils/own"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/goods.yaml", "the config file")

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
	own.InitElasticsearch(&nc.ElasticSearch)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterGoodsServiceServer(grpcServer, server.NewGoodsServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
