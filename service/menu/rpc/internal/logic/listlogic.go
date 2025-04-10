package logic

import (
	"context"

	"zore/service/menu/rpc/internal/svc"
	"zore/service/menu/rpc/pb/menu"

	"github.com/zeromicro/go-zero/core/logx"
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
