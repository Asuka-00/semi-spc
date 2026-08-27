package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	holder(publicGroup, privateGroup)

	// 注册SPC路由
	spcRouter := router.RouterGroupApp.Spc
	spcRouter.InitSpcSiteRouter(privateGroup, publicGroup)
	spcRouter.InitSpcAreaRouter(privateGroup)
	spcRouter.InitSpcEquipmentRouter(privateGroup)
	spcRouter.InitSpcChamberRouter(privateGroup)
	spcRouter.InitSpcTechnologyRouter(privateGroup)
	spcRouter.InitSpcProductRouter(privateGroup)
	spcRouter.InitSpcProcessStepRouter(privateGroup)
	spcRouter.InitSpcRecipeRouter(privateGroup)
	spcRouter.InitSpcLotRouter(privateGroup)
	spcRouter.InitSpcWaferRouter(privateGroup)
	spcRouter.InitSpcParameterRouter(privateGroup)
	spcRouter.InitSpcSpecRouter(privateGroup)
	spcRouter.InitSpcChartRouter(privateGroup)
	spcRouter.InitSpcControlLimitRouter(privateGroup)
	spcRouter.InitSpcRuleRouter(privateGroup)
	spcRouter.InitSpcSampleRouter(privateGroup)
	spcRouter.InitSpcMeasurementRouter(privateGroup)
	spcRouter.InitSpcCollectRouter(privateGroup, publicGroup)
	spcRouter.InitSpcAlarmRouter(privateGroup)
	spcRouter.InitSpcOcapRouter(privateGroup)
	spcRouter.InitSpcOcapExecutionRouter(privateGroup)
	spcRouter.InitSpcCapabilityRouter(privateGroup)
	spcRouter.InitSpcRuntimeRouter(privateGroup, publicGroup)
}
