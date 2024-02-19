package middleware

import (
	"fmt"
	"net/http"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/zeromicro/go-zero/rest/httpx"

	sen "2106A-zg6/utils/sentinel"
)

type AgentMiddleware struct {
}

func NewAgentMiddleware() *AgentMiddleware {
	return &AgentMiddleware{}
}

func (m *AgentMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sen.RegisterSentinel()
		e, b := sentinel.Entry("gateway", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf(b.Error()))
			return
		} else {
			e.Exit()
		}
		next(w, r)
	}
}
