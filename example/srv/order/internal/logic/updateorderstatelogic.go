package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/order/internal/svc"
	"2106A-zg6/srv/order/model"
	"2106A-zg6/srv/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStateLogic {
	return &UpdateOrderStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkUpdateOrderState(in *pb.UpdateOrderStateReq) error {
	if in == nil {
		return status.Error(codes.Aborted, "invalid update order state")
	}
	if in.OrderNO == "" {
		return status.Error(codes.Aborted, "order_no is required")
	}
	if in.Status <= 0 {
		return status.Error(codes.Aborted, "status is required")
	}
	return nil
}

func (l *UpdateOrderStateLogic) UpdateOrderState(in *pb.UpdateOrderStateReq) (*pb.OrderEmpty, error) {
	err := checkUpdateOrderState(in)
	if err != nil {
		return nil, err
	}
	err = model.UpdateOrder(in.OrderNO, int8(in.Status))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "update order status error: %v", err)
	}
	return &pb.OrderEmpty{}, nil
}
