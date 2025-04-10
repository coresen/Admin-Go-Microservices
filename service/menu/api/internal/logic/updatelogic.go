package logic

import (
	"context"
	"zore/service/menu/rpc/pb/menu"

	"zore/service/menu/api/internal/svc"
	"zore/service/menu/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest, id int64) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.MenuRpc.Update(l.ctx, &menu.UpdateRequest{
		Method:    req.Method,
		Path:      req.Path,
		Id:        id,
		Status:    req.Status,
		SortOrder: req.SortOrder,
		MenuName:  req.MenuName,
		ParentId:  req.ParentId,
		Icon:      req.Icon,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{
		Method:    req.Method,
		Path:      req.Path,
		Id:        id,
		Status:    req.Status,
		SortOrder: req.SortOrder,
		MenuName:  req.MenuName,
		ParentId:  req.ParentId,
		Icon:      req.Icon,
	}, nil
}
