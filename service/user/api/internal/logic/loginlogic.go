package logic

import (
	"context"
	"fmt"
	"user-api/internal/svc"
	"user-api/internal/types"
	"zero-common/dict"
	jwtx "zero-common/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest, ip string) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
		Ip:       ip,
	})
	if err != nil {
		return nil, err
	}

	userClaims := jwtx.UserClaims{
		Userid: res.Id,
	}

	if res.Username == dict.SYSTEM_USERNAME {
		userClaims.Role = dict.SUPER_ADMIN
	} else {
		userClaims.Role = dict.GENRAL_ADMIN

		roleRes, err := l.svcCtx.RoleRpc.GetRolesByUserIds(l.ctx, &role.GetRolesByUserIdsReq{
			GetRolesByUserIdsReq: []int64{res.Id},
		})

		if err != nil {
			return nil, err
		}

		if len(roleRes.Roles) > 0 {
			roleIds := make([]int64, 0, len(roleRes.Roles))
			for i, _ := range roleRes.Roles {
				roleIds = append(roleIds, i)
			}
			permissionRes, err := l.svcCtx.RoleRpc.GetPermissionByRole(l.ctx, &role.GetPermissionByRoleRequest{
				UserId:  res.Id,
				RoleIds: roleIds,
			})
			fmt.Println("permissionRes::", permissionRes.Data)
			if err == nil && len(permissionRes.Data) > 0 {
				userClaims.Menu = permissionRes.Data
			}
		}
	}

	accessToken, expireTime, err := jwtx.GetToken(userClaims, l.svcCtx.Config.Auth.AccessExpire, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: expireTime.UnixMilli(),
		Username:     res.Username,
	}, nil
}
