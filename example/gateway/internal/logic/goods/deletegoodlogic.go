package goods

import (
	"context"

	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodLogic {
	return &DeleteGoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGoodLogic) DeleteGood(req *types.DeleteGoodReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
