package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type LotService struct{}

func (s *LotService) CreateSpcLot(lot *spc.SpcLot) error {
	return global.GVA_DB.Create(lot).Error
}

func (s *LotService) DeleteSpcLot(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcLot{}, id).Error
}

func (s *LotService) UpdateSpcLot(lot *spc.SpcLot) error {
	return global.GVA_DB.Save(lot).Error
}

func (s *LotService) GetSpcLot(id uint) (lot spc.SpcLot, err error) {
	err = global.GVA_DB.Preload("Site").Preload("Product").Where("id = ?", id).First(&lot).Error
	return
}

func (s *LotService) GetSpcLotList(info request.PageInfo, siteID, productID uint) (list []spc.SpcLot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcLot{}).Preload("Site").Preload("Product")
	
	if siteID > 0 {
		db = db.Where("site_id = ?", siteID)
	}
	if productID > 0 {
		db = db.Where("product_id = ?", productID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type WaferService struct{}

func (s *WaferService) CreateSpcWafer(wafer *spc.SpcWafer) error {
	return global.GVA_DB.Create(wafer).Error
}

func (s *WaferService) DeleteSpcWafer(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcWafer{}, id).Error
}

func (s *WaferService) UpdateSpcWafer(wafer *spc.SpcWafer) error {
	return global.GVA_DB.Save(wafer).Error
}

func (s *WaferService) GetSpcWafer(id uint) (wafer spc.SpcWafer, err error) {
	err = global.GVA_DB.Preload("Lot").Where("id = ?", id).First(&wafer).Error
	return
}

func (s *WaferService) GetSpcWaferList(info request.PageInfo, lotID uint) (list []spc.SpcWafer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcWafer{}).Preload("Lot")
	
	if lotID > 0 {
		db = db.Where("lot_id = ?", lotID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type ParameterService struct{}

func (s *ParameterService) CreateSpcParameter(param *spc.SpcParameter) error {
	return global.GVA_DB.Create(param).Error
}

func (s *ParameterService) DeleteSpcParameter(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcParameter{}, id).Error
}

func (s *ParameterService) UpdateSpcParameter(param *spc.SpcParameter) error {
	return global.GVA_DB.Save(param).Error
}

func (s *ParameterService) GetSpcParameter(id uint) (param spc.SpcParameter, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&param).Error
	return
}

func (s *ParameterService) GetSpcParameterList(info request.PageInfo) (list []spc.SpcParameter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcParameter{})
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type SpecService struct{}

func (s *SpecService) CreateSpcSpec(spec *spc.SpcSpec) error {
	return global.GVA_DB.Create(spec).Error
}

func (s *SpecService) DeleteSpcSpec(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcSpec{}, id).Error
}

func (s *SpecService) UpdateSpcSpec(spec *spc.SpcSpec) error {
	return global.GVA_DB.Save(spec).Error
}

func (s *SpecService) GetSpcSpec(id uint) (spec spc.SpcSpec, err error) {
	err = global.GVA_DB.Preload("Parameter").Preload("Product").Preload("ProcessStep").Preload("Equipment").
		Where("id = ?", id).First(&spec).Error
	return
}

func (s *SpecService) GetSpcSpecList(info request.PageInfo, parameterID, productID uint) (list []spc.SpcSpec, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcSpec{}).
		Preload("Parameter").Preload("Product").Preload("ProcessStep").Preload("Equipment")
	
	if parameterID > 0 {
		db = db.Where("parameter_id = ?", parameterID)
	}
	if productID > 0 {
		db = db.Where("product_id = ?", productID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}
