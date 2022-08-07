package logic

import (
	"context"
	"leexsh/file_sys/core/define"
	"leexsh/file_sys/core/helper"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefrshAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefrshAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefrshAuthLogic {
	return &RefrshAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefrshAuthLogic) RefrshAuth(req *types.RefreshAuthRequest, authorization string) (resp *types.RefreshAuthResponse, err error) {
	uc, err := helper.AnalyzeToke(authorization)
	if err != nil {
		return nil, err
	}
	token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	refreshToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = &types.RefreshAuthResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
	resp.Msg = "更新Token成功"
	return
}
