package goods

import (
	"context"

	"github.com/ryon-wen/own-utils/code"

	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"
	"2106A-zg6/srv/goods/goodsservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodLogic {
	return &GetGoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodLogic) GetGood(req *types.GetGoodReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp, err = &types.Response{}, nil
	rpc, err := l.svcCtx.GoodsRpc.Get(l.ctx, &goodsservice.GetReq{
		Id: req.Id,
	})
	if err != nil {
		resp.Code = code.InternalError
		resp.Msg = err.Error()
		return
	}
	resp.Code = code.Success
	resp.Msg = "Success"
	resp.Data = rpc.Good
	return
}
