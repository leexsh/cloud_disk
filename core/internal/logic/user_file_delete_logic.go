package logic

import (
	"context"
	"errors"
	"leexsh/file_sys/core/models"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteResponse, err error) {
	_, err = l.svcCtx.Engine.Where("user_identity = ? AND identity = ?", userIdentity, req.Identity).Delete(&models.UserRepository{})
	if err != nil {
		return nil, errors.New("删除失败")
	}
	resp = &types.UserFileDeleteResponse{Msg: "文件删除成功"}
	return
}
