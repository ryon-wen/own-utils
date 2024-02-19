package logic

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/order/internal/svc"
	"2106A-zg6/srv/order/model"
	"2106A-zg6/srv/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBackLogic {
	return &CreateBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkCreateBack(in *pb.CreateReq) error {
	if in == nil {
		return status.Error(codes.Aborted, "invalid create request")
	}
	if in.OrderNO == "" {
		return status.Error(codes.Aborted, "order_no is required")
	}
	return nil
}

func (l *CreateBackLogic) CreateBack(in *pb.CreateReq) (*pb.OrderEmpty, error) {
	fmt.Println("Create back logic")
	err := checkCreateBack(in)
	if err != nil {
		return nil, err
	}
	err = model.DeleteOrder(in.OrderNO)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "failed to delete order: %v", err)
	}
	return &pb.OrderEmpty{}, nil
}
