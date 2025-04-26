package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	model "role-model"
	"role-rpc/internal/svc"
	"role-rpc/pb/role"
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

	node, err := snowflake.NewNode(1) // 1 是节点ID，可以根据实际需要设置不同的值
	if err != nil {
		return nil, err
	}

	data := make([]*model.Permissions, len(in.Data))
	for i, item := range in.GetData() {
		id := node.Generate().Int64()
		data[i] = &model.Permissions{
			Id:           id,
			RoleId:       item.GetRoleId(),
			PermissionId: item.GetPermissionId(),
		}
	}
	err = l.svcCtx.PermissionsModel.BatchInsert(l.ctx, data)
	if err != nil {
		return nil, err
	}

	return &role.PermissionCreateResponse{}, nil
}
