package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zore/service/menu/api/internal/config"
	"zore/service/menu/rpc/menuclient"
)

type ServiceContext struct {
	Config  config.Config
	MenuRpc menuclient.Menu
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		MenuRpc: menuclient.NewMenu(zrpc.MustNewClient(c.MenuRpc)),
	}
}
