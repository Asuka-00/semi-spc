package spc

import (
	spcService "github.com/flipped-aurora/gin-vue-admin/server/service/spc"
)

type RouterGroup struct {
	SiteRouter
	AreaRouter
	EquipmentRouter
	ChartRouter
	CollectRouter
	AlarmRouter
	CapabilityRouter
}

var (
	_ = spcService.ServiceGroup{}
)
