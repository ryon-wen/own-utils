// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package server

import (
	"context"

	"2106A-zg6/srv/usercenter/internal/logic"
	"2106A-zg6/srv/usercenter/internal/svc"
	"2106A-zg6/srv/usercenter/pb"
)

type UserCenterServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserCenterServiceServer
}

func NewUserCenterServiceServer(svcCtx *svc.ServiceContext) *UserCenterServiceServer {
	return &UserCenterServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserCenterServiceServer) Create(ctx context.Context, in *pb.CreateReq) (*pb.Empty, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *UserCenterServiceServer) Update(ctx context.Context, in *pb.UpdateReq) (*pb.Empty, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *UserCenterServiceServer) Get(ctx context.Context, in *pb.GetReq) (*pb.GetResp, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}

func (s *UserCenterServiceServer) Delete(ctx context.Context, in *pb.DeleteReq) (*pb.Empty, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}
