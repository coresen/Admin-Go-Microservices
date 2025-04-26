package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	model "menu-model"
	"menu-rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	MenuModel model.MenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		MenuModel: model.NewMenuModel(conn, c.CacheRedis),
	}
}
