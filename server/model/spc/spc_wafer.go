package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcWafer 晶圆/Wafer
type SpcWafer struct {
	global.GVA_MODEL
	LotID   uint   `json:"lotId" gorm:"not null;index:idx_spc_wafer_lot;comment:批次ID"`
	SlotNo  int    `json:"slotNo" gorm:"not null;comment:槽位号 1-25"`
	WaferID string `json:"waferId" gorm:"type:varchar(100);not null;uniqueIndex:idx_spc_wafer_waferid;comment:晶圆ID"`
	Status  int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark  string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Lot *SpcLot `json:"lot,omitempty" gorm:"foreignKey:LotID"`
}

func (SpcWafer) TableName() string {
	return "spc_wafer"
}
