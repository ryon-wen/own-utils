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

type UpdateStockBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockBackLogic {
	return &UpdateStockBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkUpdateStockBack(in *pb.UpdateStockReq) error {
	if in == nil {
		return status.Error(codes.Aborted, "update stock back request cannot be nil")
	}
	return nil
}

func (l *UpdateStockBackLogic) UpdateStockBack(in *pb.UpdateStockReq) (*pb.GoodsEmpty, error) {
	err := checkUpdateStockBack(in)
	if err != nil {
		return nil, err
	}
	if model.GidExist(in.GID) {
		return &pb.GoodsEmpty{}, nil
	}
	err = model.UpdateGoodsStockBack(in.Goods)
	if err != nil {
		zap.S().Debugf("update stock failed: %v", err)
		return nil, status.Errorf(codes.Aborted, "update stock failed: %v", err)
	}
	return &pb.GoodsEmpty{}, nil
}
