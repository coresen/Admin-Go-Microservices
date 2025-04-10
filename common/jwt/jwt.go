package jwtx

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	Userid int64       `json:"userid"` // 用户 ID
	Role   string      `json:"role"`
	Menu   interface{} `json:"menu"`
	Data   []int64     `json:"data"`
}

type MyCustomClaims struct {
	UserClaims
	jwt.RegisteredClaims
}

func GetToken(userClaims UserClaims, accessExpire int64, accessSecret string) (string, time.Time, error) {
	expireTime := time.Now().Add(time.Duration(accessExpire) * time.Second)
	claims := MyCustomClaims{
		UserClaims: UserClaims{
			Userid: userClaims.Userid, // 设置用户 ID
			Role:   userClaims.Role,
			Menu:   userClaims.Menu,
			Data:   userClaims.Data,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 过期时间7天
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(accessSecret))

	return tokenStr, expireTime, err

}
