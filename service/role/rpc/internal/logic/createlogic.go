package logic

import (
	"context"

	"zore/service/role/rpc/internal/svc"
	"zore/service/role/rpc/pb/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *role.CreateRequest) (*role.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &role.CreateResponse{}, nil
}
