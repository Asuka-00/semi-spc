package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcProcessStep 工艺步骤/Process Step
type SpcProcessStep struct {
	global.GVA_MODEL
	Code     string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_process_step_code;comment:工艺步骤代码"`
	Name     string `json:"name" gorm:"type:varchar(200);not null;comment:工艺步骤名称"`
	StepType string `json:"stepType" gorm:"type:varchar(50);comment:步骤类型"`
	Status   int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark   string `json:"remark" gorm:"type:varchar(500);comment:备注"`
}

func (SpcProcessStep) TableName() string {
	return "spc_process_step"
}
