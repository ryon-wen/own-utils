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

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkCreate(in *pb.CreateReq) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "invalid create request")
	}
	if in.User.Mobile == "" {
		return status.Error(codes.InvalidArgument, "mobile is required")
	}
	if in.User.Nickname == "" {
		return status.Error(codes.InvalidArgument, "nickname is required")
	}
	if in.User.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}

func (l *CreateLogic) Create(in *pb.CreateReq) (*pb.Empty, error) {
	err := checkCreate(in)
	if err != nil {
		return nil, err
	}
	err = model.CreateUser(&model.User{
		Nickname:  in.User.Nickname,
		Mobile:    in.User.Mobile,
		Password:  in.User.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "create user failed: %v", err)
	}
	return &pb.Empty{}, nil
}
