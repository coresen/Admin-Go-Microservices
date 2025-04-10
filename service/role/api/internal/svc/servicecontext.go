package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zore/service/role/api/internal/config"
	"zore/service/role/rpc/roleclient"
)

type ServiceContext struct {
	Config  config.Config
	RoleRpc roleclient.Role
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		RoleRpc: roleclient.NewRole(zrpc.MustNewClient(c.RoleRpc)),
	}
}
