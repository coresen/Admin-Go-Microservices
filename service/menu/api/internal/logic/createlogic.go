package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zore/service/menu/api/internal/svc"
	"zore/service/menu/api/internal/types"
	"zore/service/menu/rpc/pb/menu"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.MenuRpc.Create(l.ctx, &menu.CreateRequest{
		ParentId:  req.ParentId,
		Icon:      req.Icon,
		Path:      req.Path,
		Method:    req.Method,
		Status:    req.Status,
		MenuName:  req.MenuName,
		SortOrder: req.SortOrder,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, err

}
