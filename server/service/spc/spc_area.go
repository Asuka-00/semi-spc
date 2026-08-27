package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type AreaService struct{}

func (s *AreaService) CreateSpcArea(area *spc.SpcArea) error {
	return global.GVA_DB.Create(area).Error
}

func (s *AreaService) DeleteSpcArea(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcArea{}, id).Error
}

func (s *AreaService) UpdateSpcArea(area *spc.SpcArea) error {
	return global.GVA_DB.Save(area).Error
}

func (s *AreaService) GetSpcArea(id uint) (area spc.SpcArea, err error) {
	err = global.GVA_DB.Preload("Site").Where("id = ?", id).First(&area).Error
	return
}

func (s *AreaService) GetSpcAreaList(info request.PageInfo, siteID uint) (list []spc.SpcArea, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcArea{}).Preload("Site")
	
	if siteID > 0 {
		db = db.Where("site_id = ?", siteID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}
