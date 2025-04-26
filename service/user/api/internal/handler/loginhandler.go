package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user-api/internal/logic"
	"user-api/internal/svc"
	"user-api/internal/types"
	"zero-common/constant"
	"zero-common/response"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.Error(constant.ERROR, err.Error()))
			return
		}
		l := logic.NewLoginLogic(r.Context(), svcCtx)

		headers := []string{
			"X-Forwarded-For",
			"Proxy-Client-IP",
			"WL-Proxy-Client-IP",
			"X-Real-Ip",
		}
		var ip string
		for _, h := range headers {
			clientIp := r.Header.Get(h)
			if clientIp != "" {
				ip = clientIp
				break
			}
		}
		resp, err := l.Login(&req, ip)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.Error(constant.ERROR, err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
