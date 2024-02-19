package goods

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"2106A-zg6/gateway/internal/logic/goods"
	"2106A-zg6/gateway/internal/svc"
	"2106A-zg6/gateway/internal/types"
)

func GetGoodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetGoodReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goods.NewGetGoodLogic(r.Context(), svcCtx)
		resp, err := l.GetGood(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
