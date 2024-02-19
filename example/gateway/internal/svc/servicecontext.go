package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"2106A-zg6/gateway/internal/config"
	"2106A-zg6/gateway/internal/middleware"
	gr "2106A-zg6/srv/goods/pb"
	or "2106A-zg6/srv/order/pb"
)

type ServiceContext struct {
	Config config.Config
	//UserRpc  ur.UserCenterServiceClient
	GoodsRpc        gr.GoodsServiceClient
	OrderRpc        or.OrderServiceClient
	AgentMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//UserRpc:  ur.NewUserCenterServiceClient(zrpc.MustNewClient(zrpc.RpcClientConf{Target: "dns:///127.0.0.1:8080"}).Conn()),
		GoodsRpc:        gr.NewGoodsServiceClient(zrpc.MustNewClient(zrpc.RpcClientConf{Target: "dns:///127.0.0.1:8081"}).Conn()),
		OrderRpc:        or.NewOrderServiceClient(zrpc.MustNewClient(zrpc.RpcClientConf{Target: "dns:///127.0.0.1:8082"}).Conn()),
		AgentMiddleware: middleware.NewAgentMiddleware().Handle,
	}
}
