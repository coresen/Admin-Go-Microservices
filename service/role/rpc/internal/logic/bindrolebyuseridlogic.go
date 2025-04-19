package logic

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
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
	fmt.Println("进入", in)
	err := l.svcCtx.UserRoleModel.DeleteByUserId(l.ctx, in.UserId)
	fmt.Println("err:", err)
	if err != nil {
		return nil, err
	}

	node, err := snowflake.NewNode(1) // 1 是节点ID，可以根据实际需要设置不同的值
	if err != nil {
		return nil, err
	}

	fmt.Println("asd")
	if len(in.RoleId) == 0 {
		return &role.BindRoleByUserResponse{}, nil
	}

	roles := make([]*model.UserRole, len(in.RoleId))
	for i, v := range in.RoleId {
		roles[i] = &model.UserRole{
			Id:     node.Generate().Int64(),
			RoleId: v,
			UserId: in.UserId,
		}
	}
	err = l.svcCtx.UserRoleModel.BatchInsert(l.ctx, roles)
	if err != nil {
		return nil, err
	}
	return &role.BindRoleByUserResponse{}, nil
}
