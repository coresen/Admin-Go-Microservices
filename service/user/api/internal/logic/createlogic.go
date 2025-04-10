package logic

import (
	"context"
	"zore/service/role/rpc/pb/role"
	"zore/service/user/rpc/pb/user"

	"zore/service/user/api/internal/svc"
	"zore/service/user/api/internal/types"

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
	userID, _ := l.ctx.Value("userid").(int64)
	res, err := l.svcCtx.UserRpc.Create(l.ctx, &user.CreateUserRequest{
		Username:  req.Username,
		Password:  req.Password,
		ParentId:  req.ParentId,
		Status:    req.Status,
		CreateUid: userID,
	})

	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.RoleRpc.BindRoleByUserId(l.ctx, &role.BindRoleByUserRequest{
		UserId: res.Id,
		RoleId: req.RoleId,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
