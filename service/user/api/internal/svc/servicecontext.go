package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zore/common/interceptor"
	"zore/service/role/rpc/roleclient"
	"zore/service/user/api/internal/config"
	"zore/service/user/rpc/userclient"
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
