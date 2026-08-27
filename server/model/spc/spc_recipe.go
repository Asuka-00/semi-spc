package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcRecipe 配方/Recipe
type SpcRecipe struct {
	global.GVA_MODEL
	EquipmentID   uint   `json:"equipmentId" gorm:"not null;index:idx_spc_recipe_equipment;comment:设备ID"`
	ProcessStepID uint   `json:"processStepId" gorm:"not null;index:idx_spc_recipe_step;comment:工艺步骤ID"`
	Code          string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_recipe_code;comment:配方代码"`
	Name          string `json:"name" gorm:"type:varchar(200);not null;comment:配方名称"`
	Version       string `json:"version" gorm:"type:varchar(50);comment:配方版本"`
	Status        int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark        string `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Equipment   *SpcEquipment   `json:"equipment,omitempty" gorm:"foreignKey:EquipmentID"`
	ProcessStep *SpcProcessStep `json:"processStep,omitempty" gorm:"foreignKey:ProcessStepID"`
}

func (SpcRecipe) TableName() string {
	return "spc_recipe"
}
