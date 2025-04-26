package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"user-rpc/internal/svc"
	"user-rpc/pb/user"
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

func (l *DetailLogic) Detail(in *user.DetailRequest) (*user.DetailResponse, error) {
	// todo: add your logic here and delete this line

	return &user.DetailResponse{}, nil
}
