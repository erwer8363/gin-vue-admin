package bbs

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type XkBbsApi struct {
}

func (e XkBbsApi) CreateXkBbs(c *gin.Context) {
	var xkBbs bbs.XkBbs
	err := c.ShouldBindJSON(&xkBbs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = xkBbsService.CreatXkBbs(&xkBbs)
	if err != nil {
		global.GVA_LOG.Error("保存失败!", zap.Error(err))
		response.FailWithMessage("保存失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

func (e XkBbsApi) UpdateXkBbs(c *gin.Context) {
	var xkBbs bbs.XkBbs
	err := c.ShouldBindJSON(&xkBbs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = xkBbsService.UpdateXkBbs(&xkBbs)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetExaCustomer
// @Tags      ExaCustomer
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaCustomer                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=exampleRes.ExaCustomerResponse,msg=string}  "获取单一客户信息,返回包括客户详情"
// @Router    /bbs/get?id=1  [get]
func (e *XkBbsApi) GetXkBbs(c *gin.Context) {
	var xkBbs bbs.XkBbs
	err := c.ShouldBindQuery(&xkBbs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//bbsService := service.ServiceGroupApp.XkBBsServiceGroup
	data, err := xkBbsService.GetXkBbs(xkBbs.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (e *XkBbsApi) GetXkBbsDetail(c *gin.Context) {
	id := c.Param("id")
	bbsId, _ := strconv.ParseUint(id, 10, 64)
	data, err := xkBbsService.GetXkBbs(uint(bbsId))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
