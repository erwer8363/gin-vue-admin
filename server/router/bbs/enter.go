package bbs

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	XkBbsRouter
}

var (
	xkBbsApi = api.ApiGroupApp.XkBbsApiGroup.XkBbsApi
)
