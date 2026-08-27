package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcEquipment 设备/Equipment
type SpcEquipment struct {
	global.GVA_MODEL
	SiteID  uint   `json:"siteId" gorm:"not null;index:idx_spc_equipment_site;comment:厂区ID"`
	AreaID  uint   `json:"areaId" gorm:"not null;index:idx_spc_equipment_area;comment:区域ID"`
	Code    string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_equipment_code;comment:设备代码"`
	Name    string `json:"name" gorm:"type:varchar(200);not null;comment:设备名称"`
	EqpType string `json:"eqpType" gorm:"type:varchar(50);comment:设备类型 LITHO/ETCH/CVD/PVD/IMP/DIFF/CMP/METROLOGY/OTHER"`
	Vendor  string `json:"vendor" gorm:"type:varchar(100);comment:供应商"`
	Status  int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark  string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Site *SpcSite `json:"site,omitempty" gorm:"foreignKey:SiteID"`
	Area *SpcArea `json:"area,omitempty" gorm:"foreignKey:AreaID"`
}

func (SpcEquipment) TableName() string {
	return "spc_equipment"
}
