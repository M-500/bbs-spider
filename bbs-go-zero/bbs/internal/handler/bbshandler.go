package handler

import (
	"net/http"

	"bbs-go-zero/bbs/internal/logic"
	"bbs-go-zero/bbs/internal/svc"
	"bbs-go-zero/bbs/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BbsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBbsLogic(r.Context(), svcCtx)
		resp, err := l.Bbs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
