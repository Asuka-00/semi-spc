package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChartApi struct{}

// CreateSpcChart
// @Tags      SpcChart
// @Summary   创建控制图
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcChart                   true  "控制图信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createChart [post]
func (a *ChartApi) CreateSpcChart(c *gin.Context) {
	var chart spc.SpcChart
	err := c.ShouldBindJSON(&chart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = chartService.CreateSpcChart(&chart)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(chart, c)
}

// DeleteSpcChart
// @Tags      SpcChart
// @Summary   删除控制图
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "控制图ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteChart [delete]
func (a *ChartApi) DeleteSpcChart(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = chartService.DeleteSpcChart(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcChart
// @Tags      SpcChart
// @Summary   更新控制图
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcChart                   true  "控制图信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateChart [put]
func (a *ChartApi) UpdateSpcChart(c *gin.Context) {
	var chart spc.SpcChart
	err := c.ShouldBindJSON(&chart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = chartService.UpdateSpcChart(&chart)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcChart
// @Tags      SpcChart
// @Summary   根据ID获取控制图
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "控制图ID"
// @Success   200   {object}  response.Response{data=spc.SpcChart,msg=string}  "获取成功"
// @Router    /spc/findChart [get]
func (a *ChartApi) FindSpcChart(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chart, err := chartService.GetSpcChart(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(chart, c)
}

// GetSpcChartList
// @Tags      SpcChart
// @Summary   分页获取控制图列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getChartList [get]
func (a *ChartApi) GetSpcChartList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := chartService.GetSpcChartList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
