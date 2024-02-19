package goods

import (
	"context"

	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodLogic {
	return &UpdateGoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGoodLogic) UpdateGood(req *types.UpdateGoodReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
