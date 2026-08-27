package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcArea 区域/Area
type SpcArea struct {
	global.GVA_MODEL
	SiteID uint   `json:"siteId" gorm:"not null;index:idx_spc_area_site;comment:厂区ID"`
	Code   string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_area_code;comment:区域代码"`
	Name   string `json:"name" gorm:"type:varchar(200);not null;comment:区域名称"`
	Status int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Site *SpcSite `json:"site,omitempty" gorm:"foreignKey:SiteID"`
}

func (SpcArea) TableName() string {
	return "spc_area"
}
