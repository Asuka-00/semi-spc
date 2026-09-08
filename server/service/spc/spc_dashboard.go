package spc

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type DashboardService struct{}

// ChartRuntime 控制图运行时数据
type ChartRuntime struct {
	Chart         spc.SpcChart         `json:"chart"`
	Spec          *spc.SpcSpec         `json:"spec"`
	ControlLimit  *spc.SpcControlLimit `json:"controlLimit"`
	Samples       []spc.SpcSample      `json:"samples"`
	Measurements  []spc.SpcMeasurement `json:"measurements,omitempty"`
}

// GetDashboardOverview 获取仪表盘总览
func (s *DashboardService) GetDashboardOverview(days int) (map[string]interface{}, error) {
	if days <= 0 {
		days = 7
	}
	stats := make(map[string]interface{})
	now := time.Now()
	from := now.AddDate(0, 0, -days)
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	var todayAlarms int64
	if err := global.GVA_DB.Model(&spc.SpcAlarm{}).Where("created_at >= ?", todayStart).Count(&todayAlarms).Error; err != nil {
		return nil, err
	}

	var openAlarms int64
	if err := global.GVA_DB.Model(&spc.SpcAlarm{}).Where("status IN ?", []string{"OPEN", "ACK"}).Count(&openAlarms).Error; err != nil {
		return nil, err
	}

	var totalSamples int64
	var oocSamples int64
	_ = global.GVA_DB.Model(&spc.SpcSample{}).Where("sample_time >= ?", from).Count(&totalSamples).Error
	_ = global.GVA_DB.Model(&spc.SpcSample{}).Where("sample_time >= ? AND ooc_flag = ?", from, true).Count(&oocSamples).Error
	oocRate := 0.0
	if totalSamples > 0 {
		oocRate = float64(oocSamples) / float64(totalSamples) * 100
	}

	var activeCharts int64
	_ = global.GVA_DB.Model(&spc.SpcChart{}).Where("status = ?", 1).Count(&activeCharts).Error

	var totalAlarms int64
	_ = global.GVA_DB.Model(&spc.SpcAlarm{}).Where("created_at >= ?", from).Count(&totalAlarms).Error

	alarmService := &AlarmService{}
	detail, err := alarmService.GetAlarmStatistics(days)
	if err != nil {
		return nil, err
	}

	trend := make([]map[string]interface{}, 0, days)
	for i := days - 1; i >= 0; i-- {
		day := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -i)
		next := day.AddDate(0, 0, 1)
		var count int64
		_ = global.GVA_DB.Model(&spc.SpcAlarm{}).Where("created_at >= ? AND created_at < ?", day, next).Count(&count).Error
		trend = append(trend, map[string]interface{}{
			"day":   day.Format("01-02"),
			"count": count,
		})
	}

	type EquipStat struct {
		EquipmentID   uint   `json:"equipmentId"`
		EquipmentName string `json:"equipmentName"`
		EquipmentCode string `json:"equipmentCode"`
		AlarmCount    int64  `json:"alarmCount"`
		OocCount      int64  `json:"oocCount"`
		OosCount      int64  `json:"oosCount"`
	}
	topEquipment := make([]EquipStat, 0)
	_ = global.GVA_DB.Table("spc_alarm AS a").
		Select("e.id AS equipment_id, e.name AS equipment_name, e.code AS equipment_code, COUNT(*) AS alarm_count, SUM(CASE WHEN a.alarm_type = 'OOC' THEN 1 ELSE 0 END) AS ooc_count, SUM(CASE WHEN a.alarm_type = 'OOS' THEN 1 ELSE 0 END) AS oos_count").
		Joins("JOIN spc_sample AS s ON s.id = a.sample_id AND s.deleted_at IS NULL").
		Joins("JOIN spc_equipment AS e ON e.id = s.equipment_id AND e.deleted_at IS NULL").
		Where("a.deleted_at IS NULL AND a.created_at >= ?", from).
		Group("e.id, e.name, e.code").
		Order("alarm_count DESC").
		Limit(8).
		Scan(&topEquipment).Error

	var recentAlarms []spc.SpcAlarm
	_ = global.GVA_DB.Preload("Chart").Preload("Sample").Order("created_at DESC").Limit(8).Find(&recentAlarms).Error

	stats["todayAlarms"] = todayAlarms
	stats["openAlarms"] = openAlarms
	stats["oocRate"] = oocRate
	stats["activeCharts"] = activeCharts
	stats["totalAlarms"] = totalAlarms
	stats["totalSamples"] = totalSamples
	stats["oocSamples"] = oocSamples
	stats["trend"] = trend
	stats["topEquipment"] = topEquipment
	stats["recentAlarms"] = recentAlarms
	stats["byType"] = detail["byType"]
	stats["byStatus"] = detail["byStatus"]
	stats["bySeverity"] = detail["bySeverity"]
	return stats, nil
}

// GetChartRuntime 获取控制图运行时数据（图配置、规格、控制限、样本）
func (s *DashboardService) GetChartRuntime(chartID uint, limit int) (*ChartRuntime, error) {
	if limit <= 0 {
		limit = 50
	}
	chartService := &ChartService{}
	chart, err := chartService.GetSpcChart(chartID)
	if err != nil {
		return nil, err
	}

	runtime := &ChartRuntime{Chart: chart}
	if chart.Spec != nil {
		runtime.Spec = chart.Spec
	} else {
		var spec spc.SpcSpec
		if e := global.GVA_DB.Where("id = ?", chart.SpecID).First(&spec).Error; e == nil {
			runtime.Spec = &spec
		}
	}

	limitRow, err := chartService.GetCurrentControlLimit(chartID)
	if err == nil {
		runtime.ControlLimit = &limitRow
	}

	var samples []spc.SpcSample
	err = global.GVA_DB.Where("chart_id = ?", chartID).
		Order("sample_time ASC, subgroup_no ASC").
		Limit(limit).
		Find(&samples).Error
	if err != nil {
		return nil, err
	}
	runtime.Samples = samples
	return runtime, nil
}
