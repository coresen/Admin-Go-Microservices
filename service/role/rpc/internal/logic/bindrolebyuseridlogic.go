package logic

import (
	"context"
	"zore/service/role/model"

	"zore/service/role/rpc/internal/svc"
	"zore/service/role/rpc/pb/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindRoleByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindRoleByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindRoleByUserIdLogic {
	return &BindRoleByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindRoleByUserIdLogic) BindRoleByUserId(in *role.BindRoleByUserRequest) (*role.BindRoleByUserResponse, error) {
	// todo: add your logic here and delete this line

	_ = l.svcCtx.UserRoleModel.DeleteByUserId(l.ctx, in.UserId)

	roles := make([]*model.UserRole, len(in.RoleId))
	for i, v := range in.RoleId {
		roles[i] = &model.UserRole{
			RoleId: v,
			UserId: in.UserId,
		}
	}
	err := l.svcCtx.UserRoleModel.BatchInsert(l.ctx, roles)
	if err != nil {
		return nil, err
	}
	return &role.BindRoleByUserResponse{}, nil
}
