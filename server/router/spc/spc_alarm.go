package spc

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlarmRouter struct{}

func (r *AlarmRouter) InitSpcAlarmRouter(Router *gin.RouterGroup) {
	alarmRouter := Router.Group("spc").Use(middleware.OperationRecord())
	alarmRouterWithoutRecord := Router.Group("spc")

	alarmApi := api.ApiGroupApp.SpcApiGroup.AlarmApi
	{
		alarmRouter.POST("acknowledgeAlarm", alarmApi.AcknowledgeAlarm) // 确认告警
		alarmRouter.POST("closeAlarm", alarmApi.CloseAlarm)             // 关闭告警
	}
	{
		alarmRouterWithoutRecord.GET("getAlarmList", alarmApi.GetSpcAlarmList)           // 获取告警列表
		alarmRouterWithoutRecord.GET("getAlarmStatistics", alarmApi.GetAlarmStatistics) // 获取告警统计
	}
}
