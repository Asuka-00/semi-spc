package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type ChamberService struct{}

func (s *ChamberService) CreateSpcChamber(chamber *spc.SpcChamber) error {
	return global.GVA_DB.Create(chamber).Error
}

func (s *ChamberService) DeleteSpcChamber(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcChamber{}, id).Error
}

func (s *ChamberService) UpdateSpcChamber(chamber *spc.SpcChamber) error {
	return global.GVA_DB.Save(chamber).Error
}

func (s *ChamberService) GetSpcChamber(id uint) (chamber spc.SpcChamber, err error) {
	err = global.GVA_DB.Preload("Equipment").Where("id = ?", id).First(&chamber).Error
	return
}

func (s *ChamberService) GetSpcChamberList(info request.PageInfo, equipmentID uint) (list []spc.SpcChamber, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcChamber{}).Preload("Equipment")
	
	if equipmentID > 0 {
		db = db.Where("equipment_id = ?", equipmentID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type TechnologyService struct{}

func (s *TechnologyService) CreateSpcTechnology(tech *spc.SpcTechnology) error {
	return global.GVA_DB.Create(tech).Error
}

func (s *TechnologyService) DeleteSpcTechnology(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcTechnology{}, id).Error
}

func (s *TechnologyService) UpdateSpcTechnology(tech *spc.SpcTechnology) error {
	return global.GVA_DB.Save(tech).Error
}

func (s *TechnologyService) GetSpcTechnology(id uint) (tech spc.SpcTechnology, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tech).Error
	return
}

func (s *TechnologyService) GetSpcTechnologyList(info request.PageInfo) (list []spc.SpcTechnology, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcTechnology{})
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type ProductService struct{}

func (s *ProductService) CreateSpcProduct(product *spc.SpcProduct) error {
	return global.GVA_DB.Create(product).Error
}

func (s *ProductService) DeleteSpcProduct(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcProduct{}, id).Error
}

func (s *ProductService) UpdateSpcProduct(product *spc.SpcProduct) error {
	return global.GVA_DB.Save(product).Error
}

func (s *ProductService) GetSpcProduct(id uint) (product spc.SpcProduct, err error) {
	err = global.GVA_DB.Preload("Technology").Where("id = ?", id).First(&product).Error
	return
}

func (s *ProductService) GetSpcProductList(info request.PageInfo, techID uint) (list []spc.SpcProduct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcProduct{}).Preload("Technology")
	
	if techID > 0 {
		db = db.Where("technology_id = ?", techID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type ProcessStepService struct{}

func (s *ProcessStepService) CreateSpcProcessStep(step *spc.SpcProcessStep) error {
	return global.GVA_DB.Create(step).Error
}

func (s *ProcessStepService) DeleteSpcProcessStep(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcProcessStep{}, id).Error
}

func (s *ProcessStepService) UpdateSpcProcessStep(step *spc.SpcProcessStep) error {
	return global.GVA_DB.Save(step).Error
}

func (s *ProcessStepService) GetSpcProcessStep(id uint) (step spc.SpcProcessStep, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&step).Error
	return
}

func (s *ProcessStepService) GetSpcProcessStepList(info request.PageInfo) (list []spc.SpcProcessStep, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcProcessStep{})
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type RecipeService struct{}

func (s *RecipeService) CreateSpcRecipe(recipe *spc.SpcRecipe) error {
	return global.GVA_DB.Create(recipe).Error
}

func (s *RecipeService) DeleteSpcRecipe(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcRecipe{}, id).Error
}

func (s *RecipeService) UpdateSpcRecipe(recipe *spc.SpcRecipe) error {
	return global.GVA_DB.Save(recipe).Error
}

func (s *RecipeService) GetSpcRecipe(id uint) (recipe spc.SpcRecipe, err error) {
	err = global.GVA_DB.Preload("Equipment").Preload("ProcessStep").Where("id = ?", id).First(&recipe).Error
	return
}

func (s *RecipeService) GetSpcRecipeList(info request.PageInfo, equipmentID, stepID uint) (list []spc.SpcRecipe, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcRecipe{}).Preload("Equipment").Preload("ProcessStep")
	
	if equipmentID > 0 {
		db = db.Where("equipment_id = ?", equipmentID)
	}
	if stepID > 0 {
		db = db.Where("process_step_id = ?", stepID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}
