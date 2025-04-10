package logic

import (
	"context"
	"zore/service/role/rpc/pb/role"
	"zore/service/user/rpc/pb/user"

	"zore/service/user/api/internal/svc"
	"zore/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.List(l.ctx, &user.ListRequest{
		Username: req.Username,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	userIds := make([]int64, len(res.Data))
	for i, userItem := range res.Data {
		userIds[i] = userItem.Id
	}

	roleRes, err := l.svcCtx.RoleRpc.GetRolesByUserIds(l.ctx, &role.GetRolesByUserIdsReq{
		GetRolesByUserIdsReq: userIds,
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.UserItem, len(res.Data))
	for i, item := range res.Data {

		list[i] = types.UserItem{
			Id:        uint64(item.Id),
			ParentId:  item.ParentId,
			Username:  item.Username,
			Password:  "**********",
			Status:    item.Status,
			Ip:        item.Ip,
			LoginLast: item.LoginLast,
		}
		if roleRes.Roles[item.Id] != nil {
			roles := make([]types.Role, len(roleRes.Roles[item.Id].Roles))
			roleArr := roleRes.Roles[item.Id].Roles
			for index, roleItem := range roleArr {
				roles[index] = types.Role{
					RoleName: roleItem.RoleName,
					ParentId: uint64(roleItem.ParentId),
					Id:       uint64(roleItem.Id),
				}
			}

			list[i].Role = roles
		}

	}
	return &types.ListResponse{
		List:     list,
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    res.Count,
	}, nil
}
