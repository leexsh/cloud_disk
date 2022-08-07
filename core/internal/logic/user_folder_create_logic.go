package logic

import (
	"context"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
	"leexsh/file_sys/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	// 查询文件夹是否存在
	cnt, err := l.svcCtx.Engine.Where("name=? AND parent_id=?", req.Name, req.ParentId).Count(&models.UserRepository{})
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return &types.UserFolderCreateResponse{
			Identity: "-1",
			Msg:      "文件夹已经存在",
		}, err
	}

	ur := &models.UserRepository{
		Identity:     helper.GenerateUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	resp = &types.UserFolderCreateResponse{
		Identity: ur.Identity,
		Msg:      "创建文件夹成功",
	}
	return
}
