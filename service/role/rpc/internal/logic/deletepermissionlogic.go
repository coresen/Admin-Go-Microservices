package logic

import (
	"context"
	"role-rpc/internal/svc"
	"role-rpc/pb/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermissionLogic {
	return &DeletePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePermissionLogic) DeletePermission(in *role.PermissionDeleteRequest) (*role.PermissionDeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &role.PermissionDeleteResponse{}, nil
}
