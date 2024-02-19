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

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkGet(in *pb.GetReq) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "Invalid get request")
	}
	if in.Id == 0 {
		return status.Error(codes.InvalidArgument, "good id is required")
	}
	return nil
}

func (l *GetLogic) Get(in *pb.GetReq) (*pb.GetResp, error) {
	err := checkGet(in)
	if err != nil {
		return nil, err
	}
	good, err := model.GetGoods(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get good: %v", err)
	}
	return &pb.GetResp{
		Good: model.GoodToPb(good),
	}, nil
}
