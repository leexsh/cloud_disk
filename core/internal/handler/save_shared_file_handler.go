package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leexsh/file_sys/core/internal/logic"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
)

func SaveSharedFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveSharedFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSaveSharedFileLogic(r.Context(), svcCtx)
		resp, err := l.SaveSharedFile(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
