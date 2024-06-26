// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"2106A-zg6/srv/order/internal/logic"
	"2106A-zg6/srv/order/internal/svc"
	"2106A-zg6/srv/order/pb"
)

type OrderServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(svcCtx *svc.ServiceContext) *OrderServiceServer {
	return &OrderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServiceServer) Create(ctx context.Context, in *pb.CreateReq) (*pb.OrderEmpty, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *OrderServiceServer) CreateBack(ctx context.Context, in *pb.CreateReq) (*pb.OrderEmpty, error) {
	l := logic.NewCreateBackLogic(ctx, s.svcCtx)
	return l.CreateBack(in)
}

func (s *OrderServiceServer) UpdateOrderState(ctx context.Context, in *pb.UpdateOrderStateReq) (*pb.OrderEmpty, error) {
	l := logic.NewUpdateOrderStateLogic(ctx, s.svcCtx)
	return l.UpdateOrderState(in)
}
