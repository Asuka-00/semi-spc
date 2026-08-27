package spc

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChamberRouter struct{}

func (r *ChamberRouter) InitSpcChamberRouter(Router *gin.RouterGroup) {
	chamberRouter := Router.Group("spc").Use(middleware.OperationRecord())
	chamberRouterWithoutRecord := Router.Group("spc")
	chamberApi := api.ApiGroupApp.SpcApiGroup.ChamberApi
	{
		chamberRouter.POST("createChamber", chamberApi.CreateSpcChamber)
		chamberRouter.PUT("updateChamber", chamberApi.UpdateSpcChamber)
		chamberRouter.DELETE("deleteChamber", chamberApi.DeleteSpcChamber)
	}
	{
		chamberRouterWithoutRecord.GET("findChamber", chamberApi.FindSpcChamber)
		chamberRouterWithoutRecord.GET("getChamberList", chamberApi.GetSpcChamberList)
	}
}

type TechnologyRouter struct{}

func (r *TechnologyRouter) InitSpcTechnologyRouter(Router *gin.RouterGroup) {
	technologyRouter := Router.Group("spc").Use(middleware.OperationRecord())
	technologyRouterWithoutRecord := Router.Group("spc")
	technologyApi := api.ApiGroupApp.SpcApiGroup.TechnologyApi
	{
		technologyRouter.POST("createTechnology", technologyApi.CreateSpcTechnology)
		technologyRouter.PUT("updateTechnology", technologyApi.UpdateSpcTechnology)
		technologyRouter.DELETE("deleteTechnology", technologyApi.DeleteSpcTechnology)
	}
	{
		technologyRouterWithoutRecord.GET("findTechnology", technologyApi.FindSpcTechnology)
		technologyRouterWithoutRecord.GET("getTechnologyList", technologyApi.GetSpcTechnologyList)
	}
}

type ProductRouter struct{}

func (r *ProductRouter) InitSpcProductRouter(Router *gin.RouterGroup) {
	productRouter := Router.Group("spc").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("spc")
	productApi := api.ApiGroupApp.SpcApiGroup.ProductApi
	{
		productRouter.POST("createProduct", productApi.CreateSpcProduct)
		productRouter.PUT("updateProduct", productApi.UpdateSpcProduct)
		productRouter.DELETE("deleteProduct", productApi.DeleteSpcProduct)
	}
	{
		productRouterWithoutRecord.GET("findProduct", productApi.FindSpcProduct)
		productRouterWithoutRecord.GET("getProductList", productApi.GetSpcProductList)
	}
}

type ProcessStepRouter struct{}

func (r *ProcessStepRouter) InitSpcProcessStepRouter(Router *gin.RouterGroup) {
	processStepRouter := Router.Group("spc").Use(middleware.OperationRecord())
	processStepRouterWithoutRecord := Router.Group("spc")
	processStepApi := api.ApiGroupApp.SpcApiGroup.ProcessStepApi
	{
		processStepRouter.POST("createProcessStep", processStepApi.CreateSpcProcessStep)
		processStepRouter.PUT("updateProcessStep", processStepApi.UpdateSpcProcessStep)
		processStepRouter.DELETE("deleteProcessStep", processStepApi.DeleteSpcProcessStep)
	}
	{
		processStepRouterWithoutRecord.GET("findProcessStep", processStepApi.FindSpcProcessStep)
		processStepRouterWithoutRecord.GET("getProcessStepList", processStepApi.GetSpcProcessStepList)
	}
}

type RecipeRouter struct{}

func (r *RecipeRouter) InitSpcRecipeRouter(Router *gin.RouterGroup) {
	recipeRouter := Router.Group("spc").Use(middleware.OperationRecord())
	recipeRouterWithoutRecord := Router.Group("spc")
	recipeApi := api.ApiGroupApp.SpcApiGroup.RecipeApi
	{
		recipeRouter.POST("createRecipe", recipeApi.CreateSpcRecipe)
		recipeRouter.PUT("updateRecipe", recipeApi.UpdateSpcRecipe)
		recipeRouter.DELETE("deleteRecipe", recipeApi.DeleteSpcRecipe)
	}
	{
		recipeRouterWithoutRecord.GET("findRecipe", recipeApi.FindSpcRecipe)
		recipeRouterWithoutRecord.GET("getRecipeList", recipeApi.GetSpcRecipeList)
	}
}

type LotRouter struct{}

