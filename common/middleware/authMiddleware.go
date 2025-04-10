package middleware

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
	"zore/common/constant"
	"zore/common/dict"
	"zore/common/response"
	"zore/service/role/rpc/pb/role"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		roleTag := r.Context().Value("role").(string)
		if roleTag == dict.SUPER_ADMIN {
			next(w, r)
			return
		}

		menuTag := r.Context().Value("menu")
		if menuTag == nil {
			httpx.OkJson(w, response.Error(constant.ERROR_AUTH, "无权限访问！"))
			return
		}

		jsonData, _ := jsoniter.Marshal(menuTag)
		var menus []role.PermissionItem
		err := jsoniter.Unmarshal(jsonData, &menus)
		fmt.Println(menus)
		if err != nil {
			httpx.OkJson(w, response.Error(constant.ERROR_AUTH, err.Error()))
			return
		}

		for i, _ := range menus {
			if isPathMatch(menus[i].Path, r.URL.Path) && menus[i].Method == r.Method {
				next(w, r)
				return
			}
		}

		httpx.OkJson(w, response.Error(constant.ERROR_AUTH, "无权限访问！"))
		return

	}
}

func isPathMatch(permissionPath, requestPath string) bool {
	permParts := strings.Split(permissionPath, "/")
	reqParts := strings.Split(requestPath, "/")

	if len(permParts) != len(reqParts) {
		return false
	}

	for i := 0; i < len(permParts); i++ {
		if strings.HasPrefix(permParts[i], ":") {
			continue // 动态参数部分跳过检查
		}
		if permParts[i] != reqParts[i] {
			return false
		}
	}
	return true
}
