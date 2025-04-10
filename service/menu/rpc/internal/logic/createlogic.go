package logic

import (
	"context"
	"database/sql"
	"time"
	"zore/service/menu/model"

	"zore/service/menu/rpc/internal/svc"
	"zore/service/menu/rpc/pb/menu"

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

func (l *CreateLogic) Create(in *menu.CreateRequest) (*menu.CreateResponse, error) {
	// todo: add your logic here and delete this line

	data := model.Menu{
		ParentId: sql.NullInt64{
			Int64: in.ParentId,
			Valid: true,
		},
		Method:     in.Method,
		Status:     int64(in.Status),
		SortOrder:  int64(in.SortOrder),
		MenuName:   in.MenuName,
		CreateTime: time.Now(),
		Icon:       sql.NullString{String: in.Icon, Valid: true},
		Path:       in.Path,
	}
	resp, err := l.svcCtx.MenuModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}
	id, _ := resp.LastInsertId()
	return &menu.CreateResponse{
		Id: id,
	}, nil
}
