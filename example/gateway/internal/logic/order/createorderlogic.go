package order

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/google/uuid"
	"github.com/ryon-wen/own-utils/code"
	"github.com/ryon-wen/own-utils/own"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/core/logx"

	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"
	"2106A-zg6/srv/goods/goodsservice"
	gp "2106A-zg6/srv/goods/pb"
	op "2106A-zg6/srv/order/pb"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.Response, err error) {
	resp, err = &types.Response{}, nil

	var goodsId []int64
	for _, v := range req.Goods {
		if v.Quantity == 0 {
			resp.Code = code.BadRequest
			resp.Msg = "请求错误，商品数量不能为空"
			return
		}
		goodsId = append(goodsId, v.Id)
	}
	// 根据id查询所有商品
	rpc, err := l.svcCtx.GoodsRpc.FindByIds(l.ctx, &goodsservice.FindByIDsReq{
		IDs: goodsId,
	})
	fmt.Println(err)
	if err != nil {
		resp.Code = code.InternalError
		resp.Msg = err.Error()
		return
	}
	// 判断请求id与数据库id是否一致
	if len(req.Goods) != len(rpc.Goods) {
		resp.Code = code.BadRequest
		resp.Msg = "请求错误，商品不存在"
		return
	}
	// 验证库存是否充足
	goodsMap := make(map[int64]int64)
	for _, v := range req.Goods {
		goodsMap[v.Id] = v.Quantity
	}
	for _, v := range rpc.Goods {
		if v.Stock < goodsMap[v.ID] {
			resp.Code = code.BadRequest
			resp.Msg = "商品库存不足"
			return
		}
	}
	// 计算总金额+定义请求体
	var gs []*gp.GoodsStock
	var totalAmount float64
	for _, v := range rpc.Goods {
		totalAmount += v.Price * float64(goodsMap[v.ID])
		gs = append(gs, &gp.GoodsStock{
			ID:    v.ID,
			Stock: goodsMap[v.ID],
		})
	}
	// 生成订单编号
	orderNo := uuid.NewString()
	// 生成订单subject
	b, _ := json.Marshal(&req.Goods)
	// 生成saga事务 id
	gid := uuid.NewString()
	// 定义请求体
	goodsReq := gp.UpdateStockReq{
		Goods: gs,
		GID:   gid,
	}
	orderReq := op.CreateReq{
		OrderNO:     orderNo,
		UID:         req.UID,
		GoodsList:   string(b),
		TotalAmount: totalAmount,
	}
	// 创建saga事务
	saga := dtmgrpc.NewSagaGrpc(dtmutil.DefaultGrpcServer, gid).
		Add("192.168.0.3:8081"+"/pb.GoodsService/UpdateStock", "192.168.0.3:8081"+"/pb.GoodsService/UpdateStockBack", &goodsReq).
		Add("192.168.0.3:8082"+"/pb.OrderService/Create", "192.168.0.3:8082"+"/pb.OrderService/CreateBack", &orderReq)
	saga.WaitResult = true
	err = saga.Submit()
	fmt.Println(err)
	if err != nil {
		resp.Code = code.InternalError
		resp.Msg = err.Error()
		return
	}
	// 生成支付链接
	var p = alipay.TradePagePay{}
	p.NotifyURL = l.svcCtx.Config.AliPay.NotifyURL
	p.Subject = string(b)
	p.OutTradeNo = orderNo
	p.TotalAmount = fmt.Sprintf("%.2f", totalAmount)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	payUrl, err := own.GeneralPayURL(p)
	if err != nil {
		resp.Code = code.OperationFailed
		resp.Msg = err.Error()
		return
	}

	resp.Code = code.Success
	resp.Msg = "下单成功！"
	resp.Data = map[string]interface{}{
		"payUrl": payUrl,
	}
	return
}
