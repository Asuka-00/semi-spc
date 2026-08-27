package spc

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SiteApi
	AreaApi
	EquipmentApi
	ChamberApi
	TechnologyApi
	ProductApi
	ProcessStepApi
	RecipeApi
	LotApi
	WaferApi
	ParameterApi
	SpecApi
	ChartApi
	ControlLimitApi
	RuleApi
	SampleApi
	MeasurementApi
	CollectApi
	AlarmApi
	OcapApi
	OcapExecutionApi
	CapabilityApi
	RuntimeApi
}

var (
	siteService           = service.ServiceGroupApp.SpcServiceGroup.SiteService
	areaService           = service.ServiceGroupApp.SpcServiceGroup.AreaService
	equipmentService      = service.ServiceGroupApp.SpcServiceGroup.EquipmentService
	chamberService        = service.ServiceGroupApp.SpcServiceGroup.ChamberService
	technologyService     = service.ServiceGroupApp.SpcServiceGroup.TechnologyService
	productService        = service.ServiceGroupApp.SpcServiceGroup.ProductService
	processStepService    = service.ServiceGroupApp.SpcServiceGroup.ProcessStepService
	recipeService         = service.ServiceGroupApp.SpcServiceGroup.RecipeService
	lotService            = service.ServiceGroupApp.SpcServiceGroup.LotService
	waferService          = service.ServiceGroupApp.SpcServiceGroup.WaferService
	parameterService      = service.ServiceGroupApp.SpcServiceGroup.ParameterService
	specService           = service.ServiceGroupApp.SpcServiceGroup.SpecService
	chartService          = service.ServiceGroupApp.SpcServiceGroup.ChartService
	controlLimitService   = service.ServiceGroupApp.SpcServiceGroup.ControlLimitService
	ruleService           = service.ServiceGroupApp.SpcServiceGroup.RuleService
	sampleService         = service.ServiceGroupApp.SpcServiceGroup.SampleService
	measurementService    = service.ServiceGroupApp.SpcServiceGroup.MeasurementService
	collectService        = service.ServiceGroupApp.SpcServiceGroup.CollectService
	alarmService          = service.ServiceGroupApp.SpcServiceGroup.AlarmService
	ocapService           = service.ServiceGroupApp.SpcServiceGroup.OcapService
	ocapExecutionService  = service.ServiceGroupApp.SpcServiceGroup.OcapExecutionService
	capabilityService     = service.ServiceGroupApp.SpcServiceGroup.CapabilityService
)
