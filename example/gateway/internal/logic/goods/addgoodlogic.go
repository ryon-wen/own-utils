package goods

import (
	"context"

	"github.com/ryon-wen/own-utils/code"

	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"
	"2106A-zg6/srv/goods/goodsservice"
	"2106A-zg6/srv/goods/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodLogic {
	return &AddGoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGoodLogic) AddGood(req *types.AddGoodReq) (resp *types.Response, err error) {
	resp, err = &types.Response{}, nil

	_, err = l.svcCtx.GoodsRpc.Add(l.ctx, &goodsservice.AddReq{
		Good: &pb.Good{
			Name:  req.Name,
			Price: req.Price,
			Stock: req.Stock,
		},
	})
	if err != nil {
		resp.Code = code.InternalError
		resp.Msg = err.Error()
		return
	}

	resp.Code = code.Success
	resp.Msg = "添加成功！"
	return
}
