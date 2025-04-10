package logic

import (
	"context"

	"zore/service/menu/rpc/internal/svc"
	"zore/service/menu/rpc/pb/menu"

	"github.com/zeromicro/go-zero/core/logx"
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
