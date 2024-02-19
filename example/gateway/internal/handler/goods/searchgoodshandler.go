package goods

import (
	"net/http"

	"2106A-zg6/gateway/internal/logic/goods"
	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchGoodsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goods.NewSearchGoodsLogic(r.Context(), svcCtx)
		resp, err := l.SearchGoods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
