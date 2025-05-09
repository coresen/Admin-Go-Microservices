package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"user-rpc/internal/svc"
	"user-rpc/pb/user"
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

func (l *DeleteLogic) Delete(in *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteUserResponse{}, nil
}
