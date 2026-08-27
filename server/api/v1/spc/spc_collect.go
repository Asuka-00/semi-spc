package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	spcService "github.com/flipped-aurora/gin-vue-admin/server/service/spc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CollectApi struct{}

// CollectData
// @Tags      SpcCollect
// @Summary   SPC数据采集（支持幂等性）
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     X-Idempotency-Key  header    string                                              false  "幂等性键（可选）"
// @Param     data               body      spcService.CollectDataRequest                       true   "采集数据"
// @Success   200                {object}  response.Response{data=spcService.CollectDataResponse}  "采集成功"
// @Router    /spc/collect [post]
func (a *CollectApi) CollectData(c *gin.Context) {
	var req spcService.CollectDataRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从header读取幂等性键
	idempotencyKey := c.GetHeader("X-Idempotency-Key")
	if idempotencyKey != "" {
		req.IdempotencyKey = &idempotencyKey
	}

	result, err := collectService.CollectData(&req)
	if err != nil {
		global.GVA_LOG.Error("数据采集失败!", zap.Error(err))
		response.FailWithMessage("数据采集失败: "+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}
