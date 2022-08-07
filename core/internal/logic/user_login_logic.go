package logic

import (
	"context"
	"errors"
	"leexsh/file_sys/core/define"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/models"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 1.从数据库读取数据
	user := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("name= ? AND password= ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或者密码错误")
	}
	// 2.返回token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, int64(define.TokenExpire))
	if err != nil {
		return nil, err
	}

	// 3.生成用于刷新Token的token
	refreshToken, err := helper.GenerateToken(user.Id, user.Identity, user.Name, int64(define.RefreshTokenExpire))
	resp = new(types.LoginResponse)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
