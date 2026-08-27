package spc

import (
	spcService "github.com/flipped-aurora/gin-vue-admin/server/service/spc"
)

type RouterGroup struct {
	SiteRouter
	AreaRouter
	EquipmentRouter
	ChamberRouter
	TechnologyRouter
	ProductRouter
	ProcessStepRouter
	RecipeRouter
	LotRouter
	WaferRouter
	ParameterRouter
	SpecRouter
	ChartRouter
	ControlLimitRouter
	RuleRouter
	SampleRouter
	MeasurementRouter
	CollectRouter
	AlarmRouter
	OcapRouter
	OcapExecutionRouter
	CapabilityRouter
	RuntimeRouter
}

var (
	_ = spcService.ServiceGroup{}
)
