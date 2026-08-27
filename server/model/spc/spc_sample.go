package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// SpcSample 样本/子组/Sample
type SpcSample struct {
	global.GVA_MODEL
	ChartID     uint       `json:"chartId" gorm:"not null;index:idx_spc_sample_chart;comment:控制图ID"`
	LotID       *uint      `json:"lotId" gorm:"index:idx_spc_sample_lot;comment:批次ID"`
	WaferID     *uint      `json:"waferId" gorm:"index:idx_spc_sample_wafer;comment:晶圆ID"`
	EquipmentID *uint      `json:"equipmentId" gorm:"index:idx_spc_sample_equipment;comment:设备ID"`
	ChamberID   *uint      `json:"chamberId" gorm:"index:idx_spc_sample_chamber;comment:腔室ID"`
	RecipeID    *uint      `json:"recipeId" gorm:"index:idx_spc_sample_recipe;comment:配方ID"`
	SampleTime  *time.Time `json:"sampleTime" gorm:"index:idx_spc_sample_time;comment:采样时间"`
	SubgroupNo  int        `json:"subgroupNo" gorm:"comment:子组号"`
	N           int        `json:"n" gorm:"comment:实际测量点数"`
	MeanVal     *float64   `json:"meanVal" gorm:"type:decimal(15,6);comment:均值 X̄"`
	RangeVal    *float64   `json:"rangeVal" gorm:"type:decimal(15,6);comment:极差 R"`
	StdVal      *float64   `json:"stdVal" gorm:"type:decimal(15,6);comment:标准差 S"`
	OocFlag     bool       `json:"oocFlag" gorm:"default:false;comment:失控标志 Out Of Control"`
	OosFlag     bool       `json:"oosFlag" gorm:"default:false;comment:超规格标志 Out Of Spec"`
	Remark      string     `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Chart     *SpcChart     `json:"chart,omitempty" gorm:"foreignKey:ChartID"`
	Lot       *SpcLot       `json:"lot,omitempty" gorm:"foreignKey:LotID"`
	Wafer     *SpcWafer     `json:"wafer,omitempty" gorm:"foreignKey:WaferID"`
	Equipment *SpcEquipment `json:"equipment,omitempty" gorm:"foreignKey:EquipmentID"`
	Chamber   *SpcChamber   `json:"chamber,omitempty" gorm:"foreignKey:ChamberID"`
	Recipe    *SpcRecipe    `json:"recipe,omitempty" gorm:"foreignKey:RecipeID"`
}

func (SpcSample) TableName() string {
	return "spc_sample"
}
