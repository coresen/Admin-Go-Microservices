package logic

import (
	"context"
	"zore/service/role/rpc/pb/role"

	"zore/service/role/api/internal/svc"
	"zore/service/role/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line

	roleRes, err := l.svcCtx.RoleRpc.Create(l.ctx, &role.CreateRequest{
		ParentId:    req.ParentId,
		Status:      req.Status,
		RoleName:    req.RoleName,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	id := roleRes.Id

	if len(req.PermissionId) == 0 {
		return &types.CreateResponse{Id: id}, nil
	}

	data := make([]*role.Permission, len(req.PermissionId))
	for _, v := range req.PermissionId {
		data = append(data, &role.Permission{
			RoleId:       id,
			PermissionId: v,
		})
	}

	_, err = l.svcCtx.RoleRpc.CreatePermission(l.ctx, &role.PermissionCreateRequest{
		Data: data,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{Id: id}, nil
}
