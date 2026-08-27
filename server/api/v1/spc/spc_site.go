package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SiteApi struct{}

// CreateSpcSite
// @Tags      SpcSite
// @Summary   创建厂区
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcSite                    true  "厂区信息"
// @Success   200   {object}  response.Response{msg=string}  "创建成功"
// @Router    /spc/site [post]
func (a *SiteApi) CreateSpcSite(c *gin.Context) {
	var site spc.SpcSite
	err := c.ShouldBindJSON(&site)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = siteService.CreateSpcSite(&site)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(site, c)
}

// DeleteSpcSite
// @Tags      SpcSite
// @Summary   删除厂区
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "厂区ID"
// @Success   200   {object}  response.Response{msg=string}  "删除成功"
// @Router    /spc/site [delete]
func (a *SiteApi) DeleteSpcSite(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindJSON(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = siteService.DeleteSpcSite(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSpcSite
// @Tags      SpcSite
// @Summary   更新厂区
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcSite                    true  "厂区信息"
// @Success   200   {object}  response.Response{msg=string}  "更新成功"
// @Router    /spc/site [put]
func (a *SiteApi) UpdateSpcSite(c *gin.Context) {
	var site spc.SpcSite
	err := c.ShouldBindJSON(&site)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = siteService.UpdateSpcSite(&site)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetSpcSite
// @Tags      SpcSite
// @Summary   根据ID获取厂区
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.GetById                               true  "厂区ID"
// @Success   200   {object}  response.Response{data=spc.SpcSite,msg=string}  "获取成功"
// @Router    /spc/site [get]
func (a *SiteApi) GetSpcSite(c *gin.Context) {
	var idReq request.GetById
	err := c.ShouldBindQuery(&idReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	site, err := siteService.GetSpcSite(uint(idReq.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(site, c)
}

// GetSpcSiteList
// @Tags      SpcSite
// @Summary   分页获取厂区列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                  true  "分页参数"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取成功"
// @Router    /spc/getSiteList [get]
func (a *SiteApi) GetSpcSiteList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := siteService.GetSpcSiteList(pageInfo)
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
