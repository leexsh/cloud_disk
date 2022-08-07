package logic

import (
	"context"
	"errors"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"
	"leexsh/file_sys/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveResponse, err error) {
	parentData := &models.UserRepository{}
	has, err := l.svcCtx.Engine.Where("identity=? AND user_identity=?", req.ParentIdentity, userIdentity).Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件夹不存在")
	}

	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(models.UserRepository{
		ParentId: parentData.Id,
	})
	l.svcCtx.Engine.ShowSQL(true)
	if err != nil {
		return nil, err
	}
	resp = &types.UserFileMoveResponse{Msg: "移动成功"}
	return
}
