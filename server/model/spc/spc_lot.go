package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcLot 批次/Lot
type SpcLot struct {
	global.GVA_MODEL
	SiteID    uint   `json:"siteId" gorm:"not null;index:idx_spc_lot_site;comment:厂区ID"`
	ProductID uint   `json:"productId" gorm:"not null;index:idx_spc_lot_product;comment:产品ID"`
	LotID     string `json:"lotId" gorm:"type:varchar(100);not null;uniqueIndex:idx_spc_lot_lotid;comment:批次号"`
	LotType   string `json:"lotType" gorm:"type:varchar(20);comment:批次类型 PROD/ENG/PILOT"`
	Qty       int    `json:"qty" gorm:"default:25;comment:片数"`
	Status    int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark    string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Site    *SpcSite    `json:"site,omitempty" gorm:"foreignKey:SiteID"`
	Product *SpcProduct `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (SpcLot) TableName() string {
	return "spc_lot"
}
