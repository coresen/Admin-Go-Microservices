package logic

import (
	"context"
	"role-rpc/internal/svc"
	"role-rpc/pb/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolesByUserIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRolesByUserIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolesByUserIdsLogic {
	return &GetRolesByUserIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRolesByUserIdsLogic) GetRolesByUserIds(in *role.GetRolesByUserIdsReq) (*role.GetRolesByUserIdsResp, error) {
	// todo: add your logic here and delete this line
	if len(in.GetRolesByUserIdsReq) == 0 {
		return &role.GetRolesByUserIdsResp{}, nil
	}
	roleResp, err := l.svcCtx.UserRoleModel.FindRolesByUserIds(l.ctx, in.GetRolesByUserIdsReq)
	if err != nil {
		return nil, err
	}

	if roleResp == nil {
		return &role.GetRolesByUserIdsResp{}, nil
	}

	roles := make(map[int64]*role.RoleList)

	for i, v := range roleResp {
		roleItem := make([]*role.DetailResponse, 0)
		for _, val := range v {
			roleItem = append(roleItem, &role.DetailResponse{
				Id:          val.Id,
				ParentId:    val.ParentId,
				RoleName:    val.RoleName,
				Description: val.Description,
				Status:      int32(val.Status),
			})
		}
		roles[i] = &role.RoleList{
			Roles: roleItem,
		}
	}

	return &role.GetRolesByUserIdsResp{
		Roles: roles,
	}, nil
}
