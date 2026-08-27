package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// SpcCapability 过程能力分析/Capability Analysis
type SpcCapability struct {
	global.GVA_MODEL
	ChartID    uint       `json:"chartId" gorm:"not null;index:idx_spc_capability_chart;comment:控制图ID"`
	WindowFrom *time.Time `json:"windowFrom" gorm:"comment:分析窗口开始时间"`
	WindowTo   *time.Time `json:"windowTo" gorm:"comment:分析窗口结束时间"`
	N          int        `json:"n" gorm:"comment:样本数"`
	Cp         *float64   `json:"cp" gorm:"type:decimal(8,4);comment:短期能力指数 Cp"`
	Cpk        *float64   `json:"cpk" gorm:"type:decimal(8,4);comment:短期过程能力 Cpk"`
	Pp         *float64   `json:"pp" gorm:"type:decimal(8,4);comment:长期能力指数 Pp"`
	Ppk        *float64   `json:"ppk" gorm:"type:decimal(8,4);comment:长期过程能力 Ppk"`
	MeanVal    *float64   `json:"meanVal" gorm:"type:decimal(15,6);comment:均值 μ"`
	StdVal     *float64   `json:"stdVal" gorm:"type:decimal(15,6);comment:总体标准差 σ"`
	Remark     string     `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Chart *SpcChart `json:"chart,omitempty" gorm:"foreignKey:ChartID"`
}

func (SpcCapability) TableName() string {
	return "spc_capability"
}
