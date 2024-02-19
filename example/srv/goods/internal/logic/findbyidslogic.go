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

type FindByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdsLogic {
	return &FindByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkFindByIds(in *pb.FindByIDsReq) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "invalid findByIds request")
	}
	return nil
}

func (l *FindByIdsLogic) FindByIds(in *pb.FindByIDsReq) (*pb.FindResp, error) {
	err := checkFindByIds(in)
	if err != nil {
		return nil, err
	}
	goods, err := model.FindByIds(in.IDs)
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
