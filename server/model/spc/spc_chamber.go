package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcChamber 腔室/Chamber
type SpcChamber struct {
	global.GVA_MODEL
	EquipmentID uint   `json:"equipmentId" gorm:"not null;index:idx_spc_chamber_equipment;comment:设备ID"`
	Code        string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_chamber_code;comment:腔室代码"`
	Name        string `json:"name" gorm:"type:varchar(200);not null;comment:腔室名称"`
	Status      int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark      string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Equipment *SpcEquipment `json:"equipment,omitempty" gorm:"foreignKey:EquipmentID"`
}

func (SpcChamber) TableName() string {
	return "spc_chamber"
}