func (r *LotRouter) InitSpcLotRouter(Router *gin.RouterGroup) {
	lotRouter := Router.Group("spc").Use(middleware.OperationRecord())
	lotRouterWithoutRecord := Router.Group("spc")
	lotApi := api.ApiGroupApp.SpcApiGroup.LotApi
	{
		lotRouter.POST("createLot", lotApi.CreateSpcLot)
		lotRouter.PUT("updateLot", lotApi.UpdateSpcLot)
		lotRouter.DELETE("deleteLot", lotApi.DeleteSpcLot)
		lotRouter.POST("holdLot", lotApi.HoldSpcLot)
		lotRouter.POST("releaseLot", lotApi.ReleaseSpcLot)
	}
	{
		lotRouterWithoutRecord.GET("findLot", lotApi.FindSpcLot)
		lotRouterWithoutRecord.GET("getLotList", lotApi.GetSpcLotList)
	}
}

type WaferRouter struct{}

func (r *WaferRouter) InitSpcWaferRouter(Router *gin.RouterGroup) {
	waferRouter := Router.Group("spc").Use(middleware.OperationRecord())
	waferRouterWithoutRecord := Router.Group("spc")
	waferApi := api.ApiGroupApp.SpcApiGroup.WaferApi
	{
		waferRouter.POST("createWafer", waferApi.CreateSpcWafer)
		waferRouter.PUT("updateWafer", waferApi.UpdateSpcWafer)
		waferRouter.DELETE("deleteWafer", waferApi.DeleteSpcWafer)
	}
	{
		waferRouterWithoutRecord.GET("findWafer", waferApi.FindSpcWafer)
		waferRouterWithoutRecord.GET("getWaferList", waferApi.GetSpcWaferList)
	}
}

type ParameterRouter struct{}

func (r *ParameterRouter) InitSpcParameterRouter(Router *gin.RouterGroup) {
	parameterRouter := Router.Group("spc").Use(middleware.OperationRecord())
	parameterRouterWithoutRecord := Router.Group("spc")
	parameterApi := api.ApiGroupApp.SpcApiGroup.ParameterApi
	{
		parameterRouter.POST("createParameter", parameterApi.CreateSpcParameter)
		parameterRouter.PUT("updateParameter", parameterApi.UpdateSpcParameter)
		parameterRouter.DELETE("deleteParameter", parameterApi.DeleteSpcParameter)
	}
	{
		parameterRouterWithoutRecord.GET("findParameter", parameterApi.FindSpcParameter)
		parameterRouterWithoutRecord.GET("getParameterList", parameterApi.GetSpcParameterList)
	}
}

type SpecRouter struct{}

func (r *SpecRouter) InitSpcSpecRouter(Router *gin.RouterGroup) {
	specRouter := Router.Group("spc").Use(middleware.OperationRecord())
	specRouterWithoutRecord := Router.Group("spc")
	specApi := api.ApiGroupApp.SpcApiGroup.SpecApi
	{
		specRouter.POST("createSpec", specApi.CreateSpcSpec)
		specRouter.PUT("updateSpec", specApi.UpdateSpcSpec)
		specRouter.DELETE("deleteSpec", specApi.DeleteSpcSpec)
	}
	{
		specRouterWithoutRecord.GET("findSpec", specApi.FindSpcSpec)
		specRouterWithoutRecord.GET("getSpecList", specApi.GetSpcSpecList)
	}
}

type ControlLimitRouter struct{}

func (r *ControlLimitRouter) InitSpcControlLimitRouter(Router *gin.RouterGroup) {
	controlLimitRouter := Router.Group("spc").Use(middleware.OperationRecord())
	controlLimitRouterWithoutRecord := Router.Group("spc")
	controlLimitApi := api.ApiGroupApp.SpcApiGroup.ControlLimitApi
	{
		controlLimitRouter.POST("createControlLimit", controlLimitApi.CreateSpcControlLimit)
		controlLimitRouter.PUT("updateControlLimit", controlLimitApi.UpdateSpcControlLimit)
		controlLimitRouter.DELETE("deleteControlLimit", controlLimitApi.DeleteSpcControlLimit)
	}
	{
		controlLimitRouterWithoutRecord.GET("findControlLimit", controlLimitApi.FindSpcControlLimit)
		controlLimitRouterWithoutRecord.GET("getControlLimitList", controlLimitApi.GetSpcControlLimitList)
	}
}

type RuleRouter struct{}

func (r *RuleRouter) InitSpcRuleRouter(Router *gin.RouterGroup) {
	ruleRouter := Router.Group("spc").Use(middleware.OperationRecord())
	ruleRouterWithoutRecord := Router.Group("spc")
	ruleApi := api.ApiGroupApp.SpcApiGroup.RuleApi
	{
		ruleRouter.POST("createRule", ruleApi.CreateSpcRule)
		ruleRouter.PUT("updateRule", ruleApi.UpdateSpcRule)
		ruleRouter.DELETE("deleteRule", ruleApi.DeleteSpcRule)
	}
	{
		ruleRouterWithoutRecord.GET("findRule", ruleApi.FindSpcRule)
		ruleRouterWithoutRecord.GET("getRuleList", ruleApi.GetSpcRuleList)
	}
}

