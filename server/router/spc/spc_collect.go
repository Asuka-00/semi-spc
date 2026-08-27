package spc

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CollectRouter struct{}

func (r *CollectRouter) InitSpcCollectRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	collectRouter := Router.Group("spc")
	collectPublicRouter := PublicRouter.Group("spc")

	collectApi := api.ApiGroupApp.SpcApiGroup.CollectApi
	{
		collectRouter.POST("collect", collectApi.CollectData) // 数据采集
	}
	{
		collectPublicRouter.POST("collect", collectApi.CollectData) // 公开的采集接口(可用API Token)
	}
}
