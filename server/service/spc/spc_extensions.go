package spc

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/flipped-aurora/gin-vue-admin/server/service/spc/engine"
)

// GetSpcSampleListWithTimeRange 获取指定时间范围的样本列表
func (s *SampleService) GetSpcSampleListWithTimeRange(chartID uint, from, to *time.Time, page, pageSize int) (list []spc.SpcSample, total int64, err error) {
	db := global.GVA_DB.Model(&spc.SpcSample{}).
		Preload("Lot").Preload("Equipment").
		Where("chart_id = ?", chartID)

	if from != nil {
		db = db.Where("sample_time >= ?", from)
	}
	if to != nil {
		db = db.Where("sample_time <= ?", to)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	offset := (page - 1) * pageSize
	err = db.Order("sample_time DESC").Limit(pageSize).Offset(offset).Find(&list).Error
	return
}

// GetAlarmsByChartAndTimeRange 获取指定控制图和时间范围的告警
func (s *AlarmService) GetAlarmsByChartAndTimeRange(chartID uint, from, to *time.Time) (list []spc.SpcAlarm, err error) {
	db := global.GVA_DB.Model(&spc.SpcAlarm{}).Where("chart_id = ?", chartID)

	if from != nil {
		db = db.Where("created_at >= ?", from)
	}
	if to != nil {
		db = db.Where("created_at <= ?", to)
	}

	err = db.Order("created_at DESC").Find(&list).Error
	return
}

// GetCurrentSpec 获取当前有效规格
func (s *SpecService) GetCurrentSpec(parameterID, productID, processStepID uint, equipmentID *uint) (spec spc.SpcSpec, err error) {
	db := global.GVA_DB.Model(&spc.SpcSpec{}).
		Where("parameter_id = ? AND product_id = ? AND process_step_id = ? AND status = 1", parameterID, productID, processStepID)
	if equipmentID != nil && *equipmentID > 0 {
		db = db.Where("equipment_id = ?", *equipmentID)
	}

	err = db.Order("created_at DESC").First(&spec).Error
	return
}

// RecalculateControlLimits 重新计算控制限
func (s *ChartService) RecalculateControlLimits(chartID uint, window int) (result *spc.SpcControlLimit, err error) {
	// 获取控制图信息
	chart, err := s.GetSpcChart(chartID)
	if err != nil {
		return nil, errors.New("控制图不存在")
	}

	// 获取最近N个样本
	var samples []spc.SpcSample
	err = global.GVA_DB.Model(&spc.SpcSample{}).
		Where("chart_id = ?", chartID).
		Order("sample_time DESC").
		Limit(window).
		Find(&samples).Error
	if err != nil {
		return nil, err
	}

	if len(samples) < 10 {
		return nil, errors.New("样本数量不足，至少需要10个样本")
	}

	// 根据图表类型计算控制限
	var ucl, cl, lcl, uclS, clS, lclS *float64

	switch chart.ChartType {
	case "XBAR_R":
		// 计算均值和极差的控制限
		var means, ranges []float64
		for _, sample := range samples {
			if sample.MeanVal != nil {
				means = append(means, *sample.MeanVal)
			}
			if sample.RangeVal != nil {
				ranges = append(ranges, *sample.RangeVal)
			}
		}
		if len(means) == 0 || len(ranges) == 0 {
			return nil, errors.New("样本数据不完整")
		}
		meanBar := engine.CalculateMean(means)
		rangeBar := engine.CalculateMean(ranges)

		ucX, cX, lcX, ucR, cR, lcR := engine.XbarRLimits(meanBar, rangeBar, int(chart.SubgroupSize))
		ucl, cl, lcl, uclS, clS, lclS = &ucX, &cX, &lcX, &ucR, &cR, &lcR

	case "XBAR_S":
		// 计算均值和标准差的控制限
		var means, stddevs []float64
		for _, sample := range samples {
			if sample.MeanVal != nil {
				means = append(means, *sample.MeanVal)
			}
			if sample.StdVal != nil {
				stddevs = append(stddevs, *sample.StdVal)
			}
		}
		if len(means) == 0 || len(stddevs) == 0 {
			return nil, errors.New("样本数据不完整")
		}
		meanBar := engine.CalculateMean(means)
		sBar := engine.CalculateMean(stddevs)

		ucX, cX, lcX, ucS, cS, lcS := engine.XbarSLimits(meanBar, sBar, int(chart.SubgroupSize))
		ucl, cl, lcl, uclS, clS, lclS = &ucX, &cX, &lcX, &ucS, &cS, &lcS

	case "I_MR":
		// 计算单值-移动极差控制限
		var values []float64
		for _, sample := range samples {
			if sample.MeanVal != nil {
				values = append(values, *sample.MeanVal)
			}
		}
		if len(values) == 0 {
			return nil, errors.New("样本数据不完整")
		}
		meanVal := engine.CalculateMean(values)
		mrBar := engine.CalculateMean(engine.CalculateMovingRange(values))

		ucI, cI, lcI, ucMR, cMR, lcMR := engine.IMRLimits(meanVal, mrBar)
		ucl, cl, lcl, uclS, clS, lclS = &ucI, &cI, &lcI, &ucMR, &cMR, &lcMR

	default:
		return nil, errors.New("暂不支持该图表类型的控制限计算")
	}

	now := time.Now()
	// 创建新的控制限记录
	newLimit := &spc.SpcControlLimit{
		ChartID:       chartID,
		Source:        "CALC",
		UCL:           ucl,
		CL:            cl,
		LCL:           lcl,
		UCLS:          uclS,
		CLS:           clS,
		LCLS:          lclS,
		CalcN:         len(samples),
		EffectiveFrom: &now,
		Remark:        fmt.Sprintf("自动计算，窗口大小=%d", window),
	}

	err = global.GVA_DB.Create(newLimit).Error
	if err != nil {
		return nil, err
	}

	return newLimit, nil
}

// CollectCsv CSV批量上传
func (s *CollectService) CollectCsv(file multipart.File) (result map[string]interface{}, err error) {
	reader := csv.NewReader(file)

	// 读取表头
	header, err := reader.Read()
	if err != nil {
		return nil, errors.New("CSV文件格式错误")
	}

	// 验证表头
	expectedHeaders := []string{"chart_id", "lot_id", "wafer_id", "equipment_id", "sample_time", "measurements"}
	if len(header) < len(expectedHeaders) {
		return nil, errors.New("CSV表头不完整")
	}

	successCount := 0
	failCount := 0
	var rowErrors []map[string]interface{}

	rowNum := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			failCount++
			rowErrors = append(rowErrors, map[string]interface{}{
				"row": rowNum,
				"msg": "CSV格式错误",
			})
			continue
		}

		rowNum++

		// 解析数据
		chartID, _ := strconv.ParseUint(record[0], 10, 32)
		lotID, _ := strconv.ParseUint(record[1], 10, 32)
		_ = record[2] // waferID, not used currently
		equipmentID, _ := strconv.ParseUint(record[3], 10, 32)
		sampleTime, parseErr := time.Parse(time.RFC3339, record[4])
		if parseErr != nil {
			failCount++
			rowErrors = append(rowErrors, map[string]interface{}{
				"row": rowNum,
				"msg": "时间格式错误",
			})
			continue
		}

		// 解析测量值（逗号或空格分隔）
		measurementsStr := record[5]
		var measurements []float64
		// 简单分割，假设逗号分隔
		parts := []string{}
		current := ""
		for _, r := range measurementsStr {
			if r == ',' || r == ' ' || r == ';' {
				if current != "" {
					parts = append(parts, current)
					current = ""
				}
			} else {
				current += string(r)
			}
		}
		if current != "" {
			parts = append(parts, current)
		}

		for _, valStr := range parts {
			val, err := strconv.ParseFloat(valStr, 64)
			if err != nil || math.IsNaN(val) {
				failCount++
				rowErrors = append(rowErrors, map[string]interface{}{
					"row": rowNum,
					"msg": "测量值格式错误",
				})
				goto nextRow
			}
			measurements = append(measurements, val)
		}

		// 准备采集请求 - declare before goto
		{
			lotIDStr := strconv.FormatUint(lotID, 10)
			collectReq := CollectDataRequest{
				ChartCode:   strconv.FormatUint(chartID, 10), // 暂用ID作为Code
				LotID:       &lotIDStr,
				EquipmentID: &[]uint{uint(equipmentID)}[0],
				SampleTime:  sampleTime,
				Values:      measurements,
			}

			// 调用采集服务
			_, err = s.CollectData(&collectReq)
			if err != nil {
				failCount++
				rowErrors = append(rowErrors, map[string]interface{}{
					"row": rowNum,
					"msg": err.Error(),
				})
			} else {
				successCount++
			}
		}
	nextRow:
	}

	result = map[string]interface{}{
		"success": successCount,
		"fail":    failCount,
		"errors":  rowErrors,
	}

	return result, nil
}

