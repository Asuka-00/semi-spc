package spc

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SiteApi
	ChartApi
	CollectApi
	AlarmApi
	CapabilityApi
}

var (
	siteService       = service.ServiceGroupApp.SpcServiceGroup.SiteService
	chartService      = service.ServiceGroupApp.SpcServiceGroup.ChartService
	collectService    = service.ServiceGroupApp.SpcServiceGroup.CollectService
	alarmService      = service.ServiceGroupApp.SpcServiceGroup.AlarmService
	capabilityService = service.ServiceGroupApp.SpcServiceGroup.CapabilityService
)
