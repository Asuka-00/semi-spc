package spc

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RuntimeApi struct{}

// GetChartRuntime gets chart runtime data for monitoring
// @Tags      SpcRuntime
// @Summary   获取控制图运行时数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     chartId  query   int    true   "控制图ID"
// @Param     from     query   string false  "开始时间"
// @Param     to       query   string false  "结束时间"
// @Param     page     query   int    false  "页码"
// @Param     pageSize query   int    false  "每页大小"
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取成功"
// @Router    /spc/getChartRuntime [get]
func (a *RuntimeApi) GetChartRuntime(c *gin.Context) {
	chartIDStr := c.Query("chartId")
	chartID, err := strconv.ParseUint(chartIDStr, 10, 32)
	if err != nil || chartID == 0 {
		response.FailWithMessage("控制图ID必须提供", c)
		return
	}

	fromStr := c.Query("from")
	toStr := c.Query("to")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 50
	}
	if pageSize > 100 {
		pageSize = 100
	}

	var fromTime, toTime *time.Time
	if fromStr != "" {
		t, err := time.Parse(time.RFC3339, fromStr)
		if err == nil {
			fromTime = &t
		}
	}
	if toStr != "" {
		t, err := time.Parse(time.RFC3339, toStr)
		if err == nil {
			toTime = &t
		}
	}

	// 获取控制图信息
	chart, err := chartService.GetSpcChart(uint(chartID))
	if err != nil {
		global.GVA_LOG.Error("获取控制图失败!", zap.Error(err))
		response.FailWithMessage("控制图不存在", c)
		return
	}

	// 获取当前控制限
	currentLimit, err := chartService.GetCurrentControlLimit(uint(chartID))
	if err != nil {
		global.GVA_LOG.Warn("获取控制限失败", zap.Error(err))
	}

	// 获取当前规格（通过SpecID）
	var currentSpec interface{}
	if chart.SpecID > 0 {
		var spec spc.SpcSpec
		err = global.GVA_DB.Preload("Parameter").Preload("Product").First(&spec, chart.SpecID).Error
		if err == nil {
			currentSpec = spec
		}
	}

	// 获取样本数据
	samples, total, err := sampleService.GetSpcSampleListWithTimeRange(uint(chartID), fromTime, toTime, page, pageSize)
	if err != nil {
		global.GVA_LOG.Error("获取样本数据失败!", zap.Error(err))
		response.FailWithMessage("获取样本数据失败", c)
		return
	}

	// 获取告警数据
	var alarms []spc.SpcAlarm
	alarmsData, err := alarmService.GetAlarmsByChartAndTimeRange(uint(chartID), fromTime, toTime)
	if err != nil {
		global.GVA_LOG.Warn("获取告警数据失败", zap.Error(err))
		alarms = []spc.SpcAlarm{}
	} else {
		alarms = alarmsData
	}

	result := map[string]interface{}{
		"chart":        chart,
		"currentLimit": currentLimit,
		"currentSpec":  currentSpec,
		"samples":      samples,
		"total":        total,
		"page":         page,
		"pageSize":     pageSize,
		"alarms":       alarms,
	}

	response.OkWithData(result, c)
}

