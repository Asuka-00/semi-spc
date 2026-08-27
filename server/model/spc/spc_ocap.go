package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcOcap OCAP行动方案/Out-of-Control Action Plan
type SpcOcap struct {
	global.GVA_MODEL
	ChartID     uint   `json:"chartId" gorm:"not null;index:idx_spc_ocap_chart;comment:控制图ID"`
	Name        string `json:"name" gorm:"type:varchar(200);not null;comment:OCAP名称"`
	TriggerType string `json:"triggerType" gorm:"type:varchar(50);comment:触发类型 OOC/OOS/BOTH"`
	StepsJson   string `json:"stepsJson" gorm:"type:text;comment:步骤定义JSON"`
	Status      int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark      string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Chart *SpcChart `json:"chart,omitempty" gorm:"foreignKey:ChartID"`
}

func (SpcOcap) TableName() string {
	return "spc_ocap"
}
