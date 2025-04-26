package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"role-rpc/roleclient"
	"user-api/internal/config"
	"user-rpc/userclient"
	"zero-common/interceptor"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userclient.User
	RoleRpc roleclient.Role
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		RoleRpc: roleclient.NewRole(zrpc.MustNewClient(c.RoleRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientAuthInterceptor))),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientAuthInterceptor))),
	}
}
