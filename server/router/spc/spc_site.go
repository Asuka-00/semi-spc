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
		siteRouter.POST("site", siteApi.CreateSpcSite)
		siteRouter.PUT("site", siteApi.UpdateSpcSite)
		siteRouter.DELETE("site", siteApi.DeleteSpcSite)
	}
	{
		siteRouterWithoutRecord.GET("site", siteApi.GetSpcSite)
		siteRouterWithoutRecord.GET("getSiteList", siteApi.GetSpcSiteList)
	}
	{
		_ = sitePublicRouter
	}
}

type AreaRouter struct{}

func (r *AreaRouter) InitSpcAreaRouter(Router *gin.RouterGroup) {
	areaRouter := Router.Group("spc").Use(middleware.OperationRecord())
	areaRouterWithoutRecord := Router.Group("spc")

	areaApi := api.ApiGroupApp.SpcApiGroup.AreaApi
	{
		areaRouter.POST("createArea", areaApi.CreateSpcArea)
		areaRouter.PUT("updateArea", areaApi.UpdateSpcArea)
		areaRouter.DELETE("deleteArea", areaApi.DeleteSpcArea)
	}
	{
		areaRouterWithoutRecord.GET("findArea", areaApi.FindSpcArea)
		areaRouterWithoutRecord.GET("getAreaList", areaApi.GetSpcAreaList)
	}
}

type EquipmentRouter struct{}

func (r *EquipmentRouter) InitSpcEquipmentRouter(Router *gin.RouterGroup) {
	equipmentRouter := Router.Group("spc").Use(middleware.OperationRecord())
	equipmentRouterWithoutRecord := Router.Group("spc")

	equipmentApi := api.ApiGroupApp.SpcApiGroup.EquipmentApi
	{
		equipmentRouter.POST("createEquipment", equipmentApi.CreateSpcEquipment)
		equipmentRouter.PUT("updateEquipment", equipmentApi.UpdateSpcEquipment)
		equipmentRouter.DELETE("deleteEquipment", equipmentApi.DeleteSpcEquipment)
	}
	{
		equipmentRouterWithoutRecord.GET("findEquipment", equipmentApi.FindSpcEquipment)
		equipmentRouterWithoutRecord.GET("getEquipmentList", equipmentApi.GetSpcEquipmentList)
	}
}

type ChartRouter struct{}

func (r *ChartRouter) InitSpcChartRouter(Router *gin.RouterGroup) {
	chartRouter := Router.Group("spc").Use(middleware.OperationRecord())
	chartRouterWithoutRecord := Router.Group("spc")

	chartApi := api.ApiGroupApp.SpcApiGroup.ChartApi
	{
		chartRouter.POST("createChart", chartApi.CreateSpcChart)
		chartRouter.PUT("updateChart", chartApi.UpdateSpcChart)
		chartRouter.DELETE("deleteChart", chartApi.DeleteSpcChart)
	}
	{
		chartRouterWithoutRecord.GET("findChart", chartApi.FindSpcChart)
		chartRouterWithoutRecord.GET("getChartList", chartApi.GetSpcChartList)
	}
}

