package logic

import (
	"context"
	"errors"
	"leexsh/file_sys/core/models"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailsLogic {
	return &UserDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailsLogic) UserDetails(req *types.UserDetailsRequest) (resp *types.UserDetailsResponse, err error) {
	user := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("identity=?", req.Indentity).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	resp = new(types.UserDetailsResponse)
	resp.Name = user.Name
	resp.Email = user.Email
	return
}
