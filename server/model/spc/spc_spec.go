package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// SpcSpec 规格/Specification (版本化)
type SpcSpec struct {
	global.GVA_MODEL
	ParameterID   uint       `json:"parameterId" gorm:"not null;index:idx_spc_spec_parameter;comment:参数ID"`
	ProductID     uint       `json:"productId" gorm:"not null;index:idx_spc_spec_product;comment:产品ID"`
	ProcessStepID uint       `json:"processStepId" gorm:"not null;index:idx_spc_spec_step;comment:工艺步骤ID"`
	EquipmentID   *uint      `json:"equipmentId" gorm:"index:idx_spc_spec_equipment;comment:设备ID(nullable=全局规格)"`
	Version       string     `json:"version" gorm:"type:varchar(50);comment:版本号"`
	USL           *float64   `json:"usl" gorm:"type:decimal(15,6);comment:规格上限 Upper Specification Limit"`
	Target        *float64   `json:"target" gorm:"type:decimal(15,6);comment:目标值"`
	LSL           *float64   `json:"lsl" gorm:"type:decimal(15,6);comment:规格下限 Lower Specification Limit"`
	EffectiveFrom *time.Time `json:"effectiveFrom" gorm:"comment:生效开始时间"`
	EffectiveTo   *time.Time `json:"effectiveTo" gorm:"comment:生效结束时间"`
	Status        int8       `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark        string     `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Parameter   *SpcParameter   `json:"parameter,omitempty" gorm:"foreignKey:ParameterID"`
	Product     *SpcProduct     `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	ProcessStep *SpcProcessStep `json:"processStep,omitempty" gorm:"foreignKey:ProcessStepID"`
	Equipment   *SpcEquipment   `json:"equipment,omitempty" gorm:"foreignKey:EquipmentID"`
}

func (SpcSpec) TableName() string {
	return "spc_spec"
}
