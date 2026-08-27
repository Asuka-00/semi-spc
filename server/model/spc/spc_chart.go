package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcChart 控制图/Chart
type SpcChart struct {
	global.GVA_MODEL
	Code          string   `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_chart_code;comment:控制图代码"`
	Name          string   `json:"name" gorm:"type:varchar(200);not null;comment:控制图名称"`
	ParameterID   uint     `json:"parameterId" gorm:"not null;index:idx_spc_chart_parameter;comment:参数ID"`
	SpecID        uint     `json:"specId" gorm:"not null;index:idx_spc_chart_spec;comment:规格ID"`
	ChartType     string   `json:"chartType" gorm:"type:varchar(20);not null;comment:控制图类型 I_MR/XBAR_R/XBAR_S/P/NP/C/U/EWMA/CUSUM"`
	SubgroupSize  int      `json:"subgroupSize" gorm:"default:5;comment:子组大小 n"`
	Ruleset       string   `json:"ruleset" gorm:"type:varchar(100);default:'WE1';comment:规则集 逗号分隔 e.g. WE1,WE2,NELSON5"`
	LimitMethod   string   `json:"limitMethod" gorm:"type:varchar(20);default:'CALC';comment:控制限方法 CALC/MANUAL"`
	EwmaLambda    *float64 `json:"ewmaLambda" gorm:"type:decimal(5,4);comment:EWMA平滑系数 λ (0,1)"`
	CusumK        *float64 `json:"cusumK" gorm:"type:decimal(15,6);comment:CUSUM参考值 K"`
	CusumH        *float64 `json:"cusumH" gorm:"type:decimal(15,6);comment:CUSUM决策区间 H"`
	Status        int8     `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark        string   `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Parameter *SpcParameter `json:"parameter,omitempty" gorm:"foreignKey:ParameterID"`
	Spec      *SpcSpec      `json:"spec,omitempty" gorm:"foreignKey:SpecID"`
}

func (SpcChart) TableName() string {
	return "spc_chart"
}
