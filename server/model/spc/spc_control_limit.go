package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// SpcControlLimit 控制限/Control Limit
type SpcControlLimit struct {
	global.GVA_MODEL
	ChartID       uint       `json:"chartId" gorm:"not null;index:idx_spc_control_limit_chart;comment:控制图ID"`
	UCL           *float64   `json:"ucl" gorm:"type:decimal(15,6);comment:上控制限 Upper Control Limit"`
	CL            *float64   `json:"cl" gorm:"type:decimal(15,6);comment:中心线 Center Line"`
	LCL           *float64   `json:"lcl" gorm:"type:decimal(15,6);comment:下控制限 Lower Control Limit"`
	UCLS          *float64   `json:"uclS" gorm:"type:decimal(15,6);comment:S图上控制限(Xbar-S)或MR图上限(I-MR)"`
	CLS           *float64   `json:"clS" gorm:"type:decimal(15,6);comment:S图中心线"`
	LCLS          *float64   `json:"lclS" gorm:"type:decimal(15,6);comment:S图下控制限"`
	CalcN         int        `json:"calcN" gorm:"comment:计算样本数 n"`
	Source        string     `json:"source" gorm:"type:varchar(20);default:'CALC';comment:来源 CALC/MANUAL"`
	EffectiveFrom *time.Time `json:"effectiveFrom" gorm:"comment:生效开始时间"`
	EffectiveTo   *time.Time `json:"effectiveTo" gorm:"comment:生效结束时间"`
	Remark        string     `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Chart *SpcChart `json:"chart,omitempty" gorm:"foreignKey:ChartID"`
}

func (SpcControlLimit) TableName() string {
	return "spc_control_limit"
}
