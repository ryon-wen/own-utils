package logic

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/usercenter/internal/model"
	"2106A-zg6/srv/usercenter/internal/svc"
	"2106A-zg6/srv/usercenter/pb"

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
		return status.Error(codes.InvalidArgument, "invalid update request")
	}
	if in.User.ID == 0 {
		return status.Error(codes.InvalidArgument, "user id is required")
	}
	return nil
}

func (l *UpdateLogic) Update(in *pb.UpdateReq) (*pb.Empty, error) {
	err := checkUpdate(in)
	if err != nil {
		return nil, err
	}
	err = model.UpdateUser(&model.User{
		Id:        in.User.ID,
		Nickname:  in.User.Nickname,
		Mobile:    in.User.Mobile,
		Password:  in.User.Password,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
