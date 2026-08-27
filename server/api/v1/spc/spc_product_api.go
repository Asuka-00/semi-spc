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

type ChamberApi struct{}

// CreateSpcChamber
// @Tags      SpcChamber
// @Summary   创建腔体
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcChamber                 true  "腔体信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createChamber [post]
func (a *ChamberApi) CreateSpcChamber(c *gin.Context) {
	var chamber spc.SpcChamber
	err := c.ShouldBindJSON(&chamber)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = chamberService.CreateSpcChamber(&chamber)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(chamber, c)
}

// DeleteSpcChamber
// @Tags      SpcChamber
// @Summary   删除腔体
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "腔体ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteChamber [delete]
func (a *ChamberApi) DeleteSpcChamber(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = chamberService.DeleteSpcChamber(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcChamber
// @Tags      SpcChamber
// @Summary   更新腔体
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcChamber                 true  "腔体信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateChamber [put]
func (a *ChamberApi) UpdateSpcChamber(c *gin.Context) {
	var chamber spc.SpcChamber
	err := c.ShouldBindJSON(&chamber)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = chamberService.UpdateSpcChamber(&chamber)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcChamber
// @Tags      SpcChamber
// @Summary   根据ID获取腔体
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "腔体ID"
// @Success   200   {object}  response.Response{data=spc.SpcChamber,msg=string}  "获取成功"
// @Router    /spc/findChamber [get]
func (a *ChamberApi) FindSpcChamber(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chamber, err := chamberService.GetSpcChamber(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(chamber, c)
}

// GetSpcChamberList
// @Tags      SpcChamber
// @Summary   分页获取腔体列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     equipmentId  query   int                                           false "设备ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getChamberList [get]
func (a *ChamberApi) GetSpcChamberList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	equipmentIDStr := c.Query("equipmentId")
	equipmentID, _ := strconv.ParseUint(equipmentIDStr, 10, 32)

	list, total, err := chamberService.GetSpcChamberList(pageInfo, uint(equipmentID))
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

type TechnologyApi struct{}

// CreateSpcTechnology
// @Tags      SpcTechnology
// @Summary   创建工艺技术
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcTechnology              true  "工艺技术信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createTechnology [post]
func (a *TechnologyApi) CreateSpcTechnology(c *gin.Context) {
	var tech spc.SpcTechnology
	err := c.ShouldBindJSON(&tech)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = technologyService.CreateSpcTechnology(&tech)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(tech, c)
}

// DeleteSpcTechnology
// @Tags      SpcTechnology
// @Summary   删除工艺技术
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "工艺技术ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteTechnology [delete]
func (a *TechnologyApi) DeleteSpcTechnology(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = technologyService.DeleteSpcTechnology(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcTechnology
// @Tags      SpcTechnology
// @Summary   更新工艺技术
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcTechnology              true  "工艺技术信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateTechnology [put]
func (a *TechnologyApi) UpdateSpcTechnology(c *gin.Context) {
	var tech spc.SpcTechnology
	err := c.ShouldBindJSON(&tech)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = technologyService.UpdateSpcTechnology(&tech)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcTechnology
// @Tags      SpcTechnology
// @Summary   根据ID获取工艺技术
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "工艺技术ID"
// @Success   200   {object}  response.Response{data=spc.SpcTechnology,msg=string}  "获取成功"
// @Router    /spc/findTechnology [get]
func (a *TechnologyApi) FindSpcTechnology(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	tech, err := technologyService.GetSpcTechnology(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(tech, c)
}

// GetSpcTechnologyList
// @Tags      SpcTechnology
// @Summary   分页获取工艺技术列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getTechnologyList [get]
func (a *TechnologyApi) GetSpcTechnologyList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := technologyService.GetSpcTechnologyList(pageInfo)
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

type ProductApi struct{}

// CreateSpcProduct
// @Tags      SpcProduct
// @Summary   创建产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcProduct                 true  "产品信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createProduct [post]
func (a *ProductApi) CreateSpcProduct(c *gin.Context) {
	var product spc.SpcProduct
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = productService.CreateSpcProduct(&product)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(product, c)
}

// DeleteSpcProduct
// @Tags      SpcProduct
// @Summary   删除产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "产品ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteProduct [delete]
func (a *ProductApi) DeleteSpcProduct(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = productService.DeleteSpcProduct(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcProduct
// @Tags      SpcProduct
// @Summary   更新产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcProduct                 true  "产品信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateProduct [put]
func (a *ProductApi) UpdateSpcProduct(c *gin.Context) {
	var product spc.SpcProduct
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = productService.UpdateSpcProduct(&product)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcProduct
// @Tags      SpcProduct
// @Summary   根据ID获取产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "产品ID"
// @Success   200   {object}  response.Response{data=spc.SpcProduct,msg=string}  "获取成功"
// @Router    /spc/findProduct [get]
func (a *ProductApi) FindSpcProduct(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	product, err := productService.GetSpcProduct(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(product, c)
}

// GetSpcProductList
// @Tags      SpcProduct
// @Summary   分页获取产品列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     techId  query   int                                                false "工艺技术ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getProductList [get]
func (a *ProductApi) GetSpcProductList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	techIDStr := c.Query("techId")
	techID, _ := strconv.ParseUint(techIDStr, 10, 32)

	list, total, err := productService.GetSpcProductList(pageInfo, uint(techID))
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
