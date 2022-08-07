package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leexsh/file_sys/core/internal/logic"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
)

func FileUploadPrepareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadPrepareRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadPrepareLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadPrepare(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
