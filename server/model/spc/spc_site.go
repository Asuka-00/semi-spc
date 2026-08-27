package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcSite 厂区/Fab
type SpcSite struct {
	global.GVA_MODEL
	Code     string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_site_code;comment:厂区代码"`
	Name     string `json:"name" gorm:"type:varchar(200);not null;comment:厂区名称"`
	Timezone string `json:"timezone" gorm:"type:varchar(50);default:'Asia/Shanghai';comment:时区"`
	Status   int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark   string `json:"remark" gorm:"type:varchar(500);comment:备注"`
}

func (SpcSite) TableName() string {
	return "spc_site"
}
