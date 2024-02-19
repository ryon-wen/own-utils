package goods

import (
	"context"

	"github.com/ryon-wen/own-utils/code"

	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"
	"2106A-zg6/srv/goods/goodsservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodsLogic {
	return &SearchGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchGoodsLogic) SearchGoods(req *types.SearchGoodsReq) (resp *types.Response, err error) {
	resp, err = &types.Response{}, nil
	rpc, rpcErr := l.svcCtx.GoodsRpc.SearchGoods(l.ctx, &goodsservice.SearchGoodsReq{
		Content: req.Content,
		Offset:  req.Offset,
		Limit:   req.Limit,
	})
	if rpcErr != nil {
		resp.Code = code.InternalError
		resp.Msg = rpcErr.Error()
		return
	}
	resp.Code = code.Success
	resp.Msg = "success"
	resp.Data = rpc.Goods
	return
}
