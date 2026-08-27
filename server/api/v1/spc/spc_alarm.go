package spc

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AlarmApi struct{}

// GetSpcAlarmList
// @Tags      SpcAlarm
// @Summary   分页获取告警列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     status  query   string                                           false "告警状态"
// @Param     alarmType query string                                           false "告警类型"
// @Success   200   {object}  response.Response{data=response.PageResult}  "获取成功"
// @Router    /spc/getAlarmList [get]
func (a *AlarmApi) GetSpcAlarmList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	status := c.Query("status")
	alarmType := c.Query("alarmType")

	list, total, err := alarmService.GetSpcAlarmList(pageInfo, status, alarmType)
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

// AcknowledgeAlarm
// @Tags      SpcAlarm
// @Summary   确认告警
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "告警ID"
// @Success   200   {object}  response.Response{msg=string}  "确认成功"
// @Router    /spc/acknowledgeAlarm [post]
func (a *AlarmApi) AcknowledgeAlarm(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	remark := c.Query("remark")
	err = alarmService.AcknowledgeAlarm(uint(req.ID), remark)
	if err != nil {
		global.GVA_LOG.Error("确认失败!", zap.Error(err))
		response.FailWithMessage("确认失败", c)
		return
	}
	response.OkWithMessage("确认成功", c)
}

// CloseAlarm
// @Tags      SpcAlarm
// @Summary   关闭告警
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "告警ID"
// @Success   200   {object}  response.Response{msg=string}  "关闭成功"
// @Router    /spc/closeAlarm [post]
func (a *AlarmApi) CloseAlarm(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	remark := c.Query("remark")
	err = alarmService.CloseAlarm(uint(uint(req.ID)), remark)
	if err != nil {
		global.GVA_LOG.Error("关闭失败!", zap.Error(err))
		response.FailWithMessage("关闭失败", c)
		return
	}
	response.OkWithMessage("关闭成功", c)
}

// GetAlarmStatistics
// @Tags      SpcAlarm
// @Summary   获取告警统计
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     days  query     int                                          false "统计天数"
// @Success   200   {object}  response.Response{data=map[string]interface{}}  "获取成功"
// @Router    /spc/getAlarmStatistics [get]
func (a *AlarmApi) GetAlarmStatistics(c *gin.Context) {
	daysStr := c.DefaultQuery("days", "7")
	days, _ := strconv.Atoi(daysStr)

	stats, err := alarmService.GetAlarmStatistics(days)
	if err != nil {
		global.GVA_LOG.Error("获取统计失败!", zap.Error(err))
		response.FailWithMessage("获取统计失败", c)
		return
	}

	response.OkWithData(stats, c)
}
