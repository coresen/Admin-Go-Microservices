package logic

import (
	"context"
	"errors"
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
	detail, err := l.svcCtx.UserRpc.Detail(l.ctx, &user.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	if detail == nil {
		return nil, errors.New("user is nil")
	}

	data := user.UpdateUserRequest{
		Id:       req.Id,
		ParentId: req.ParentId,
	}

	if req.Password != nil {
		data.Password = *req.Password
	}

	if req.Status != nil {
		data.Status = *req.Status
	}

	_, err = l.svcCtx.UserRpc.Update(l.ctx, &data)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.RoleRpc.BindRoleByUserId(l.ctx, &role.BindRoleByUserRequest{
		UserId: req.Id,
		RoleId: req.RoleId,
	})

	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{
		Id: uint64(req.Id),
	}, nil

}
