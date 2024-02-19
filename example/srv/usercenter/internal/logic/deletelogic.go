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
		return status.Error(codes.InvalidArgument, "invalid create request")
	}
	if in.ID == 0 {
		return status.Error(codes.InvalidArgument, "ID is required")
	}
	return nil
}

func (l *DeleteLogic) Delete(in *pb.DeleteReq) (*pb.Empty, error) {
	err := checkDelete(in)
	if err != nil {
		return nil, err
	}
	err = model.DeleteUser(in.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete user error: %v", err)
	}
	return &pb.Empty{}, nil
}
