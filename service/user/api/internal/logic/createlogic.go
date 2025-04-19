package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zore/service/role/rpc/pb/role"
	"zore/service/user/rpc/pb/user"

	"zore/service/user/api/internal/svc"
	"zore/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	fmt.Println("进入")

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println("err", err.Error())
		return nil, err
	}

	userID, _ := l.ctx.Value("userid").(int64)
	var newId int64
	fmt.Println("进行11")
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		res, err := l.svcCtx.UserRpc.Create(l.ctx, &user.CreateUserRequest{
			Username:  req.Username,
			Password:  req.Password,
			ParentId:  req.ParentId,
			Status:    req.Status,
			CreateUid: userID,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}

		newId = res.Id

		_, err = l.svcCtx.RoleRpc.BindRoleByUserId(l.ctx, &role.BindRoleByUserRequest{
			UserId: res.Id,
			RoleId: req.RoleId,
		})

		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	})

	fmt.Println("123234", err)
	return &types.CreateResponse{
		Id: newId,
	}, err
}
