package goods

import (
	"context"

	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindGoodsLogic {
	return &FindGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindGoodsLogic) FindGoods(req *types.FindGoodsReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
