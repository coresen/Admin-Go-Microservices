package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"menu-rpc/internal/svc"
	"menu-rpc/pb/menu"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *menu.DetailRequest) (*menu.DetailResponse, error) {
	// todo: add your logic here and delete this line

	return &menu.DetailResponse{}, nil
}
