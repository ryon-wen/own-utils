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

type FindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLogic {
	return &FindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkFind(in *pb.FindReq) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "invalid find request")
	}
	if in.Limit <= 0 {
		return status.Error(codes.InvalidArgument, "limit must > 0")
	}
	if in.Offset <= 0 {
		return status.Error(codes.InvalidArgument, "offset must > 0")
	}
	return nil
}

func (l *FindLogic) Find(in *pb.FindReq) (*pb.FindResp, error) {
	err := checkFind(in)
	if err != nil {
		return nil, err
	}
	goods, err := model.FindGoods(int(in.Offset), int(in.Limit))
	if err != nil {
		return nil, err
	}
	var g []*pb.Good
	for _, v := range goods {
		g = append(g, model.GoodToPb(v))
	}
	return &pb.FindResp{
		Goods: g,
	}, nil
}
