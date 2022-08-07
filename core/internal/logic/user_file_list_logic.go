package logic

import (
	"context"
	"leexsh/file_sys/core/define"
	"leexsh/file_sys/core/models"
	"time"

	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIndentity string) (resp *types.UserFileListResponse, err error) {
	userfile := make([]*types.UserFile, 0)
	var size int = req.Size
	if size == 0 {
		size = define.Pagesize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id=? AND user_identity=?", req.Id, userIndentity).Select("user_repository.id, "+
		"user_repository.identity, user_repository.repository_identity, user_repository.ext, user_repository.name, repository_pool.path, repository_pool.size").Join("LEFT",
		"repository_pool", "user_repository.repository_identity = repository_pool.identity").Where("user_repository.deleted_at=? OR user_repository.deleted_at=NULL", time.Time{}.Format(define.DateTime)).
		Limit(size, offset).Find(&userfile)

	if err != nil {
		return &types.UserFileListResponse{
			List:  nil,
			Count: 0,
			Msg:   "get error",
		}, err
	}
	count, err := l.svcCtx.Engine.Table("user_repository").Where("parent_id=? AND user_identity=?", req.Id, userIndentity).Count(&models.UserRepository{})
	if err != nil {
		return nil, err
	}
	resp = &types.UserFileListResponse{
		List:  userfile,
		Count: count,
		Msg:   "查询成功!",
	}
	return
}
