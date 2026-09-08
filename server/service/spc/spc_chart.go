package spc

import (
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/flipped-aurora/gin-vue-admin/server/service/spc/engine"
	"gorm.io/gorm"
)

type ChartService struct{}

func (s *ChartService) CreateSpcChart(chart *spc.SpcChart) error {
	err := global.GVA_DB.Create(chart).Error
	if err != nil {
		return err
	}
	defaults := []string{"WE1", "WE2", "WE4"}
	rules := make([]spc.SpcRule, 0, len(defaults))
	for _, code := range defaults {
		def := engine.DefaultRuleConfig(code)
		rules = append(rules, spc.SpcRule{
			ChartID:  chart.ID,
			RuleCode: code,
			Enabled:  code == "WE1",
			N:        def.N,
			Hits:     def.Hits,
			K:        def.K,
			Severity: def.Severity,
		})
	}
	return global.GVA_DB.Create(&rules).Error
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
	info = normalizePage(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcChart{}).Preload("Parameter").Preload("Spec")
	if info.Keyword != "" {
		kw := "%" + info.Keyword + "%"
		db = db.Where("code LIKE ? OR name LIKE ?", kw, kw)
	}

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
	now := time.Now()
	err = global.GVA_DB.Where("chart_id = ? AND (effective_to IS NULL OR effective_to > ?)", chartID, now).
		Order("created_at DESC").First(&limit).Error
	return
}

func f64(v float64) *float64 { return &v }

// CalculateAndSaveControlLimit 按样本重算当前控制限（CALC），并关闭旧有效限
func (s *ChartService) CalculateAndSaveControlLimit(chartID uint, sampleN int) (*spc.SpcControlLimit, error) {
	if sampleN <= 0 {
		sampleN = 50
	}
	chart, err := s.GetSpcChart(chartID)
	if err != nil {
		return nil, err
	}
	var samples []spc.SpcSample
	err = global.GVA_DB.Where("chart_id = ?", chartID).
		Order("sample_time DESC, subgroup_no DESC").
		Limit(sampleN).
		Find(&samples).Error
	if err != nil {
		return nil, err
	}
	for i, j := 0, len(samples)-1; i < j; i, j = i+1, j-1 {
		samples[i], samples[j] = samples[j], samples[i]
	}
	if len(samples) < 2 {
		return nil, errors.New("样本数量不足，至少需要 2 个子组才能计算控制限")
	}

	means := make([]float64, 0, len(samples))
	ranges := make([]float64, 0, len(samples))
	stds := make([]float64, 0, len(samples))
	for _, item := range samples {
		if item.MeanVal != nil {
			means = append(means, *item.MeanVal)
		}
		if item.RangeVal != nil {
			ranges = append(ranges, *item.RangeVal)
		}
		if item.StdVal != nil {
			stds = append(stds, *item.StdVal)
		}
	}
	if len(means) < 2 {
		return nil, errors.New("有效均值样本不足")
	}

	var ucl, cl, lcl, uclS, clS, lclS float64
	n := chart.SubgroupSize
	if n < 1 {
		n = 1
	}
	switch chart.ChartType {
	case "XBAR_S":
		if len(stds) < 2 {
			return nil, errors.New("标准差样本不足，无法计算 X̄-S 控制限")
		}
		ucl, cl, lcl, uclS, clS, lclS = engine.XbarSLimits(engine.CalculateMean(means), engine.CalculateMean(stds), n)
	case "I_MR", "EWMA", "CUSUM":
		mrs := engine.CalculateMovingRange(means)
		if len(mrs) == 0 {
			return nil, errors.New("无法计算移动极差")
		}
		ucl, cl, lcl, uclS, clS, lclS = engine.IMRLimits(engine.CalculateMean(means), engine.CalculateMean(mrs))
	case "XBAR_R":
		if len(ranges) < 2 {
			return nil, errors.New("极差样本不足，无法计算 X̄-R 控制限")
		}
		ucl, cl, lcl, uclS, clS, lclS = engine.XbarRLimits(engine.CalculateMean(means), engine.CalculateMean(ranges), n)
	default:
		sigma := engine.CalculateStdDev(means)
		cl = engine.CalculateMean(means)
		ucl = cl + 3*sigma
		lcl = cl - 3*sigma
	}

	now := time.Now()
	if err = global.GVA_DB.Model(&spc.SpcControlLimit{}).
		Where("chart_id = ? AND effective_to IS NULL", chartID).
		Update("effective_to", now).Error; err != nil {
		return nil, err
	}

	limit := &spc.SpcControlLimit{
		ChartID:       chartID,
		UCL:           f64(ucl),
		CL:            f64(cl),
		LCL:           f64(lcl),
		CalcN:         len(means),
		Source:        "CALC",
		EffectiveFrom: &now,
	}
	if uclS != 0 || clS != 0 || lclS != 0 {
		limit.UCLS = f64(uclS)
		limit.CLS = f64(clS)
		limit.LCLS = f64(lclS)
	}
	if err = global.GVA_DB.Create(limit).Error; err != nil {
		return nil, err
	}
	return limit, nil
}

// SaveChartRules 覆盖保存控制图规则
func (s *ChartService) SaveChartRules(chartID uint, rules []spc.SpcRule) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("chart_id = ?", chartID).Delete(&spc.SpcRule{}).Error; err != nil {
			return err
		}
		if len(rules) == 0 {
			return nil
		}
		for i := range rules {
			rules[i].ID = 0
			rules[i].ChartID = chartID
			if rules[i].Severity == "" {
				rules[i].Severity = "WARN"
			}
		}
		return tx.Create(&rules).Error
	})
}

// GetActiveRules 获取控制图的启用规则
func (s *ChartService) GetActiveRules(chartID uint) (rules []spc.SpcRule, err error) {
	err = global.GVA_DB.Where("chart_id = ? AND enabled = ?", chartID, true).Find(&rules).Error
	return
}
