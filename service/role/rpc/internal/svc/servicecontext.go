package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	roleModel "role-model"
	"role-rpc/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	RoleModel        roleModel.RoleModel
	UserRoleModel    roleModel.UserRoleModel
	PermissionsModel roleModel.PermissionsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:           c,
		RoleModel:        roleModel.NewRoleModel(conn),
		PermissionsModel: roleModel.NewPermissionsModel(conn),
		UserRoleModel:    roleModel.NewUserRoleModel(conn),
	}
}
