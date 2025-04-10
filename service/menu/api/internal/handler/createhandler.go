package handler

import (
	"net/http"
	"zore/common/constant"
	"zore/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zore/service/menu/api/internal/logic"
	"zore/service/menu/api/internal/svc"
	"zore/service/menu/api/internal/types"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, response.Error(constant.ERROR, err.Error()))
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			httpx.OkJson(w, response.Error(constant.ERROR, err.Error()))
		} else {
			httpx.OkJson(w, response.Success(resp))
		}
	}
}
