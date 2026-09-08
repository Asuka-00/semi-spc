package spc

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EntityRouter struct{}
type CapabilityRouter struct{}
type DashboardRouter struct{}

func (r *EntityRouter) InitSpcEntityRouter(Router *gin.RouterGroup) {
	write := Router.Group("spc").Use(middleware.OperationRecord())
	read := Router.Group("spc")
	g := api.ApiGroupApp.SpcApiGroup

	{
		write.POST("createChamber", g.CreateSpcChamber)
		write.PUT("updateChamber", g.UpdateSpcChamber)
		write.DELETE("deleteChamber", g.DeleteSpcChamber)
		write.POST("createTechnology", g.CreateSpcTechnology)
		write.PUT("updateTechnology", g.UpdateSpcTechnology)
		write.DELETE("deleteTechnology", g.DeleteSpcTechnology)
		write.POST("createProduct", g.CreateSpcProduct)
		write.PUT("updateProduct", g.UpdateSpcProduct)
		write.DELETE("deleteProduct", g.DeleteSpcProduct)
		write.POST("createProcessStep", g.CreateSpcProcessStep)
		write.PUT("updateProcessStep", g.UpdateSpcProcessStep)
		write.DELETE("deleteProcessStep", g.DeleteSpcProcessStep)
		write.POST("createRecipe", g.CreateSpcRecipe)
		write.PUT("updateRecipe", g.UpdateSpcRecipe)
		write.DELETE("deleteRecipe", g.DeleteSpcRecipe)
		write.POST("createLot", g.CreateSpcLot)
		write.PUT("updateLot", g.UpdateSpcLot)
		write.DELETE("deleteLot", g.DeleteSpcLot)
		write.POST("createWafer", g.CreateSpcWafer)
		write.PUT("updateWafer", g.UpdateSpcWafer)
		write.DELETE("deleteWafer", g.DeleteSpcWafer)
		write.POST("createParameter", g.CreateSpcParameter)
		write.PUT("updateParameter", g.UpdateSpcParameter)
		write.DELETE("deleteParameter", g.DeleteSpcParameter)
		write.POST("createSpec", g.CreateSpcSpec)
		write.PUT("updateSpec", g.UpdateSpcSpec)
		write.DELETE("deleteSpec", g.DeleteSpcSpec)
		write.POST("createControlLimit", g.CreateSpcControlLimit)
		write.PUT("updateControlLimit", g.UpdateSpcControlLimit)
		write.DELETE("deleteControlLimit", g.DeleteSpcControlLimit)
		write.POST("createRule", g.CreateSpcRule)
		write.PUT("updateRule", g.UpdateSpcRule)
		write.DELETE("deleteRule", g.DeleteSpcRule)
		write.POST("createOcap", g.CreateSpcOcap)
		write.PUT("updateOcap", g.UpdateSpcOcap)
		write.DELETE("deleteOcap", g.DeleteSpcOcap)
		write.POST("createOcapExecution", g.CreateSpcOcapExecution)
		write.PUT("updateOcapExecution", g.UpdateSpcOcapExecution)
		write.POST("startOcapExecution", g.StartSpcOcapExecution)
		write.POST("completeOcapExecution", g.CompleteSpcOcapExecution)
		write.POST("calculateCapability", g.CalculateCapability)
		write.POST("saveChartRules", g.SaveChartRules)
	}
	{
		read.GET("findChamber", g.FindSpcChamber)
		read.GET("getChamberList", g.GetSpcChamberList)
		read.GET("findTechnology", g.FindSpcTechnology)
		read.GET("getTechnologyList", g.GetSpcTechnologyList)
		read.GET("findProduct", g.FindSpcProduct)
		read.GET("getProductList", g.GetSpcProductList)
		read.GET("findProcessStep", g.FindSpcProcessStep)
		read.GET("getProcessStepList", g.GetSpcProcessStepList)
		read.GET("findRecipe", g.FindSpcRecipe)
		read.GET("getRecipeList", g.GetSpcRecipeList)
		read.GET("findLot", g.FindSpcLot)
		read.GET("getLotList", g.GetSpcLotList)
		read.GET("findWafer", g.FindSpcWafer)
		read.GET("getWaferList", g.GetSpcWaferList)
		read.GET("findParameter", g.FindSpcParameter)
		read.GET("getParameterList", g.GetSpcParameterList)
		read.GET("findSpec", g.FindSpcSpec)
		read.GET("getSpecList", g.GetSpcSpecList)
		read.GET("findControlLimit", g.FindSpcControlLimit)
		read.GET("getControlLimitList", g.GetSpcControlLimitList)
		read.GET("findRule", g.FindSpcRule)
		read.GET("getRuleList", g.GetSpcRuleList)
		read.GET("findOcap", g.FindSpcOcap)
		read.GET("getOcapList", g.GetSpcOcapList)
		read.GET("findOcapExecution", g.FindSpcOcapExecution)
		read.GET("getOcapExecutionList", g.GetSpcOcapExecutionList)
		read.GET("getSampleList", g.GetSpcSampleList)
		read.GET("findSample", g.FindSpcSample)
		read.GET("getMeasurementList", g.GetSpcMeasurementList)
		read.GET("getCapabilityList", g.GetCapabilityList)
		read.GET("getCapabilityHistory", g.GetCapabilityHistory)
		read.GET("getDashboardOverview", g.GetDashboardOverview)
		read.GET("getChartRuntime", g.GetChartRuntime)
		read.GET("getRuleCatalog", g.GetRuleCatalog)
	}
}
