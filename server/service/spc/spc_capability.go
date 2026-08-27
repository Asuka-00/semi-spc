package spc

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/flipped-aurora/gin-vue-admin/server/service/spc/engine"
)

type CapabilityService struct{}

// CalculateCapabilityRequest 计算能力请求
type CalculateCapabilityRequest struct {
	ChartID    uint      `json:"chartId" binding:"required"`
	WindowFrom time.Time `json:"windowFrom" binding:"required"`
	WindowTo   time.Time `json:"windowTo" binding:"required"`
}

// CalculateCapability 计算过程能力
func (s *CapabilityService) CalculateCapability(req *CalculateCapabilityRequest) (*spc.SpcCapability, error) {
	chartService := &ChartService{}

	// 获取控制图配置
	chart, err := chartService.GetSpcChart(req.ChartID)
	if err != nil {
		return nil, err
	}

	// 获取规格
	var spec spc.SpcSpec
	err = global.GVA_DB.Where("id = ?", chart.SpecID).First(&spec).Error
	if err != nil {
		return nil, err
	}

	// 获取时间窗口内的样本数据
	var samples []spc.SpcSample
	err = global.GVA_DB.Where("chart_id = ? AND sample_time >= ? AND sample_time <= ?",
		req.ChartID, req.WindowFrom, req.WindowTo).
		Order("sample_time ASC").Find(&samples).Error
	if err != nil {
		return nil, err
	}

	if len(samples) < 2 {
		return nil, err
	}

	// 提取均值和极差/标准差
	means := make([]float64, len(samples))
	ranges := make([]float64, len(samples))
	stds := make([]float64, len(samples))

	for i, s := range samples {
		if s.MeanVal != nil {
			means[i] = *s.MeanVal
		}
		if s.RangeVal != nil {
			ranges[i] = *s.RangeVal
		}
		if s.StdVal != nil {
			stds[i] = *s.StdVal
		}
	}

	// 使用计算引擎计算能力指数
	var result *engine.CapabilityResult
	if chart.ChartType == "XBAR_R" {
		result = engine.CalculateCapabilityFromSubgroups(means, ranges, nil, chart.SubgroupSize, spec.USL, spec.LSL, spec.Target)
	} else if chart.ChartType == "XBAR_S" {
		result = engine.CalculateCapabilityFromSubgroups(means, nil, stds, chart.SubgroupSize, spec.USL, spec.LSL, spec.Target)
	} else {
		// I-MR或其他类型
		result = engine.CalculateCapability(means, spec.USL, spec.LSL, spec.Target, 0)
	}

	if result == nil {
		return nil, err
	}

	// 保存能力分析结果
	capability := &spc.SpcCapability{
		ChartID:    req.ChartID,
		WindowFrom: &req.WindowFrom,
		WindowTo:   &req.WindowTo,
		N:          result.N,
		Cp:         &result.Cp,
		Cpk:        &result.Cpk,
		Pp:         &result.Pp,
		Ppk:        &result.Ppk,
		MeanVal:    &result.Mean,
		StdVal:     &result.Sigma,
	}

	err = global.GVA_DB.Create(capability).Error
	if err != nil {
		return nil, err
	}

	return capability, nil
}

// GetCapabilityHistory 获取能力分析历史
func (s *CapabilityService) GetCapabilityHistory(chartID uint, limit int) (list []spc.SpcCapability, err error) {
	err = global.GVA_DB.Where("chart_id = ?", chartID).
		Order("created_at DESC").Limit(limit).Find(&list).Error
	return
}
