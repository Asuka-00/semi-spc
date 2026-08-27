package spc_test

import (
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/stretchr/testify/assert"
)

var siteService = service.ServiceGroupApp.SpcServiceGroup.SiteService
var areaService = service.ServiceGroupApp.SpcServiceGroup.AreaService
var parameterService = service.ServiceGroupApp.SpcServiceGroup.ParameterService
var specService = service.ServiceGroupApp.SpcServiceGroup.SpecService
var chartService = service.ServiceGroupApp.SpcServiceGroup.ChartService
var collectService = service.ServiceGroupApp.SpcServiceGroup.CollectService

func TestDuplicateCodeRejection(t *testing.T) {
	// 清理
	global.GVA_DB.Where("code = ?", "TEST_DUP").Delete(&spc.SpcSite{})

	// 创建第一个Site
	site1 := &spc.SpcSite{
		Code:   "TEST_DUP",
		Name:   "测试厂区1",
		Status: 1,
	}
	err := siteService.CreateSpcSite(site1)
	assert.NoError(t, err)

	// 尝试创建重复代码的Site
	site2 := &spc.SpcSite{
		Code:   "TEST_DUP",
		Name:   "测试厂区2",
		Status: 1,
	}
	err = siteService.CreateSpcSite(site2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "代码已存在")

	// 清理
	global.GVA_DB.Unscoped().Delete(site1)
}

func TestReferentialIntegrityDelete(t *testing.T) {
	// 清理
	global.GVA_DB.Where("code = ?", "TEST_REF_SITE").Delete(&spc.SpcSite{})
	global.GVA_DB.Where("code = ?", "TEST_REF_AREA").Delete(&spc.SpcArea{})

	// 创建Site
	site := &spc.SpcSite{
		Code:   "TEST_REF_SITE",
		Name:   "测试厂区",
		Status: 1,
	}
	err := siteService.CreateSpcSite(site)
	assert.NoError(t, err)

	// 创建Area引用Site
	area := &spc.SpcArea{
		Code:   "TEST_REF_AREA",
		Name:   "测试区域",
		SiteID: site.ID,
		Status: 1,
	}
	err = areaService.CreateSpcArea(area)
	assert.NoError(t, err)

	// 尝试删除有子记录的Site
	err = siteService.DeleteSpcSite(site.ID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "存在关联数据")

	// 清理
	global.GVA_DB.Unscoped().Delete(area)
	global.GVA_DB.Unscoped().Delete(site)
}

func TestSpecValidationUSLGreaterThanLSL(t *testing.T) {
	// 清理
	global.GVA_DB.Where("code = ?", "TEST_PARAM").Delete(&spc.SpcParameter{})

	// 创建Parameter
	param := &spc.SpcParameter{
		Code:          "TEST_PARAM",
		Name:          "测试参数",
		DataType:      "VARIABLE",
		UnitOfMeasure: "nm",
		DecimalPlaces: 2,
		Status:        1,
	}
	err := parameterService.CreateSpcParameter(param)
	assert.NoError(t, err)

	// 尝试创建USL < LSL的Spec
	usl := 100.0
	lsl := 200.0
	spec := &spc.SpcSpec{
		ParameterID: param.ID,
		Version:     "1.0",
		USL:         &usl,
		LSL:         &lsl,
		Status:      1,
	}
	err = specService.CreateSpcSpec(spec)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "规格上限必须大于下限")

	// 清理
	global.GVA_DB.Unscoped().Delete(param)
}

func TestChartTypeDataTypeMismatch(t *testing.T) {
	// 清理
	global.GVA_DB.Where("code = ?", "TEST_PARAM_ATTR").Delete(&spc.SpcParameter{})
	global.GVA_DB.Where("code = ?", "TEST_CHART_MISMATCH").Delete(&spc.SpcChart{})

	// 创建ATTRIBUTE类型的Parameter
	param := &spc.SpcParameter{
		Code:          "TEST_PARAM_ATTR",
		Name:          "测试属性参数",
		DataType:      "ATTRIBUTE",
		UnitOfMeasure: "defects",
		DecimalPlaces: 0,
		Status:        1,
	}
	err := parameterService.CreateSpcParameter(param)
	assert.NoError(t, err)

	// 尝试创建不匹配的Chart (ATTRIBUTE参数不能用I_MR图)
	chart := &spc.SpcChart{
		Code:         "TEST_CHART_MISMATCH",
		Name:         "测试控制图",
		ChartType:    "I_MR",
		ParameterID:  param.ID,
		SubgroupSize: 1,
		LimitMethod:  "CALC",
		Ruleset:      "WE1",
		Status:       1,
	}
	err = chartService.CreateSpcChart(chart)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "图表类型与参数数据类型不匹配")

	// 清理
	global.GVA_DB.Unscoped().Delete(param)
}

func TestSubgroupSizeValidation(t *testing.T) {
	// 清理
	global.GVA_DB.Where("code = ?", "TEST_PARAM_VAR").Delete(&spc.SpcParameter{})
	global.GVA_DB.Where("code = ?", "TEST_CHART_SIZE").Delete(&spc.SpcChart{})

	// 创建VARIABLE类型的Parameter
	param := &spc.SpcParameter{
		Code:          "TEST_PARAM_VAR",
		Name:          "测试变量参数",
		DataType:      "VARIABLE",
		UnitOfMeasure: "nm",
		DecimalPlaces: 2,
		Status:        1,
	}
	err := parameterService.CreateSpcParameter(param)
	assert.NoError(t, err)

	// 创建Chart
	chart := &spc.SpcChart{
		Code:         "TEST_CHART_SIZE",
		Name:         "测试控制图",
		ChartType:    "XBAR_R",
		ParameterID:  param.ID,
		SubgroupSize: 5,
		LimitMethod:  "MANUAL",
		Ruleset:      "WE1",
		Status:       1,
	}
	err = chartService.CreateSpcChart(chart)
	assert.NoError(t, err)

	// 尝试采集错误数量的测量值
	now := time.Now()
	collectReq := &spc.CollectDataRequest{
		ChartCode:  "TEST_CHART_SIZE",
		LotID:      "LOT001",
		Values:     []float64{100.0, 101.0, 102.0}, // 只有3个值，期望5个
		SampleTime: &now,
	}
	_, err = collectService.CollectData(collectReq)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "测量值数量必须等于子组大小")

	// 清理
	global.GVA_DB.Unscoped().Delete(chart)
	global.GVA_DB.Unscoped().Delete(param)
}

func TestUniqueSubgroupNo(t *testing.T) {
	// 这个测试需要先有Chart和Sample数据，这里简化为逻辑验证
	// 实际业务中，同一个Chart下的subgroup_no应该唯一
	// 如果再次提交相同的subgroup_no，应该拒绝
	t.Skip("需要完整的数据采集流程，暂时跳过")
}
