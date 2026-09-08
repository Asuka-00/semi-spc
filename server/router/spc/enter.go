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
	EntityRouter
}

var (
	_ = spcService.ServiceGroup{}
)
