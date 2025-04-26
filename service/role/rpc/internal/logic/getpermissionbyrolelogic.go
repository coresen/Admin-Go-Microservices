package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"role-rpc/internal/svc"
	"role-rpc/pb/role"
)

type GetPermissionByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionByRoleLogic {
	return &GetPermissionByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPermissionByRoleLogic) GetPermissionByRole(in *role.GetPermissionByRoleRequest) (*role.GetPermissionByRoleResponse, error) {
	res, err := l.svcCtx.PermissionsModel.FindRoleList(l.ctx, in.RoleIds)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	data := make([]*role.PermissionItem, len(res))
	for i, item := range res {
		data[i] = &role.PermissionItem{
			Id:     int64(item.Id),
			Path:   item.Path,
			Method: item.Method,
		}
	}

	return &role.GetPermissionByRoleResponse{
		Data: data,
	}, nil
}
