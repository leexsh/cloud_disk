package logic

import (
	"context"
	"errors"
	"leexsh/file_sys/core/models"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateResponse, err error) {
	// 判定name是否已经存在
	cnt, err := l.svcCtx.Engine.Where("name=? AND parent_id=(SELECT parent_id FROM user_repository WHERE user_repository.identity=?)", req.Name, req.Identity).Count(&models.UserRepository{})
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return &types.UserFileNameUpdateResponse{Msg: "文件名称已经存在"}, errors.New("err: file existed")
	}
	data := &models.UserRepository{Name: req.Name}
	_, err = l.svcCtx.Engine.Where("identity=? AND user_identity=?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return &types.UserFileNameUpdateResponse{Msg: "update fail"}, err
	}
	resp = &types.UserFileNameUpdateResponse{
		Msg: "update success",
	}
	return
}
