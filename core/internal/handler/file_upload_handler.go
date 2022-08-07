package handler

import (
	"crypto/md5"
	"fmt"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/models"
	"net/http"
	"path"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leexsh/file_sys/core/internal/logic"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// add file
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		// 判定文件是否存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		request := &models.RepositoryPool{}
		has, err := svcCtx.Engine.Where("hash=?", hash).Exist(request)
		if err != nil {
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadResponse{
				Identity: request.Identity,
				Ext:      request.Ext,
				Name:     request.Name,
				Msg:      fmt.Sprintf("文件%s已经存在", fileHeader.Filename),
			})
			return
		}
		cospath, err := helper.UploadCos(r)
		if err != nil {
			return
		}
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cospath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
