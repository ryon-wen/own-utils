package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/usercenter/internal/model"
	"2106A-zg6/srv/usercenter/internal/svc"
	"2106A-zg6/srv/usercenter/pb"

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
		return status.Error(codes.InvalidArgument, "invalid create request")
	}
	if in.ID == 0 {
		return status.Error(codes.InvalidArgument, "ID is required")
	}
	return nil
}

func (l *GetLogic) Get(in *pb.GetReq) (*pb.GetResp, error) {
	err := checkGet(in)
	if err != nil {
		return nil, err
	}
	user, err := model.GetUser(in.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user failed: %v", err)
	}
	return &pb.GetResp{
		User: &pb.User{
			ID:       user.Id,
			Nickname: user.Nickname,
			Mobile:   user.Mobile,
			Password: user.Password,
		},
	}, nil
}
