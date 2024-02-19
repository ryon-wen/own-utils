package logic

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/goods/internal/svc"
	"2106A-zg6/srv/goods/model"
	"2106A-zg6/srv/goods/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkUpdate(in *pb.UpdateReq) error {
	if in == nil {
		return fmt.Errorf("invalid add request")
	}
	if in.Good.ID == 0 {
		return status.Error(codes.InvalidArgument, "good id is required")
	}
	if in.Good.Stock <= 0 {
		return status.Error(codes.InvalidArgument, "good stock must > 0")
	}
	return nil
}

func (l *UpdateLogic) Update(in *pb.UpdateReq) (*pb.GoodsEmpty, error) {
	err := checkUpdate(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = model.UpdateGoods(&model.Goods{
		Id:        in.Good.ID,
		Name:      in.Good.Name,
		Price:     in.Good.Price,
		Stock:     in.Good.Stock,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update good: %v", err)
	}
	return &pb.GoodsEmpty{}, nil
}
