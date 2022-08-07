package logic

import (
	"context"
	"errors"
	"fmt"
	"leexsh/file_sys/core/define"
	"leexsh/file_sys/core/helper"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
	"leexsh/file_sys/core/models"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.EmailCodeRequest) (resp *types.EmailCodeResponse, err error) {
	// 1.查询邮箱是否存在
	has, err := l.svcCtx.Engine.Where("email=?", req.Email).Exist(&models.UserBasic{})
	if err != nil {
		return nil, err
	}
	fmt.Println("liz has,", has)
	resp = &types.EmailCodeResponse{}
	if has {
		errors.New("email existed")
		resp.Msg = "邮箱已经存在，请换个邮箱"
		resp.Code = "-1"
		return resp, err
	}

	// 2.邮箱不存在
	code := helper.GenerateEmailCode()
	l.svcCtx.RedisEngine.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExprie))
	resp.Code = code
	resp.Msg = fmt.Sprintf("已经向%s的邮箱成功发送验证码, 验证码是%s", req.Email, code)
	return
}
