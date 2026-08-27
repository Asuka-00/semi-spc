package spc

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SiteApi
	AreaApi
	EquipmentApi
	ChartApi
	CollectApi
	AlarmApi
	CapabilityApi
}

var (
	siteService       = service.ServiceGroupApp.SpcServiceGroup.SiteService
	areaService       = service.ServiceGroupApp.SpcServiceGroup.AreaService
	equipmentService  = service.ServiceGroupApp.SpcServiceGroup.EquipmentService
	chartService      = service.ServiceGroupApp.SpcServiceGroup.ChartService
	collectService    = service.ServiceGroupApp.SpcServiceGroup.CollectService
	alarmService      = service.ServiceGroupApp.SpcServiceGroup.AlarmService
	capabilityService = service.ServiceGroupApp.SpcServiceGroup.CapabilityService
)
