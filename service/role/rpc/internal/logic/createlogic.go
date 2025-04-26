package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	model "role-model"
	"role-rpc/internal/svc"
	"role-rpc/pb/role"
	"time"
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

	node, err := snowflake.NewNode(1) // 1 是节点ID，可以根据实际需要设置不同的值
	if err != nil {
		return nil, err
	}

	id := node.Generate().Int64()
	_, err = l.svcCtx.RoleModel.Insert(l.ctx, &model.Role{
		Id:          id,
		ParentId:    in.ParentId,
		RoleName:    in.RoleName,
		Description: in.Description,
		CreatedAt:   time.Now(),
		Status:      in.Status,
	})

	if err != nil {
		return nil, err
	}

	return &role.CreateResponse{
		Id: id,
	}, nil
}