// GetCapability calculates process capability
// @Tags      SpcCapability
// @Summary   计算能力指数
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     chartId  query   int    true   "控制图ID"
// @Param     from     query   string true   "开始时间"
// @Param     to       query   string true   "结束时间"
// @Success   200   {object}  response.Response{data=object,msg=string}  "计算成功"
// @Router    /spc/getCapability [get]
func (a *RuntimeApi) GetCapability(c *gin.Context) {
	chartIDStr := c.Query("chartId")
	chartID, err := strconv.ParseUint(chartIDStr, 10, 32)
	if err != nil || chartID == 0 {
		response.FailWithMessage("控制图ID必须提供", c)
		return
	}

	fromStr := c.Query("from")
	toStr := c.Query("to")
	if fromStr == "" || toStr == "" {
		response.FailWithMessage("时间范围必须提供", c)
		return
	}

	fromTime, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		response.FailWithMessage("开始时间格式错误", c)
		return
	}

	toTime, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		response.FailWithMessage("结束时间格式错误", c)
		return
	}

	result, err := capabilityService.CalculateAndSaveCapability(uint(chartID), fromTime, toTime)
	if err != nil {
		global.GVA_LOG.Error("能力计算失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// CalculateLimits recalculates control limits
// @Tags      SpcChart
// @Summary   重算控制限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      object{chartId=uint,window=int}  true  "控制图ID和窗口大小"
// @Success   200   {object}  response.Response{msg=string}    "计算成功"
// @Router    /spc/calculateLimits [post]
func (a *RuntimeApi) CalculateLimits(c *gin.Context) {
	var req struct {
		ChartID uint `json:"chartId" binding:"required"`
		Window  int  `json:"window" binding:"required"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Window < 10 || req.Window > 1000 {
		response.FailWithMessage("窗口大小必须在10-1000之间", c)
		return
	}

	result, err := chartService.RecalculateControlLimits(req.ChartID, req.Window)
	if err != nil {
		global.GVA_LOG.Error("控制限计算失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// CollectCsv handles CSV batch upload
// @Tags      SpcCollect
// @Summary   CSV批量上传
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file  true  "CSV文件"
// @Success   200   {object}  response.Response{data=object,msg=string}  "上传成功"
// @Router    /spc/collectCsv [post]
func (a *RuntimeApi) CollectCsv(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("文件上传失败", c)
		return
	}

	f, err := file.Open()
	if err != nil {
		response.FailWithMessage("文件打开失败", c)
		return
	}
	defer f.Close()

	result, err := collectService.CollectCsv(f)
	if err != nil {
		global.GVA_LOG.Error("CSV处理失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// GetDashboard gets dashboard statistics
// @Tags      SpcDashboard
// @Summary   获取Dashboard统计数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取成功"
// @Router    /spc/getDashboard [get]
func (a *RuntimeApi) GetDashboard(c *gin.Context) {
	// 今日告警数
	todayAlarms, err := alarmService.GetTodayAlarmCount()
	if err != nil {
		global.GVA_LOG.Warn("获取今日告警数失败", zap.Error(err))
		todayAlarms = 0
	}

	// 未处理告警数
	openAlarms, err := alarmService.GetOpenAlarmCount()
	if err != nil {
		global.GVA_LOG.Warn("获取未处理告警数失败", zap.Error(err))
		openAlarms = 0
	}

	// OOC率
	oocRate, err := alarmService.GetOOCRate(7) // 近7天
	if err != nil {
		global.GVA_LOG.Warn("获取OOC率失败", zap.Error(err))
		oocRate = 0
	}

	// 未关闭OCAP数
	openOcap, err := ocapExecutionService.GetOpenOcapCount()
	if err != nil {
		global.GVA_LOG.Warn("获取OCAP数失败", zap.Error(err))
		openOcap = 0
	}

	// TOP设备（按告警数排序）
	topEquipment, err := alarmService.GetTopEquipmentByAlarms(5, 7) // 近7天TOP 5
	if err != nil {
		global.GVA_LOG.Warn("获取TOP设备失败", zap.Error(err))
		topEquipment = make([]map[string]interface{}, 0)
	}

	// TOP参数（按告警数排序）
	topParameter, err := alarmService.GetTopParameterByAlarms(5, 7) // 近7天TOP 5
	if err != nil {
		global.GVA_LOG.Warn("获取TOP参数失败", zap.Error(err))
		topParameter = make([]map[string]interface{}, 0)
	}

	result := map[string]interface{}{
		"todayAlarms":  todayAlarms,
		"openAlarms":   openAlarms,
		"oocRate":      oocRate,
		"openOcap":     openOcap,
		"topEquipment": topEquipment,
		"topParameter": topParameter,
	}

	response.OkWithData(result, c)
}
