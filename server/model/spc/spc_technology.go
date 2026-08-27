package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcTechnology 工艺技术节点/Technology
type SpcTechnology struct {
	global.GVA_MODEL
	Code   string  `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_technology_code;comment:技术节点代码"`
	Name   string  `json:"name" gorm:"type:varchar(200);not null;comment:技术节点名称"`
	NodeNm float64 `json:"nodeNm" gorm:"type:decimal(10,2);comment:工艺节点(纳米) e.g. 28.00"`
	Status int8    `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark string  `json:"remark" gorm:"type:varchar(500);comment:备注"`
}

func (SpcTechnology) TableName() string {
	return "spc_technology"
}
