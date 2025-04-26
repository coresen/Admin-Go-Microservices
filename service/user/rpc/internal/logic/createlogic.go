package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"time"
	model "user-model"
	"user-rpc/internal/svc"
	"user-rpc/pb/user"
	"zero-common/crypt"

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

func (l *CreateLogic) Create(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {

	password := crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)

	node, err := snowflake.NewNode(1) // 1 是节点ID，可以根据实际需要设置不同的值
	if err != nil {
		return nil, err
	}

	id := node.Generate().Int64()
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Id:         id,
		ParentId:   in.ParentId,
		Username:   in.Username,
		Status:     in.Status,
		Password:   password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		Id: id,
	}, nil
}
