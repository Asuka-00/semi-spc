package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type EquipmentService struct{}

func (s *EquipmentService) CreateSpcEquipment(equipment *spc.SpcEquipment) error {
	return global.GVA_DB.Create(equipment).Error
}

func (s *EquipmentService) DeleteSpcEquipment(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcEquipment{}, id).Error
}

func (s *EquipmentService) UpdateSpcEquipment(equipment *spc.SpcEquipment) error {
	return global.GVA_DB.Save(equipment).Error
}

func (s *EquipmentService) GetSpcEquipment(id uint) (equipment spc.SpcEquipment, err error) {
	err = global.GVA_DB.Preload("Site").Preload("Area").Where("id = ?", id).First(&equipment).Error
	return
}

func (s *EquipmentService) GetSpcEquipmentList(info request.PageInfo, siteID, areaID uint) (list []spc.SpcEquipment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcEquipment{}).Preload("Site").Preload("Area")
	
	if siteID > 0 {
		db = db.Where("site_id = ?", siteID)
	}
	if areaID > 0 {
		db = db.Where("area_id = ?", areaID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}
