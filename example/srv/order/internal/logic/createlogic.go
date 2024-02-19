package logic

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/order/internal/svc"
	"2106A-zg6/srv/order/model"
	"2106A-zg6/srv/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkCreate(in *pb.CreateReq) error {
	if in == nil {
		return status.Error(codes.Aborted, "invalid create request")
	}
	if in.UID <= 0 {
		return status.Error(codes.Aborted, "uid is required")
	}
	if in.OrderNO == "" {
		return status.Error(codes.Aborted, "order_no is required")
	}
	if in.TotalAmount == 0 {
		return status.Error(codes.Aborted, "total_amount is required")
	}
	if in.GoodsList == "" {
		return status.Error(codes.Aborted, "total_amount is required")
	}
	return nil
}

func (l *CreateLogic) Create(in *pb.CreateReq) (*pb.OrderEmpty, error) {
	err := checkCreate(in)
	if err != nil {
		return nil, err
	}
	if in.UID == 1 {
		return nil, status.Error(codes.Aborted, "未知错误")
	}
	err = model.CreateOrder(&model.Orders{
		UserId:      in.UID,
		GoodsList:   in.GoodsList,
		TotalAmount: in.TotalAmount,
		TradeNo:     in.OrderNO,
		TradeStatus: int8(in.Status),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "create order failed: %v", err)
	}
	return &pb.OrderEmpty{}, nil
}
