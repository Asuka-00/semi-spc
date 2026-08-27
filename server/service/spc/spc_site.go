package spc

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type SiteService struct{}

// CreateSpcSite 创建厂区
func (s *SiteService) CreateSpcSite(site *spc.SpcSite) error {
	// 验证代码唯一性
	var count int64
	err := global.GVA_DB.Model(&spc.SpcSite{}).Where("code = ? AND deleted_at IS NULL", site.Code).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("代码已存在")
	}
	return global.GVA_DB.Create(site).Error
}

// DeleteSpcSite 删除厂区
func (s *SiteService) DeleteSpcSite(id uint) error {
	// 检查是否有关联的Area
	var areaCount int64
	err := global.GVA_DB.Model(&spc.SpcArea{}).Where("site_id = ?", id).Count(&areaCount).Error
	if err != nil {
		return err
	}
	if areaCount > 0 {
		return errors.New("存在关联数据，无法删除")
	}
	return global.GVA_DB.Delete(&spc.SpcSite{}, id).Error
}

// UpdateSpcSite 更新厂区
func (s *SiteService) UpdateSpcSite(site *spc.SpcSite) error {
	return global.GVA_DB.Save(site).Error
}

// GetSpcSite 根据id获取厂区
func (s *SiteService) GetSpcSite(id uint) (site spc.SpcSite, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&site).Error
	return
}

// GetSpcSiteByCode 根据code获取厂区
func (s *SiteService) GetSpcSiteByCode(code string) (site spc.SpcSite, err error) {
	err = global.GVA_DB.Where("code = ?", code).First(&site).Error
	return
}

// GetSpcSiteList 分页获取厂区列表
func (s *SiteService) GetSpcSiteList(info request.PageInfo) (list []spc.SpcSite, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcSite{})
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

// GetAllSpcSites 获取所有厂区（不分页）
func (s *SiteService) GetAllSpcSites() (list []spc.SpcSite, err error) {
	err = global.GVA_DB.Where("status = ?", 1).Find(&list).Error
	return
}
