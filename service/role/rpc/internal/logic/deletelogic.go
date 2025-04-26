package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"role-rpc/internal/svc"
	"role-rpc/pb/role"
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

func (l *DeleteLogic) Delete(in *role.DeleteRequest) (*role.DeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &role.DeleteResponse{}, nil
}
