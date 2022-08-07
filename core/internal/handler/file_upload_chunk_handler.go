package handler

import (
	"errors"
	"leexsh/file_sys/core/helper"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leexsh/file_sys/core/internal/logic"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileChunkUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if r.PostForm.Get("key") == "" {
			httpx.Error(w, errors.New("key is empty"))
			return
		}
		if r.PostForm.Get("uploadId") == "" {
			httpx.Error(w, errors.New("upload id is empty"))
			return
		}
		if r.PostForm.Get("partNumber") == "" {
			httpx.Error(w, errors.New("partNumber is empty"))
			return
		}
		etag, err := helper.CosPartUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)
		resp = &types.FileChunkUploadResponse{Etag: etag}
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
