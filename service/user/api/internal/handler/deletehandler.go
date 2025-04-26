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

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.Error(constant.ERROR, err.Error()))
			return
		}

		l := logic.NewDeleteLogic(r.Context(), svcCtx)
		resp, err := l.Delete(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.Error(constant.ERROR, err.Error()))
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Success(resp))
		}
	}
}
