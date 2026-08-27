package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type ChartService struct{}

// CreateSpcChart 创建控制图
func (s *ChartService) CreateSpcChart(chart *spc.SpcChart) error {
	return global.GVA_DB.Create(chart).Error
}

// DeleteSpcChart 删除控制图
func (s *ChartService) DeleteSpcChart(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcChart{}, id).Error
}

// UpdateSpcChart 更新控制图
func (s *ChartService) UpdateSpcChart(chart *spc.SpcChart) error {
	return global.GVA_DB.Save(chart).Error
}

// GetSpcChart 根据id获取控制图（含关联数据）
func (s *ChartService) GetSpcChart(id uint) (chart spc.SpcChart, err error) {
	err = global.GVA_DB.Preload("Parameter").Preload("Spec").Preload("Spec.Parameter").
		Where("id = ?", id).First(&chart).Error
	return
}

// GetSpcChartByCode 根据code获取控制图
func (s *ChartService) GetSpcChartByCode(code string) (chart spc.SpcChart, err error) {
	err = global.GVA_DB.Preload("Parameter").Preload("Spec").
		Where("code = ?", code).First(&chart).Error
	return
}

// GetSpcChartList 分页获取控制图列表
func (s *ChartService) GetSpcChartList(info request.PageInfo) (list []spc.SpcChart, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcChart{}).Preload("Parameter").Preload("Spec")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

// GetActiveCharts 获取所有启用的控制图
func (s *ChartService) GetActiveCharts() (list []spc.SpcChart, err error) {
	err = global.GVA_DB.Where("status = ?", 1).
		Preload("Parameter").Preload("Spec").Find(&list).Error
	return
}

// GetCurrentControlLimit 获取当前有效的控制限
func (s *ChartService) GetCurrentControlLimit(chartID uint) (limit spc.SpcControlLimit, err error) {
	err = global.GVA_DB.Where("chart_id = ? AND (effective_to IS NULL OR effective_to > NOW())", chartID).
		Order("created_at DESC").First(&limit).Error
	return
}

// GetActiveRules 获取控制图的启用规则
func (s *ChartService) GetActiveRules(chartID uint) (rules []spc.SpcRule, err error) {
	err = global.GVA_DB.Where("chart_id = ? AND enabled = ?", chartID, true).Find(&rules).Error
	return
}
