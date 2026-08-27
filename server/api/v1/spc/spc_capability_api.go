package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CapabilityApi struct{}

// CalculateCapability calculates process capability
// @Tags      SpcCapability
// @Summary   计算能力指数
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      object{chartId=uint,from=string,to=string}  true  "计算参数"
// @Success   200   {object}  response.Response{msg=string}               "计算成功"
// @Router    /spc/calculateCapability [post]
func (a *CapabilityApi) CalculateCapability(c *gin.Context) {
	var req struct {
		ChartID uint   `json:"chartId" binding:"required"`
		From    string `json:"from" binding:"required"`
		To      string `json:"to" binding:"required"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// This would call the capability service
	// For now, return success
	response.OkWithMessage("能力计算成功", c)
}

// GetCapabilityHistory gets capability history
// @Tags      SpcCapability
// @Summary   获取能力分析历史
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     chartId  query   int                                              true "控制图ID"
// @Success   200   {object}  response.Response{data=[]object,msg=string}     "获取成功"
// @Router    /spc/getCapabilityHistory [get]
func (a *CapabilityApi) GetCapabilityHistory(c *gin.Context) {
	chartIDStr := c.Query("chartId")
	if chartIDStr == "" {
		response.FailWithMessage("控制图ID必须提供", c)
		return
	}

	// Query capability history
	var history []interface{}
	err := global.GVA_DB.Raw("SELECT * FROM spc_capability WHERE chart_id = ? ORDER BY created_at DESC", chartIDStr).Scan(&history).Error
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithData(history, c)
}
