package logic

import (
	"context"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
	"leexsh/file_sys/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepoSaveRequest, userIdentity string) (resp *types.UserRepoSaveResponse, err error) {
	ur := &models.UserRepository{}
	has, err := l.svcCtx.Engine.Where("name=?", req.Name).Exist(ur)
	if err != nil {
		return nil, err
	}
	if has {
		return &types.UserRepoSaveResponse{
			Msg:      "Error: file existed",
			Inentity: "error",
		}, err
	}
	ur.Identity = helper.GenerateUUID()
	ur.UserIdentity = userIdentity
	ur.ParentId = req.ParentId
	ur.RepositoryIdentity = req.RepoInentityId
	ur.Ext = req.Ext
	ur.Name = req.Name

	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	resp = &types.UserRepoSaveResponse{}
	resp.Msg = "success"
	resp.Inentity = ur.Identity
	return
}
