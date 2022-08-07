package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
	"leexsh/file_sys/core/models"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateResponse, err error) {
	uuid := helper.GenerateUUID()
	ur := &models.UserRepository{}
	has, err := l.svcCtx.Engine.Where("identity=?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("没有发现该文件")
	}
	data := &models.ShareBasic{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     ur.RepositoryIdentity,
		ExpiredTime:            req.ExpireTime,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	l.svcCtx.Engine.ShowSQL(true)
	if err != nil {
		return nil, err
	}
	resp = &types.ShareBasicCreateResponse{
		Identity: uuid,
		Msg:      "创建分享链接成功",
	}
	return
}
