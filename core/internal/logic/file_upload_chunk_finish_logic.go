package logic

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"leexsh/file_sys/core/define"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/models"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkFinishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkFinishLogic {
	return &FileUploadChunkFinishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkFinishLogic) FileUploadChunkFinish(req *types.FileChunkUploadFinishRequest) (resp *types.FileChunkUploadFinishResponse, err error) {
	co := make([]cos.Object, 0)
	for _, object := range req.CosObjects {
		co = append(co, cos.Object{
			ETag:       object.Etag,
			PartNumber: object.PartNumber,
		})
	}
	err = helper.CosChunkUploadFinish(req.Key, req.UploadId, co)
	resp = &types.FileChunkUploadFinishResponse{}
	if err != nil {
		resp.Msg = "upload fail"
		return resp, err

	}
	// 数据入库
	rp := &models.RepositoryPool{
		Identity: helper.GenerateUUID(),
		Hash:     req.Md5,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     define.COSADDR + "/" + req.Key,
	}
	l.svcCtx.Engine.Insert(rp)
	resp.Msg = "upload success"
	return
}
