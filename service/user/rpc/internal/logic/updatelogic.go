package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	model "user-model"
	"user-rpc/internal/svc"
	"user-rpc/pb/user"
	"zero-common/crypt"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	// todo: add your logic here and delete this line

	in.Password = crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)

	err := l.svcCtx.UserModel.Update(l.ctx, &model.User{
		Id:         in.Id,
		Password:   in.Password,
		Status:     in.Status,
		ParentId:   in.ParentId,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &user.UpdateUserResponse{
		Id: in.Id,
	}, nil
}
