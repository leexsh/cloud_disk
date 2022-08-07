package logic

import (
	"context"
	"errors"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
	"leexsh/file_sys/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveSharedFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveSharedFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveSharedFileLogic {
	return &SaveSharedFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveSharedFileLogic) SaveSharedFile(req *types.SaveSharedFileRequest, userIndentity string) (resp *types.SaveSharedFileResponse, err error) {
	//获取资源详情
	rp := &models.RepositoryPool{}
	has, err := l.svcCtx.Engine.Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件不存在")
	}
	// 写入数据
	ur := models.UserRepository{
		Identity:           helper.GenerateUUID(),
		UserIdentity:       userIndentity,
		ParentId:           int(req.ParentId),
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	_, err = l.svcCtx.Engine.Insert(&ur)
	if err != nil {
		return nil, err
	}
	resp = &types.SaveSharedFileResponse{
		Identity: ur.Identity,
		Msg:      "写入成功",
	}
	return
}
