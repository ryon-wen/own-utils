package goods

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"2106A-zg6/gateway/internal/logic/goods"
	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"
)

func FindGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindGoodsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goods.NewFindGoodsLogic(r.Context(), svcCtx)
		resp, err := l.FindGoods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
