package logic

import (
	"context"

	"zore/service/role/rpc/internal/svc"
	"zore/service/role/rpc/pb/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermissionLogic {
	return &CreatePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePermissionLogic) CreatePermission(in *role.PermissionCreateRequest) (*role.PermissionCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &role.PermissionCreateResponse{}, nil
}
