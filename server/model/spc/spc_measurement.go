package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcMeasurement 测量值/Measurement
type SpcMeasurement struct {
	global.GVA_MODEL
	SampleID    uint     `json:"sampleId" gorm:"not null;index:idx_spc_measurement_sample;comment:样本ID"`
	SeqNo       int      `json:"seqNo" gorm:"comment:序号 1-n"`
	SiteX       *int     `json:"siteX" gorm:"comment:测量站点X坐标"`
	SiteY       *int     `json:"siteY" gorm:"comment:测量站点Y坐标"`
	Value       *float64 `json:"value" gorm:"type:decimal(15,6);comment:测量值(变量型)"`
	DefectCount *int     `json:"defectCount" gorm:"comment:缺陷数(计数型)"`
	Remark      string   `json:"remark" gorm:"type:varchar(500);comment:备注"`

	Sample *SpcSample `json:"sample,omitempty" gorm:"foreignKey:SampleID"`
}

func (SpcMeasurement) TableName() string {
	return "spc_measurement"
}
