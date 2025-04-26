package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"menu-rpc/internal/svc"
	"menu-rpc/pb/menu"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *menu.UpdateRequest) (*menu.UpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &menu.UpdateResponse{}, nil
}
