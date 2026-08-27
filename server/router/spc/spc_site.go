package spc

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SiteRouter struct{}

func (r *SiteRouter) InitSpcSiteRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	siteRouter := Router.Group("spc").Use(middleware.OperationRecord())
	siteRouterWithoutRecord := Router.Group("spc")
	sitePublicRouter := PublicRouter.Group("spc")

	siteApi := api.ApiGroupApp.SpcApiGroup.SiteApi
	{
		siteRouter.POST("site", siteApi.CreateSpcSite)       // 创建厂区
		siteRouter.PUT("site", siteApi.UpdateSpcSite)        // 更新厂区
		siteRouter.DELETE("site", siteApi.DeleteSpcSite)     // 删除厂区
	}
	{
		siteRouterWithoutRecord.GET("site", siteApi.GetSpcSite)           // 获取单个厂区
		siteRouterWithoutRecord.GET("getSiteList", siteApi.GetSpcSiteList) // 获取厂区列表
	}
	{
		_ = sitePublicRouter
	}
}
