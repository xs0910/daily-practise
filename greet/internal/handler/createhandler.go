package handler

import (
	"net/http"

	"daily-practise/greet/internal/logic"
	"daily-practise/greet/internal/svc"
	"daily-practise/greet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		err := l.Create(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
