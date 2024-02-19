package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/goods/internal/svc"
	"2106A-zg6/srv/goods/model"
	"2106A-zg6/srv/goods/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkDelete(in *pb.DeleteReq) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "invalid delete request")
	}
	if in.Id == 0 {
		return status.Error(codes.InvalidArgument, "good id is required")
	}
	return nil
}

func (l *DeleteLogic) Delete(in *pb.DeleteReq) (*pb.GoodsEmpty, error) {
	err := checkDelete(in)
	if err != nil {
		return nil, err
	}
	err = model.DeleteGoods(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete good: %v", err)
	}
	return &pb.GoodsEmpty{}, nil
}
