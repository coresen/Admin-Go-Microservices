package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zore/common/constant"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 封装成功响应
func Success(data interface{}) Response {
	return Response{
		Code: constant.SUCCESS,
		Msg:  constant.GetMessage(constant.SUCCESS),
		Data: data,
	}
}

// 封装错误响应
func Error(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusOK, &Response{401, "鉴权失败", nil})
}
