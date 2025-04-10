package logic

import (
	"context"
	"zore/service/user/rpc/internal/svc"
	"zore/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *user.ListRequest) (*user.ListResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.FindPageList(l.ctx, in.Username, in.Status, 1, 10)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.UserModel.FindCount(l.ctx, in.Username, in.Status)
	if err != nil {
		return nil, err
	}

	data := make([]*user.UserItem, len(res))

	for i, item := range res {
		data[i] = &user.UserItem{
			Id:         int64(item.Id),
			ParentId:   item.ParentId,
			Username:   item.Username,
			Password:   item.Password,
			Status:     item.Status,
			Ip:         item.Ip,
			CreateTime: item.CreateTime.Unix(),
			LoginLast:  item.LoginLast.Time.Unix(),
		}
	}

	return &user.ListResponse{
		Data:  data,
		Count: count,
	}, nil
}
