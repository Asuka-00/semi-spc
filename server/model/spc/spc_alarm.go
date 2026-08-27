package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcAlarm 告警/Alarm
type SpcAlarm struct {
	global.GVA_MODEL
	SampleID  uint   `json:"sampleId" gorm:"not null;index:idx_spc_alarm_sample;comment:样本ID"`
	ChartID   uint   `json:"chartId" gorm:"not null;index:idx_spc_alarm_chart;comment:控制图ID"`
	AlarmType string `json:"alarmType" gorm:"type:varchar(10);not null;comment:告警类型 OOC/OOS"`
	RuleCode  string `json:"ruleCode" gorm:"type:varchar(20);comment:触发规则代码 WE1/NELSON5等"`
	Severity  string `json:"severity" gorm:"type:varchar(10);default:'WARN';comment:严重度 INFO/WARN/CRIT"`
	Status    string `json:"status" gorm:"type:varchar(20);default:'OPEN';comment:状态 OPEN/ACK/CLOSED"`
	HoldLot   bool   `json:"holdLot" gorm:"default:false;comment:是否Hold批次"`
	Remark    string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Sample *SpcSample `json:"sample,omitempty" gorm:"foreignKey:SampleID"`
	Chart  *SpcChart  `json:"chart,omitempty" gorm:"foreignKey:ChartID"`
}

func (SpcAlarm) TableName() string {
	return "spc_alarm"
}
