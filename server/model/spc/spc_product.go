package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcProduct 产品/Product
type SpcProduct struct {
	global.GVA_MODEL
	TechnologyID uint   `json:"technologyId" gorm:"not null;index:idx_spc_product_technology;comment:技术节点ID"`
	Code         string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_product_code;comment:产品代码"`
	Name         string `json:"name" gorm:"type:varchar(200);not null;comment:产品名称"`
	Status       int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark       string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Technology *SpcTechnology `json:"technology,omitempty" gorm:"foreignKey:TechnologyID"`
}

func (SpcProduct) TableName() string {
	return "spc_product"
}
