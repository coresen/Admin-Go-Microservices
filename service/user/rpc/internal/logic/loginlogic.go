package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"zore/common/constant"
	"zore/common/crypt"
	"zore/service/user/model"
	"zore/service/user/rpc/internal/svc"
	"zore/service/user/rpc/pb/user"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	password := crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)

	if password != res.Password {
		return nil, errors.New("密码错误")
	}

	if res.Status == constant.DISABLE {
		return nil, errors.New(fmt.Sprintf("【%s】账号已被停用", in.Username))
	}

	if in.Ip != "" {
		res.Ip = in.Ip
	}

	res.LoginLast = sql.NullTime{Time: time.Now(), Valid: true}
	res.UpdateTime = time.Now()

	go func() {
		err = l.svcCtx.UserModel.Update(context.Background(), res)
	}()

	return &user.LoginResponse{
		Id:       int64(res.Id),
		Username: res.Username,
	}, nil

}
