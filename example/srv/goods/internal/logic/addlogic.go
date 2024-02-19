package logic

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/goods/internal/svc"
	"2106A-zg6/srv/goods/model"
	"2106A-zg6/srv/goods/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkAdd(in *pb.AddReq) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "invalid add request")
	}
	if in.Good.Price == 0 {
		return status.Error(codes.InvalidArgument, "good price is required")
	}
	if in.Good.Name == "" {
		return status.Error(codes.InvalidArgument, "good name is required")
	}
	if in.Good.Stock <= 0 {
		return status.Error(codes.InvalidArgument, "good stock must > 0")
	}
	return nil
}

func (l *AddLogic) Add(in *pb.AddReq) (*pb.GoodsEmpty, error) {
	err := checkAdd(in)
	if err != nil {
		return nil, err
	}
	err = model.AddGoods(&model.Goods{
		Name:      in.Good.Name,
		Price:     in.Good.Price,
		Stock:     in.Good.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add good: %v", err)
	}
	return &pb.GoodsEmpty{}, nil
}
