package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// SpcOcapExecution OCAP执行记录/OCAP Execution
type SpcOcapExecution struct {
	global.GVA_MODEL
	AlarmID   uint       `json:"alarmId" gorm:"not null;index:idx_spc_ocap_execution_alarm;comment:告警ID"`
	OcapID    uint       `json:"ocapId" gorm:"not null;index:idx_spc_ocap_execution_ocap;comment:OCAP方案ID"`
	Status    string     `json:"status" gorm:"type:varchar(20);default:'PENDING';comment:执行状态 PENDING/IN_PROGRESS/COMPLETED/CANCELLED"`
	Owner     string     `json:"owner" gorm:"type:varchar(100);comment:负责人"`
	StartedAt *time.Time `json:"startedAt" gorm:"comment:开始时间"`
	ClosedAt  *time.Time `json:"closedAt" gorm:"comment:关闭时间"`
	Comment   string     `json:"comment" gorm:"type:text;comment:处理备注"`
	Remark    string     `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Alarm *SpcAlarm `json:"alarm,omitempty" gorm:"foreignKey:AlarmID"`
	Ocap  *SpcOcap  `json:"ocap,omitempty" gorm:"foreignKey:OcapID"`
}

func (SpcOcapExecution) TableName() string {
	return "spc_ocap_execution"
}