// GetTodayAlarmCount 获取今日告警数
func (s *AlarmService) GetTodayAlarmCount() (count int64, err error) {
	today := time.Now().Truncate(24 * time.Hour)
	err = global.GVA_DB.Model(&spc.SpcAlarm{}).
		Where("created_at >= ?", today).Count(&count).Error
	return
}

// GetOOCRate 获取OOC率
func (s *AlarmService) GetOOCRate(days int) (rate float64, err error) {
	var totalSamples, oocSamples int64

	// 总样本数
	err = global.GVA_DB.Model(&spc.SpcSample{}).
		Where("created_at >= DATE_SUB(NOW(), INTERVAL ? DAY)", days).
		Count(&totalSamples).Error
	if err != nil {
		return 0, err
	}

	// OOC样本数
	err = global.GVA_DB.Model(&spc.SpcSample{}).
		Where("created_at >= DATE_SUB(NOW(), INTERVAL ? DAY) AND is_ooc = 1", days).
		Count(&oocSamples).Error
	if err != nil {
		return 0, err
	}

	if totalSamples > 0 {
		rate = float64(oocSamples) / float64(totalSamples) * 100
	}

	return rate, nil
}

// GetOpenOcapCount 获取未关闭OCAP数量
func (s *OcapExecutionService) GetOpenOcapCount() (count int64, err error) {
	err = global.GVA_DB.Model(&spc.SpcOcapExecution{}).
		Where("status = ?", "OPEN").Count(&count).Error
	return
}

