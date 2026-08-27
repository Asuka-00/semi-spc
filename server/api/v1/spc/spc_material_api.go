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

type ProcessStepApi struct{}

// CreateSpcProcessStep creates a process step
// @Tags      SpcProcessStep
// @Summary   创建工艺步骤
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcProcessStep             true  "工艺步骤信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createProcessStep [post]
func (a *ProcessStepApi) CreateSpcProcessStep(c *gin.Context) {
	var step spc.SpcProcessStep
	err := c.ShouldBindJSON(&step)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = processStepService.CreateSpcProcessStep(&step)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(step, c)
}

// DeleteSpcProcessStep deletes a process step
// @Tags      SpcProcessStep
// @Summary   删除工艺步骤
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "工艺步骤ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteProcessStep [delete]
func (a *ProcessStepApi) DeleteSpcProcessStep(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = processStepService.DeleteSpcProcessStep(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcProcessStep updates a process step
// @Tags      SpcProcessStep
// @Summary   更新工艺步骤
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcProcessStep             true  "工艺步骤信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateProcessStep [put]
func (a *ProcessStepApi) UpdateSpcProcessStep(c *gin.Context) {
	var step spc.SpcProcessStep
	err := c.ShouldBindJSON(&step)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = processStepService.UpdateSpcProcessStep(&step)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcProcessStep finds a process step by ID
// @Tags      SpcProcessStep
// @Summary   根据ID获取工艺步骤
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                                    true  "工艺步骤ID"
// @Success   200   {object}  response.Response{data=spc.SpcProcessStep,msg=string}  "获取成功"
// @Router    /spc/findProcessStep [get]
func (a *ProcessStepApi) FindSpcProcessStep(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	step, err := processStepService.GetSpcProcessStep(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(step, c)
}

// GetSpcProcessStepList gets process step list with pagination
// @Tags      SpcProcessStep
// @Summary   分页获取工艺步骤列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getProcessStepList [get]
func (a *ProcessStepApi) GetSpcProcessStepList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := processStepService.GetSpcProcessStepList(pageInfo)
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

type RecipeApi struct{}

// CreateSpcRecipe creates a recipe
// @Tags      SpcRecipe
// @Summary   创建工艺配方
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcRecipe                  true  "配方信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createRecipe [post]
func (a *RecipeApi) CreateSpcRecipe(c *gin.Context) {
	var recipe spc.SpcRecipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = recipeService.CreateSpcRecipe(&recipe)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(recipe, c)
}

// DeleteSpcRecipe deletes a recipe
// @Tags      SpcRecipe
// @Summary   删除工艺配方
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "配方ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteRecipe [delete]
func (a *RecipeApi) DeleteSpcRecipe(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = recipeService.DeleteSpcRecipe(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcRecipe updates a recipe
// @Tags      SpcRecipe
// @Summary   更新工艺配方
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcRecipe                  true  "配方信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateRecipe [put]
func (a *RecipeApi) UpdateSpcRecipe(c *gin.Context) {
	var recipe spc.SpcRecipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = recipeService.UpdateSpcRecipe(&recipe)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcRecipe finds a recipe by ID
// @Tags      SpcRecipe
// @Summary   根据ID获取工艺配方
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "配方ID"
// @Success   200   {object}  response.Response{data=spc.SpcRecipe,msg=string}  "获取成功"
// @Router    /spc/findRecipe [get]
func (a *RecipeApi) FindSpcRecipe(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	recipe, err := recipeService.GetSpcRecipe(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(recipe, c)
}

// GetSpcRecipeList gets recipe list with pagination
// @Tags      SpcRecipe
// @Summary   分页获取工艺配方列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     equipmentId  query   int                                           false "设备ID"
// @Param     stepId  query   int                                                false "工艺步骤ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getRecipeList [get]
func (a *RecipeApi) GetSpcRecipeList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	equipmentIDStr := c.Query("equipmentId")
	stepIDStr := c.Query("stepId")
	equipmentID, _ := strconv.ParseUint(equipmentIDStr, 10, 32)
	stepID, _ := strconv.ParseUint(stepIDStr, 10, 32)

	list, total, err := recipeService.GetSpcRecipeList(pageInfo, uint(equipmentID), uint(stepID))
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

type LotApi struct{}

// CreateSpcLot creates a lot
// @Tags      SpcLot
// @Summary   创建批次
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcLot                     true  "批次信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createLot [post]
func (a *LotApi) CreateSpcLot(c *gin.Context) {
	var lot spc.SpcLot
	err := c.ShouldBindJSON(&lot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = lotService.CreateSpcLot(&lot)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(lot, c)
}

// DeleteSpcLot deletes a lot
// @Tags      SpcLot
// @Summary   删除批次
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "批次ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteLot [delete]
func (a *LotApi) DeleteSpcLot(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = lotService.DeleteSpcLot(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcLot updates a lot
// @Tags      SpcLot
// @Summary   更新批次
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcLot                     true  "批次信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateLot [put]
func (a *LotApi) UpdateSpcLot(c *gin.Context) {
	var lot spc.SpcLot
	err := c.ShouldBindJSON(&lot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = lotService.UpdateSpcLot(&lot)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcLot finds a lot by ID
// @Tags      SpcLot
// @Summary   根据ID获取批次
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "批次ID"
// @Success   200   {object}  response.Response{data=spc.SpcLot,msg=string}  "获取成功"
// @Router    /spc/findLot [get]
func (a *LotApi) FindSpcLot(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	lot, err := lotService.GetSpcLot(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(lot, c)
}

// GetSpcLotList gets lot list with pagination
// @Tags      SpcLot
// @Summary   分页获取批次列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     siteId  query   int                                                false "厂区ID"
// @Param     productId  query   int                                             false "产品ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getLotList [get]
func (a *LotApi) GetSpcLotList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	siteIDStr := c.Query("siteId")
	productIDStr := c.Query("productId")
	siteID, _ := strconv.ParseUint(siteIDStr, 10, 32)
	productID, _ := strconv.ParseUint(productIDStr, 10, 32)

	list, total, err := lotService.GetSpcLotList(pageInfo, uint(siteID), uint(productID))
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

// HoldSpcLot holds a lot
// @Tags      SpcLot
// @Summary   Hold批次
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "批次ID"
// @Success   200   {object}  response.Response{msg=string}  "Hold成功"
// @Router    /spc/holdLot [post]
func (a *LotApi) HoldSpcLot(c *gin.Context) {
	var req struct {
		ID      uint   `json:"ID" binding:"required"`
		Comment string `json:"comment"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = lotService.HoldSpcLot(req.ID, req.Comment)
	if err != nil {
		global.GVA_LOG.Error("Hold失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Hold成功", c)
}

// ReleaseSpcLot releases a lot
// @Tags      SpcLot
// @Summary   Release批次
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "批次ID"
// @Success   200   {object}  response.Response{msg=string}  "Release成功"
// @Router    /spc/releaseLot [post]
func (a *LotApi) ReleaseSpcLot(c *gin.Context) {
	var req struct {
		ID      uint   `json:"ID" binding:"required"`
		Comment string `json:"comment"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = lotService.ReleaseSpcLot(req.ID, req.Comment)
	if err != nil {
		global.GVA_LOG.Error("Release失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Release成功", c)
}

type WaferApi struct{}

// CreateSpcWafer creates a wafer
// @Tags      SpcWafer
// @Summary   创建晶圆
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcWafer                   true  "晶圆信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createWafer [post]
func (a *WaferApi) CreateSpcWafer(c *gin.Context) {
	var wafer spc.SpcWafer
	err := c.ShouldBindJSON(&wafer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = waferService.CreateSpcWafer(&wafer)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(wafer, c)
}

// DeleteSpcWafer deletes a wafer
// @Tags      SpcWafer
// @Summary   删除晶圆
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "晶圆ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteWafer [delete]
func (a *WaferApi) DeleteSpcWafer(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = waferService.DeleteSpcWafer(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcWafer updates a wafer
// @Tags      SpcWafer
// @Summary   更新晶圆
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcWafer                   true  "晶圆信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateWafer [put]
func (a *WaferApi) UpdateSpcWafer(c *gin.Context) {
	var wafer spc.SpcWafer
	err := c.ShouldBindJSON(&wafer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = waferService.UpdateSpcWafer(&wafer)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcWafer finds a wafer by ID
// @Tags      SpcWafer
// @Summary   根据ID获取晶圆
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "晶圆ID"
// @Success   200   {object}  response.Response{data=spc.SpcWafer,msg=string}  "获取成功"
// @Router    /spc/findWafer [get]
func (a *WaferApi) FindSpcWafer(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	wafer, err := waferService.GetSpcWafer(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(wafer, c)
}

// GetSpcWaferList gets wafer list with pagination
// @Tags      SpcWafer
// @Summary   分页获取晶圆列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     lotId  query   int                                                false "批次ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getWaferList [get]
func (a *WaferApi) GetSpcWaferList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	lotIDStr := c.Query("lotId")
	lotID, _ := strconv.ParseUint(lotIDStr, 10, 32)

	list, total, err := waferService.GetSpcWaferList(pageInfo, uint(lotID))
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