type SampleRouter struct{}

func (r *SampleRouter) InitSpcSampleRouter(Router *gin.RouterGroup) {
	sampleRouterWithoutRecord := Router.Group("spc")
	sampleApi := api.ApiGroupApp.SpcApiGroup.SampleApi
	{
		sampleRouterWithoutRecord.GET("findSample", sampleApi.FindSpcSample)
		sampleRouterWithoutRecord.GET("getSampleList", sampleApi.GetSpcSampleList)
	}
}

type MeasurementRouter struct{}

func (r *MeasurementRouter) InitSpcMeasurementRouter(Router *gin.RouterGroup) {
	measurementRouterWithoutRecord := Router.Group("spc")
	measurementApi := api.ApiGroupApp.SpcApiGroup.MeasurementApi
	{
		measurementRouterWithoutRecord.GET("getMeasurementList", measurementApi.GetSpcMeasurementList)
	}
}

type OcapRouter struct{}

func (r *OcapRouter) InitSpcOcapRouter(Router *gin.RouterGroup) {
	ocapRouter := Router.Group("spc").Use(middleware.OperationRecord())
	ocapRouterWithoutRecord := Router.Group("spc")
	ocapApi := api.ApiGroupApp.SpcApiGroup.OcapApi
	{
		ocapRouter.POST("createOcap", ocapApi.CreateSpcOcap)
		ocapRouter.PUT("updateOcap", ocapApi.UpdateSpcOcap)
		ocapRouter.DELETE("deleteOcap", ocapApi.DeleteSpcOcap)
		ocapRouter.POST("startOcap", ocapApi.StartOcap)
	}
	{
		ocapRouterWithoutRecord.GET("findOcap", ocapApi.FindSpcOcap)
		ocapRouterWithoutRecord.GET("getOcapList", ocapApi.GetSpcOcapList)
	}
}

type OcapExecutionRouter struct{}

func (r *OcapExecutionRouter) InitSpcOcapExecutionRouter(Router *gin.RouterGroup) {
	ocapExecutionRouter := Router.Group("spc").Use(middleware.OperationRecord())
	ocapExecutionRouterWithoutRecord := Router.Group("spc")
	ocapExecutionApi := api.ApiGroupApp.SpcApiGroup.OcapExecutionApi
	{
		ocapExecutionRouter.PUT("updateOcapExecution", ocapExecutionApi.UpdateSpcOcapExecution)
	}
	{
		ocapExecutionRouterWithoutRecord.GET("getOcapExecutionList", ocapExecutionApi.GetSpcOcapExecutionList)
	}
}

type CapabilityRouter struct{}

func (r *CapabilityRouter) InitSpcCapabilityRouter(Router *gin.RouterGroup) {
	capabilityRouter := Router.Group("spc").Use(middleware.OperationRecord())
	capabilityRouterWithoutRecord := Router.Group("spc")
	capabilityApi := api.ApiGroupApp.SpcApiGroup.CapabilityApi
	{
		capabilityRouter.POST("calculateCapability", capabilityApi.CalculateCapability)
	}
	{
		capabilityRouterWithoutRecord.GET("getCapabilityHistory", capabilityApi.GetCapabilityHistory)
	}
}

type RuntimeRouter struct{}

func (r *RuntimeRouter) InitSpcRuntimeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	runtimeRouter := Router.Group("spc").Use(middleware.OperationRecord())
	runtimeRouterWithoutRecord := Router.Group("spc")
	runtimePublicRouter := PublicRouter.Group("spc").Use(middleware.SpcCollectAuth()) // 公开CSV采集需要认证
	
	runtimeApi := api.ApiGroupApp.SpcApiGroup.RuntimeApi
	{
		runtimeRouter.POST("calculateLimits", runtimeApi.CalculateLimits)
		runtimeRouter.POST("collectCsv", runtimeApi.CollectCsv) // JWT认证
	}
	{
		runtimePublicRouter.POST("collectCsv", runtimeApi.CollectCsv) // JWT或API Token认证
	}
	{
		runtimeRouterWithoutRecord.GET("getChartRuntime", runtimeApi.GetChartRuntime)
		runtimeRouterWithoutRecord.GET("getCapability", runtimeApi.GetCapability)
		runtimeRouterWithoutRecord.GET("getDashboard", runtimeApi.GetDashboard)
	}
}
