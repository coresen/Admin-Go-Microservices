package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"role-rpc/internal/svc"
	"role-rpc/pb/role"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *role.ListRequest) (*role.ListResponse, error) {
	// todo: add your logic here and delete this line

	return &role.ListResponse{}, nil
}
