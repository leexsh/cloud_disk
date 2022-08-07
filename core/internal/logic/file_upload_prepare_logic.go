package logic

import (
	"context"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
	"leexsh/file_sys/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareResponse, err error) {
	rp := &models.RepositoryPool{}
	has, err := l.svcCtx.Engine.Where("hash=?", req.Md5).Get(rp)
	if err != nil {
		return nil, err
	}
	// 文件存在
	resp = &types.FileUploadPrepareResponse{}
	if has {
		resp.Identity = rp.Identity
		resp.UploadId = "-1"
		resp.Msg = "已经存在这文件，可以走秒传"
	} else {
		// 文件不存在 需要分片上传
		key, uploadId, err := helper.CosInitPart(req.Ext)
		if err != nil {
			return nil, err
		}
		resp.Identity = rp.Identity
		resp.Key = key
		resp.UploadId = uploadId
		resp.Msg = "文件不存在，需要分片上传"
	}
	return
}
