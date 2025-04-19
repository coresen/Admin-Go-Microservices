package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zore/service/role/model"
	"zore/service/role/rpc/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	RoleModel        model.RoleModel
	UserRoleModel    model.UserRoleModel
	PermissionsModel model.PermissionsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:           c,
		RoleModel:        model.NewRoleModel(conn, c.CacheRedis), //UserModel: model.NewUserModel(conn, c.CacheRedis),
		PermissionsModel: model.NewPermissionsModel(conn),
		UserRoleModel:    model.NewUserRoleModel(conn),
	}
}
