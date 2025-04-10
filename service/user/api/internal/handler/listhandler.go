package handler

import (
	"net/http"
	"zore/common/constant"
	"zore/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zore/service/user/api/internal/logic"
	"zore/service/user/api/internal/svc"
	"zore/service/user/api/internal/types"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			httpx.OkJson(w, response.Error(constant.ERROR, err.Error()))
		} else {
			httpx.OkJson(w, response.Success(resp))
		}
	}
}
