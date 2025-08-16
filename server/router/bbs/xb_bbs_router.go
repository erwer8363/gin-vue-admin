package bbs

import (
	"github.com/gin-gonic/gin"
)

type XkBbsRouter struct{}

func (e *XkBbsRouter) InitXkBbsRouter(Router *gin.RouterGroup) {

	xkBbsCustomerRouterWithoutRecord := Router.Group("bbs")
	//.Use(middleware.OperationRecord())
	{
		xkBbsCustomerRouterWithoutRecord.POST("save", xkBbsApi.CreateXkBbs) // 新增客户信息
		xkBbsCustomerRouterWithoutRecord.POST("update", xkBbsApi.UpdateXkBbs)
	}

	xkBbsRouterWithoutRecord := Router.Group("bbs")
	{
		xkBbsRouterWithoutRecord.GET("get", xkBbsApi.GetXkBbs)              // 获取单一客户信息
		xkBbsRouterWithoutRecord.GET("detail/:id", xkBbsApi.GetXkBbsDetail) // 获取bbs详情
	}
}
