package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	spcService "github.com/flipped-aurora/gin-vue-admin/server/service/spc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DashboardApi struct{}
type SampleApi struct{}

// GetDashboardOverview
// @Tags      SpcDashboard
// @Summary   获取SPC仪表盘总览
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     days  query     int  false  "统计天数"
// @Success   200   {object}  response.Response{data=map[string]interface{}}
// @Router    /spc/getDashboardOverview [get]
func (a *DashboardApi) GetDashboardOverview(c *gin.Context) {
	days := int(parseUintQuery(c, "days"))
	if days == 0 {
		days = 7
	}
	stats, err := dashboardService.GetDashboardOverview(days)
	if err != nil {
		global.GVA_LOG.Error("获取仪表盘失败!", zap.Error(err))
		response.FailWithMessage("获取仪表盘失败", c)
		return
	}
	response.OkWithData(stats, c)
}

// GetChartRuntime
// @Tags      SpcDashboard
// @Summary   获取控制图运行时数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     chartId  query     int  true  "控制图ID"
// @Param     limit    query     int  false "样本数量"
// @Success   200      {object}  response.Response{data=spcService.ChartRuntime}
// @Router    /spc/getChartRuntime [get]
func (a *DashboardApi) GetChartRuntime(c *gin.Context) {
	chartID := parseUintQuery(c, "chartId")
	if chartID == 0 {
		response.FailWithMessage("请选择控制图", c)
		return
	}
	limit := int(parseUintQuery(c, "limit"))
	data, err := dashboardService.GetChartRuntime(chartID, limit)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(data, c)
}

// GetSpcSampleList
// @Tags      SpcSample
// @Summary   分页获取样本列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     chartId  query  int  false "控制图ID"
// @Success   200      {object}  response.Response{data=response.PageResult}
// @Router    /spc/getSampleList [get]
func (a *SampleApi) GetSpcSampleList(c *gin.Context) {
	page := parsePage(c)
	chartID := parseUintQuery(c, "chartId")
	list, total, err := sampleService.GetSpcSampleList(page, chartID)
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// FindSpcSample
// @Tags      SpcSample
// @Summary   根据ID获取样本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response
// @Router    /spc/findSample [get]
func (a *SampleApi) FindSpcSample(c *gin.Context) {
	id, err := bindQueryID(c)
	if err != nil || id == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}
	item, err := sampleService.GetSpcSample(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}

// GetSpcMeasurementList
// @Tags      SpcSample
// @Summary   获取样本测量值
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     sampleId  query  int  true "样本ID"
// @Success   200       {object}  response.Response
// @Router    /spc/getMeasurementList [get]
func (a *SampleApi) GetSpcMeasurementList(c *gin.Context) {
	sampleID := parseUintQuery(c, "sampleId")
	if sampleID == 0 {
		response.FailWithMessage("样本ID不能为空", c)
		return
	}
	list, err := measurementService.GetSpcMeasurementList(sampleID)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(list, c)
}

var _ = spcService.ChartRuntime{}
