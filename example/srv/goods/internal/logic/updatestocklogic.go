package logic

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/goods/internal/svc"
	"2106A-zg6/srv/goods/model"
	"2106A-zg6/srv/goods/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockLogic {
	return &UpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkUpdateStock(in *pb.UpdateStockReq) error {
	if in == nil {
		return status.Error(codes.Aborted, "update stock back request cannot be nil")
	}
	if in.GID == "" {
		return status.Error(codes.Aborted, "gid is required")
	}
	return nil
}

func (l *UpdateStockLogic) UpdateStock(in *pb.UpdateStockReq) (*pb.GoodsEmpty, error) {
	err := checkUpdateStock(in)
	if err != nil {
		return nil, err
	}
	err = model.UpdateGoodsStock(in.Goods, in.GID)
	if err != nil {
		zap.S().Debugf("update stock failed: %v", err)
		return nil, status.Errorf(codes.Aborted, "update stock failed: %v", err)
	}
	return &pb.GoodsEmpty{}, nil
}
