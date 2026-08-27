package spc

import (
	spcService "github.com/flipped-aurora/gin-vue-admin/server/service/spc"
)

type RouterGroup struct {
	SiteRouter
	ChartRouter
	CollectRouter
	AlarmRouter
	CapabilityRouter
}

var (
	siteServiceApi   = spcService.ServiceGroup{}.SiteService
	collectApiRouter = spcService.ServiceGroup{}.CollectService
)
