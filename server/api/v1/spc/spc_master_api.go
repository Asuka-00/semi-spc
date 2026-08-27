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

type AreaApi struct{}

// CreateSpcArea
// @Tags      SpcArea
// @Summary   创建区域
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcArea                    true  "区域信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createArea [post]
func (a *AreaApi) CreateSpcArea(c *gin.Context) {
	var area spc.SpcArea
	err := c.ShouldBindJSON(&area)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = areaService.CreateSpcArea(&area)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(area, c)
}

// DeleteSpcArea
// @Tags      SpcArea
// @Summary   删除区域
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "区域ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteArea [delete]
func (a *AreaApi) DeleteSpcArea(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = areaService.DeleteSpcArea(idReq.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcArea
// @Tags      SpcArea
// @Summary   更新区域
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcArea                    true  "区域信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateArea [put]
func (a *AreaApi) UpdateSpcArea(c *gin.Context) {
	var area spc.SpcArea
	err := c.ShouldBindJSON(&area)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = areaService.UpdateSpcArea(&area)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcArea
// @Tags      SpcArea
// @Summary   根据ID获取区域
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "区域ID"
// @Success   200   {object}  response.Response{data=spc.SpcArea,msg=string}  "获取成功"
// @Router    /spc/findArea [get]
func (a *AreaApi) FindSpcArea(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	area, err := areaService.GetSpcArea(idReq.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(area, c)
}

// GetSpcAreaList
// @Tags      SpcArea
// @Summary   分页获取区域列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     siteId  query   int                                                false "厂区ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getAreaList [get]
func (a *AreaApi) GetSpcAreaList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	siteIDStr := c.Query("siteId")
	siteID, _ := strconv.ParseUint(siteIDStr, 10, 32)

	list, total, err := areaService.GetSpcAreaList(pageInfo, uint(siteID))
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

type EquipmentApi struct{}

// CreateSpcEquipment
// @Tags      SpcEquipment
// @Summary   创建设备
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcEquipment               true  "设备信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/createEquipment [post]
func (a *EquipmentApi) CreateSpcEquipment(c *gin.Context) {
	var equipment spc.SpcEquipment
	err := c.ShouldBindJSON(&equipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = equipmentService.CreateSpcEquipment(&equipment)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(equipment, c)
}

// DeleteSpcEquipment
// @Tags      SpcEquipment
// @Summary   删除设备
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "设备ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/deleteEquipment [delete]
func (a *EquipmentApi) DeleteSpcEquipment(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = equipmentService.DeleteSpcEquipment(idReq.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcEquipment
// @Tags      SpcEquipment
// @Summary   更新设备
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcEquipment               true  "设备信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/updateEquipment [put]
func (a *EquipmentApi) UpdateSpcEquipment(c *gin.Context) {
	var equipment spc.SpcEquipment
	err := c.ShouldBindJSON(&equipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = equipmentService.UpdateSpcEquipment(&equipment)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSpcEquipment
// @Tags      SpcEquipment
// @Summary   根据ID获取设备
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "设备ID"
// @Success   200   {object}  response.Response{data=spc.SpcEquipment,msg=string}  "获取成功"
// @Router    /spc/findEquipment [get]
func (a *EquipmentApi) FindSpcEquipment(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	equipment, err := equipmentService.GetSpcEquipment(idReq.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(equipment, c)
}

// GetSpcEquipmentList
// @Tags      SpcEquipment
// @Summary   分页获取设备列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Param     siteId  query   int                                                false "厂区ID"
// @Param     areaId  query   int                                                false "区域ID"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getEquipmentList [get]
func (a *EquipmentApi) GetSpcEquipmentList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	siteIDStr := c.Query("siteId")
	areaIDStr := c.Query("areaId")
	siteID, _ := strconv.ParseUint(siteIDStr, 10, 32)
	areaID, _ := strconv.ParseUint(areaIDStr, 10, 32)

	list, total, err := equipmentService.GetSpcEquipmentList(pageInfo, uint(siteID), uint(areaID))
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
