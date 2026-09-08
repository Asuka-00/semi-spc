package spc

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CollectRouter struct{}

func (r *CollectRouter) InitSpcCollectRouter(Router *gin.RouterGroup) {
	collectRouter := Router.Group("spc")
	collectApi := api.ApiGroupApp.SpcApiGroup.CollectApi
	// 采集走 JWT 组。GVA API Token 本身是 JWT，设备上传使用 x-token 即可，
	// 不能再在 PublicGroup 注册同一路径，否则 Gin 会 panic。
	collectRouter.POST("collect", collectApi.CollectData)
}
