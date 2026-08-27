package spc

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ParameterApi struct{}

// CreateSpcParameter creates a parameter
// @Tags      SpcParameter
// @Summary   创建参数
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcParameter               true  "参数信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createParameter [post]
func (a *ParameterApi) CreateSpcParameter(c *gin.Context) {
	var param spc.SpcParameter
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = parameterService.CreateSpcParameter(&param)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(param, c)
}

// DeleteSpcParameter deletes a parameter
// @Tags      SpcParameter
// @Summary   删除参数
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "参数ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteParameter [delete]
func (a *ParameterApi) DeleteSpcParameter(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = parameterService.DeleteSpcParameter(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcParameter updates a parameter
// @Tags      SpcParameter
// @Summary   更新参数
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcParameter               true  "参数信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateParameter [put]
func (a *ParameterApi) UpdateSpcParameter(c *gin.Context) {
	var param spc.SpcParameter
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = parameterService.UpdateSpcParameter(&param)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcParameter finds a parameter by ID
// @Tags      SpcParameter
// @Summary   根据ID获取参数
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "参数ID"
// @Success   200   {object}  response.Response{data=spc.SpcParameter,msg=string}  "获取成功"
// @Router    /spc/findParameter [get]
func (a *ParameterApi) FindSpcParameter(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	param, err := parameterService.GetSpcParameter(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(param, c)
}

// GetSpcParameterList gets parameter list with pagination
// @Tags      SpcParameter
// @Summary   分页获取参数列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getParameterList [get]
func (a *ParameterApi) GetSpcParameterList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := parameterService.GetSpcParameterList(pageInfo)
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

type SpecApi struct{}

// CreateSpcSpec creates a spec
// @Tags      SpcSpec
// @Summary   创建规格
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcSpec                    true  "规格信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createSpec [post]
func (a *SpecApi) CreateSpcSpec(c *gin.Context) {
	var spec spc.SpcSpec
	err := c.ShouldBindJSON(&spec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = specService.CreateSpcSpec(&spec)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(spec, c)
}

// DeleteSpcSpec deletes a spec
// @Tags      SpcSpec
// @Summary   删除规格
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "规格ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteSpec [delete]
func (a *SpecApi) DeleteSpcSpec(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = specService.DeleteSpcSpec(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcSpec updates a spec
// @Tags      SpcSpec
// @Summary   更新规格
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcSpec                    true  "规格信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateSpec [put]
func (a *SpecApi) UpdateSpcSpec(c *gin.Context) {
	var spec spc.SpcSpec
	err := c.ShouldBindJSON(&spec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = specService.UpdateSpcSpec(&spec)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcSpec finds a spec by ID
// @Tags      SpcSpec
// @Summary   根据ID获取规格
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "规格ID"
// @Success   200   {object}  response.Response{data=spc.SpcSpec,msg=string}  "获取成功"
// @Router    /spc/findSpec [get]
func (a *SpecApi) FindSpcSpec(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	spec, err := specService.GetSpcSpec(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(spec, c)
}

// GetSpcSpecList gets spec list with pagination
// @Tags      SpcSpec
// @Summary   分页获取规格列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     parameterId  query   int                                           false "参数ID"
// @Param     productId  query   int                                             false "产品ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getSpecList [get]
func (a *SpecApi) GetSpcSpecList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	parameterIDStr := c.Query("parameterId")
	productIDStr := c.Query("productId")
	parameterID, _ := strconv.ParseUint(parameterIDStr, 10, 32)
	productID, _ := strconv.ParseUint(productIDStr, 10, 32)

	list, total, err := specService.GetSpcSpecList(pageInfo, uint(parameterID), uint(productID))
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

type ControlLimitApi struct{}

// CreateSpcControlLimit creates a control limit
// @Tags      SpcControlLimit
// @Summary   创建控制限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcControlLimit            true  "控制限信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createControlLimit [post]
func (a *ControlLimitApi) CreateSpcControlLimit(c *gin.Context) {
	var limit spc.SpcControlLimit
	err := c.ShouldBindJSON(&limit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = controlLimitService.CreateSpcControlLimit(&limit)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(limit, c)
}

// DeleteSpcControlLimit deletes a control limit
// @Tags      SpcControlLimit
// @Summary   删除控制限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "控制限ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteControlLimit [delete]
func (a *ControlLimitApi) DeleteSpcControlLimit(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = controlLimitService.DeleteSpcControlLimit(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcControlLimit updates a control limit
// @Tags      SpcControlLimit
// @Summary   更新控制限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcControlLimit            true  "控制限信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateControlLimit [put]
func (a *ControlLimitApi) UpdateSpcControlLimit(c *gin.Context) {
	var limit spc.SpcControlLimit
	err := c.ShouldBindJSON(&limit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = controlLimitService.UpdateSpcControlLimit(&limit)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcControlLimit finds a control limit by ID
// @Tags      SpcControlLimit
// @Summary   根据ID获取控制限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "控制限ID"
// @Success   200   {object}  response.Response{data=spc.SpcControlLimit,msg=string}  "获取成功"
// @Router    /spc/findControlLimit [get]
func (a *ControlLimitApi) FindSpcControlLimit(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	limit, err := controlLimitService.GetSpcControlLimit(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(limit, c)
}

// GetSpcControlLimitList gets control limit list with pagination
// @Tags      SpcControlLimit
// @Summary   分页获取控制限列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     chartId  query   int                                              false "控制图ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getControlLimitList [get]
func (a *ControlLimitApi) GetSpcControlLimitList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chartIDStr := c.Query("chartId")
	chartID, _ := strconv.ParseUint(chartIDStr, 10, 32)

	list, total, err := controlLimitService.GetSpcControlLimitList(pageInfo, uint(chartID))
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

type RuleApi struct{}

// CreateSpcRule creates a rule
// @Tags      SpcRule
// @Summary   创建规则
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcRule                    true  "规则信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createRule [post]
func (a *RuleApi) CreateSpcRule(c *gin.Context) {
	var rule spc.SpcRule
	err := c.ShouldBindJSON(&rule)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ruleService.CreateSpcRule(&rule)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(rule, c)
}

// DeleteSpcRule deletes a rule
// @Tags      SpcRule
// @Summary   删除规则
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "规则ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteRule [delete]
func (a *RuleApi) DeleteSpcRule(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ruleService.DeleteSpcRule(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcRule updates a rule
// @Tags      SpcRule
// @Summary   更新规则
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcRule                    true  "规则信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateRule [put]
func (a *RuleApi) UpdateSpcRule(c *gin.Context) {
	var rule spc.SpcRule
	err := c.ShouldBindJSON(&rule)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ruleService.UpdateSpcRule(&rule)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcRule finds a rule by ID
// @Tags      SpcRule
// @Summary   根据ID获取规则
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "规则ID"
// @Success   200   {object}  response.Response{data=spc.SpcRule,msg=string}  "获取成功"
// @Router    /spc/findRule [get]
func (a *RuleApi) FindSpcRule(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	rule, err := ruleService.GetSpcRule(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(rule, c)
}

// GetSpcRuleList gets rule list with pagination
// @Tags      SpcRule
// @Summary   分页获取规则列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     chartId  query   int                                              false "控制图ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getRuleList [get]
func (a *RuleApi) GetSpcRuleList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chartIDStr := c.Query("chartId")
	chartID, _ := strconv.ParseUint(chartIDStr, 10, 32)

	list, total, err := ruleService.GetSpcRuleList(pageInfo, uint(chartID))
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

type SampleApi struct{}

// FindSpcSample finds a sample by ID
// @Tags      SpcSample
// @Summary   根据ID获取样本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "样本ID"
// @Success   200   {object}  response.Response{data=spc.SpcSample,msg=string}  "获取成功"
// @Router    /spc/findSample [get]
func (a *SampleApi) FindSpcSample(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	sample, err := sampleService.GetSpcSample(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(sample, c)
}

// GetSpcSampleList gets sample list with pagination
// @Tags      SpcSample
// @Summary   分页获取样本列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     chartId  query   int                                              false "控制图ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getSampleList [get]
func (a *SampleApi) GetSpcSampleList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chartIDStr := c.Query("chartId")
	chartID, _ := strconv.ParseUint(chartIDStr, 10, 32)

	list, total, err := sampleService.GetSpcSampleList(pageInfo, uint(chartID))
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

type MeasurementApi struct{}

// GetSpcMeasurementList gets measurement list by sample ID
// @Tags      SpcMeasurement
// @Summary   获取测量值列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     sampleId  query   int                                              true "样本ID"
// @Success   200   {object}  response.Response{data=[]spc.SpcMeasurement,msg=string}  "获取成功"
// @Router    /spc/getMeasurementList [get]
func (a *MeasurementApi) GetSpcMeasurementList(c *gin.Context) {
	sampleIDStr := c.Query("sampleId")
	sampleID, err := strconv.ParseUint(sampleIDStr, 10, 32)
	if err != nil || sampleID == 0 {
		response.FailWithMessage("样本ID必须提供", c)
		return
	}

	list, err := measurementService.GetSpcMeasurementList(uint(sampleID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(list, c)
}

type OcapApi struct{}

// CreateSpcOcap creates an OCAP
// @Tags      SpcOcap
// @Summary   创建OCAP
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcOcap                    true  "OCAP信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createOcap [post]
func (a *OcapApi) CreateSpcOcap(c *gin.Context) {
	var ocap spc.SpcOcap
	err := c.ShouldBindJSON(&ocap)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ocapService.CreateSpcOcap(&ocap)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(ocap, c)
}

// DeleteSpcOcap deletes an OCAP
// @Tags      SpcOcap
// @Summary   删除OCAP
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "OCAP ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteOcap [delete]
func (a *OcapApi) DeleteSpcOcap(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ocapService.DeleteSpcOcap(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcOcap updates an OCAP
// @Tags      SpcOcap
// @Summary   更新OCAP
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcOcap                    true  "OCAP信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateOcap [put]
func (a *OcapApi) UpdateSpcOcap(c *gin.Context) {
	var ocap spc.SpcOcap
	err := c.ShouldBindJSON(&ocap)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ocapService.UpdateSpcOcap(&ocap)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcOcap finds an OCAP by ID
// @Tags      SpcOcap
// @Summary   根据ID获取OCAP
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "OCAP ID"
// @Success   200   {object}  response.Response{data=spc.SpcOcap,msg=string}  "获取成功"
// @Router    /spc/findOcap [get]
func (a *OcapApi) FindSpcOcap(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ocap, err := ocapService.GetSpcOcap(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(ocap, c)
}

// GetSpcOcapList gets OCAP list with pagination
// @Tags      SpcOcap
// @Summary   分页获取OCAP列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     chartId  query   int                                              false "控制图ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getOcapList [get]
func (a *OcapApi) GetSpcOcapList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chartIDStr := c.Query("chartId")
	chartID, _ := strconv.ParseUint(chartIDStr, 10, 32)

	list, total, err := ocapService.GetSpcOcapList(pageInfo, uint(chartID))
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

// StartOcap starts OCAP execution from an alarm
// @Tags      SpcOcap
// @Summary   启动OCAP执行
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      object{alarmId=uint,ocapId=uint}  true  "告警和OCAP ID"
// @Success   200   {object}  response.Response{msg=string}     "启动成功"
// @Router    /spc/startOcap [post]
func (a *OcapApi) StartOcap(c *gin.Context) {
	var req struct {
		AlarmID uint `json:"alarmId" binding:"required"`
		OcapID  uint `json:"ocapId" binding:"required"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ocapExecutionService.CreateSpcOcapExecution(&spc.SpcOcapExecution{
		AlarmID: req.AlarmID,
		OcapID:  req.OcapID,
		Status:  "OPEN",
	})
	if err != nil {
		global.GVA_LOG.Error("启动OCAP失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("OCAP启动成功", c)
}

type OcapExecutionApi struct{}

// UpdateSpcOcapExecution updates OCAP execution
// @Tags      SpcOcapExecution
// @Summary   更新OCAP执行
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcOcapExecution           true  "OCAP执行信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateOcapExecution [put]
func (a *OcapExecutionApi) UpdateSpcOcapExecution(c *gin.Context) {
	var exec spc.SpcOcapExecution
	err := c.ShouldBindJSON(&exec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = ocapExecutionService.UpdateSpcOcapExecution(&exec)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetSpcOcapExecutionList gets OCAP execution list
// @Tags      SpcOcapExecution
// @Summary   分页获取OCAP执行列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     alarmId  query   int                                              false "告警ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getOcapExecutionList [get]
func (a *OcapExecutionApi) GetSpcOcapExecutionList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	alarmIDStr := c.Query("alarmId")
	alarmID, _ := strconv.ParseUint(alarmIDStr, 10, 32)

	list, total, err := ocapExecutionService.GetSpcOcapExecutionList(pageInfo, uint(alarmID))
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
