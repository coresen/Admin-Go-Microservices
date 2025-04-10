package logic

import (
	"context"
	"zore/service/role/rpc/pb/role"
	"zore/service/user/rpc/pb/user"

	"zore/service/user/api/internal/svc"
	"zore/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.Update(l.ctx, &user.UpdateUserRequest{
		Id:       int64(req.Id),
		ParentId: req.ParentId,
		Password: req.Password,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.RoleRpc.BindRoleByUserId(l.ctx, &role.BindRoleByUserRequest{
		UserId: int64(req.Id),
		RoleId: req.RoleId,
	})

	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{
		Id: req.Id,
	}, nil

}
