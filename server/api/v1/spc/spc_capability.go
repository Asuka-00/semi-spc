package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	spcService "github.com/flipped-aurora/gin-vue-admin/server/service/spc"
	"github.com/gin-gonic/gin"
)

type CapabilityApi struct{}

// CalculateCapability
// @Tags      SpcCapability
// @Summary   计算过程能力
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spcService.CalculateCapabilityRequest  true  "计算参数"
// @Success   200   {object}  response.Response
// @Router    /spc/calculateCapability [post]
func (a *CapabilityApi) CalculateCapability(c *gin.Context) {
	var req spcService.CalculateCapabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	item, err := capabilityService.CalculateCapability(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(item, c)
}

// GetCapabilityList
// @Tags      SpcCapability
// @Summary   分页获取能力分析记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     chartId  query  int  false "控制图ID"
// @Success   200      {object}  response.Response{data=response.PageResult}
// @Router    /spc/getCapabilityList [get]
func (a *CapabilityApi) GetCapabilityList(c *gin.Context) {
	page := parsePage(c)
	chartID := parseUintQuery(c, "chartId")
	list, total, err := capabilityService.GetCapabilityList(page, chartID)
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// GetCapabilityHistory
// @Tags      SpcCapability
// @Summary   获取能力分析历史
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     chartId  query  int  false "控制图ID"
// @Param     limit    query  int  false "条数"
// @Success   200      {object}  response.Response
// @Router    /spc/getCapabilityHistory [get]
func (a *CapabilityApi) GetCapabilityHistory(c *gin.Context) {
	chartID := parseUintQuery(c, "chartId")
	limit := int(parseUintQuery(c, "limit"))
	list, err := capabilityService.GetCapabilityHistory(chartID, limit)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(list, c)
}
