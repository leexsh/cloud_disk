package logic

import (
	"context"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkLogic {
	return &FileUploadChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkLogic) FileUploadChunk(req *types.FileChunkUploadRequest) (resp *types.FileChunkUploadResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
