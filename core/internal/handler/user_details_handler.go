package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leexsh/file_sys/core/internal/logic"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
)

func UserDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserDetailsLogic(r.Context(), svcCtx)
		resp, err := l.UserDetails(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
