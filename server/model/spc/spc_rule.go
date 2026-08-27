package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcRule 控制规则/Rule
type SpcRule struct {
	global.GVA_MODEL
	ChartID  uint    `json:"chartId" gorm:"not null;index:idx_spc_rule_chart;comment:控制图ID"`
	RuleCode string  `json:"ruleCode" gorm:"type:varchar(20);not null;comment:规则代码 WE1/WE2/WE3/WE4/NELSON1/NELSON2等"`
	Enabled  bool    `json:"enabled" gorm:"default:true;comment:是否启用"`
	N        int     `json:"n" gorm:"comment:连续点数 n"`
	K        float64 `json:"k" gorm:"type:decimal(5,2);comment:σ倍数 k"`
	Remark   string  `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Chart *SpcChart `json:"chart,omitempty" gorm:"foreignKey:ChartID"`
}

func (SpcRule) TableName() string {
	return "spc_rule"
}
