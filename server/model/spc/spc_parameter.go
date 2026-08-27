package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SpcParameter 参数/Parameter
type SpcParameter struct {
	global.GVA_MODEL
	Code          string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:idx_spc_parameter_code;comment:参数代码"`
	Name          string `json:"name" gorm:"type:varchar(200);not null;comment:参数名称"`
	DataType      string `json:"dataType" gorm:"type:varchar(20);comment:数据类型 VARIABLE/ATTRIBUTE"`
	Unit          string `json:"unit" gorm:"type:varchar(20);comment:单位 nm/Å/℃等"`
	DecimalPlaces int    `json:"decimalPlaces" gorm:"default:2;comment:小数位数"`
	SampleLevel   string `json:"sampleLevel" gorm:"type:varchar(20);comment:采样级别 LOT/WAFER/SITE"`
	Status        int8   `json:"status" gorm:"type:tinyint;default:1;comment:状态 0=禁用 1=启用"`
	Remark        string `json:"remark" gorm:"type:varchar(500);comment:备注"`
}

func (SpcParameter) TableName() string {
	return "spc_parameter"
}
