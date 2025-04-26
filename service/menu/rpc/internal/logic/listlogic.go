package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"menu-rpc/internal/svc"
	"menu-rpc/pb/menu"
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

func (l *ListLogic) List(in *menu.ListRequest) (*menu.ListResponse, error) {
	// todo: add your logic here and delete this line

	return &menu.ListResponse{}, nil
}
