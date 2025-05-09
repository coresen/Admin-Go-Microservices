//跨域代理前缀
// const API_PROXY_PREFIX='/api'
// const BASE_URL = process.env.NODE_ENV === 'production' ? process.env.VUE_APP_API_BASE_URL : API_PROXY_PREFIX

module.exports = {
  LOGIN: `/api/user/login`,

  USER_LIST: `/api/user/list`,
  USER_CREATE: `/api/user/creat`,

  ROUTES: `/api/role/routes`,
  GOODS: `/goods`,
  GOODS_COLUMNS: `/columns`,
}
