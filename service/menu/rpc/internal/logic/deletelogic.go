package logic

import (
	"context"

	"zore/service/menu/rpc/internal/svc"
	"zore/service/menu/rpc/pb/menu"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *menu.RemoveRequest) (*menu.RemoveResponse, error) {
	// todo: add your logic here and delete this line

	return &menu.RemoveResponse{}, nil
}
