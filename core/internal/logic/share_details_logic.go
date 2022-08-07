package logic

import (
	"context"
	"leexsh/file_sys/core/internal/svc"
	"leexsh/file_sys/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareDetailsLogic {
	return &ShareDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareDetailsLogic) ShareDetails(req *types.ShareBasicDetailsRequest) (resp *types.ShareBasicDetailsResponse, err error) {
	// 1.链接的点击次数+1 使用事务
	{
		session := l.svcCtx.Engine.NewSession()
		defer session.Close()
		err := session.Begin()
		_, err = session.Exec("update share_basic set click_num=click_num+1 where identity=?", req.Identity)
		if err != nil {
			return nil, err
		}
		err = session.Commit()
		if err != nil {
			return nil, err
		}
	}
	// 2.查询详细信息
	resp = new(types.ShareBasicDetailsResponse)
	has, err := l.svcCtx.Engine.Table("share_basic").
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Join("LEFT", "user_repository", "user_repository.identity = share_basic.user_repository_identity ").
		Where("share_basic.identity=?", req.Identity).Get(resp)
	if err != nil {
		return nil, err
	}
	if !has {
		return &types.ShareBasicDetailsResponse{
			Msg: "该文件不存在",
		}, err
	}
	resp.Msg = "查询成功"
	return
}
