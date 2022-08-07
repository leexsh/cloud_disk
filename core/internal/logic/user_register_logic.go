package logic

import (
	"context"
	"errors"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/models"
	"log"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	code, err := l.svcCtx.RedisEngine.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("验证码不正确")
	}
	if code != req.Code {
		return nil, errors.New("验证码错误")
	}
	// 验证结束， 判定用户是否存在
	has, err := l.svcCtx.Engine.Where("name=?", req.Name).Exist(&models.UserBasic{})
	if err != nil {
		return nil, err
	}
	if has {
		return nil, errors.New("用户名已经存在")
	}

	// 写入用户
	uc := &models.UserBasic{
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
		Identity: helper.GenerateUUID(),
	}
	n, err := l.svcCtx.Engine.Insert(uc)
	if err != nil {
		return nil, err
	}
	log.Println("insert row is: ", n)
	resp = &types.UserRegisterResponse{}
	resp.Msg = "注册成功"
	return
}
