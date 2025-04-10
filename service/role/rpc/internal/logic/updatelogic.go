package logic

import (
	"context"

	"zore/service/role/rpc/internal/svc"
	"zore/service/role/rpc/pb/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *role.UpdateRequest) (*role.UpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &role.UpdateResponse{}, nil
}