// GetTopEquipmentByAlarms 获取TOP设备（按告警数排序）
func (s *AlarmService) GetTopEquipmentByAlarms(limit, days int) (result []map[string]interface{}, err error) {
	err = global.GVA_DB.Model(&spc.SpcAlarm{}).
		Select("equipment_id, COUNT(*) as alarm_count").
		Where("created_at >= DATE_SUB(NOW(), INTERVAL ? DAY) AND equipment_id IS NOT NULL", days).
		Group("equipment_id").
		Order("alarm_count DESC").
		Limit(limit).
		Scan(&result).Error
	return
}

// GetTopParameterByAlarms 获取TOP参数（按告警数排序）
func (s *AlarmService) GetTopParameterByAlarms(limit, days int) (result []map[string]interface{}, err error) {
	err = global.GVA_DB.Raw(`
		SELECT c.parameter_id, COUNT(*) as alarm_count
		FROM spc_alarm a
		JOIN spc_chart c ON a.chart_id = c.id
		WHERE a.created_at >= DATE_SUB(NOW(), INTERVAL ? DAY) AND c.parameter_id IS NOT NULL
		GROUP BY c.parameter_id
		ORDER BY alarm_count DESC
		LIMIT ?
	`, days, limit).Scan(&result).Error
	return
}

// HoldSpcLot Hold批次
func (s *LotService) HoldSpcLot(lotID uint, comment string) error {
	return global.GVA_DB.Model(&spc.SpcLot{}).Where("id = ?", lotID).
		Updates(map[string]interface{}{
			"status": "HELD",
			"remark": comment,
		}).Error
}

// ReleaseSpcLot Release批次
func (s *LotService) ReleaseSpcLot(lotID uint, comment string) error {
	return global.GVA_DB.Model(&spc.SpcLot{}).Where("id = ?", lotID).
		Updates(map[string]interface{}{
			"status": "RELEASED",
			"remark": comment,
		}).Error
}

// CalculateAndSaveCapability 计算并保存能力指数
func (s *CapabilityService) CalculateAndSaveCapability(chartID uint, startTime, endTime time.Time) (result map[string]interface{}, err error) {
	// 获取控制图信息
	var chart spc.SpcChart
	err = global.GVA_DB.First(&chart, chartID).Error
	if err != nil {
		return nil, errors.New("控制图不存在")
	}

	// 获取时间范围内的样本
	var samples []spc.SpcSample
	err = global.GVA_DB.Model(&spc.SpcSample{}).
		Where("chart_id = ? AND sample_time BETWEEN ? AND ?", chartID, startTime, endTime).
		Order("sample_time ASC").
		Find(&samples).Error
	if err != nil {
		return nil, err
	}

	if len(samples) < 30 {
		return nil, errors.New("样本数量不足，至少需要30个样本")
	}

	// 获取规格限
	var spec spc.SpcSpec
	err = global.GVA_DB.Model(&spc.SpcSpec{}).
		Where("parameter_id = ? AND status = 1", chart.ParameterID).
		Order("created_at DESC").
		First(&spec).Error
	if err != nil {
		return nil, errors.New("未找到有效规格")
	}

	if spec.USL == nil || spec.LSL == nil || spec.Target == nil {
		return nil, errors.New("规格限不完整")
	}

	// 提取所有测量值
	var allValues []float64
	for _, sample := range samples {
		// 获取该样本的测量值
		var measurements []spc.SpcMeasurement
		global.GVA_DB.Where("sample_id = ?", sample.ID).Find(&measurements)
		for _, m := range measurements {
			if m.Value != nil {
				allValues = append(allValues, *m.Value)
			}
		}
	}

	if len(allValues) == 0 {
		return nil, errors.New("没有测量值数据")
	}

	// 计算能力指数 (使用3倍标准差)
	capResult := engine.CalculateCapability(allValues, spec.USL, spec.LSL, spec.Target, 3.0)

	// 保存能力分析结果
	capability := &spc.SpcCapability{
		ChartID:    chartID,
		WindowFrom: &startTime,
		WindowTo:   &endTime,
		N:          len(allValues),
		Cp:         &capResult.Cp,
		Cpk:        &capResult.Cpk,
		Pp:         &capResult.Pp,
		Ppk:        &capResult.Ppk,
		MeanVal:    &capResult.Mean,
		StdVal:     &capResult.Sigma,
	}

	err = global.GVA_DB.Create(capability).Error
	if err != nil {
		return nil, err
	}

	result = map[string]interface{}{
		"capability_id": capability.ID,
		"cp":            capResult.Cp,
		"cpk":           capResult.Cpk,
		"pp":            capResult.Pp,
		"ppk":           capResult.Ppk,
		"mean":          capResult.Mean,
		"sigma":         capResult.Sigma,
	}

	return result, nil
}
