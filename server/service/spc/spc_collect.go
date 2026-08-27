package spc

import (
	"errors"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/flipped-aurora/gin-vue-admin/server/service/spc/engine"
	"gorm.io/gorm"
)

type CollectService struct{}

// CollectDataRequest 数据采集请求
type CollectDataRequest struct {
	ChartCode   string    `json:"chartCode" binding:"required"`
	LotID       *string   `json:"lotId"`
	WaferID     *string   `json:"waferId"`
	EquipmentID *uint     `json:"equipmentId"`
	ChamberID   *uint     `json:"chamberId"`
	RecipeID    *uint     `json:"recipeId"`
	SampleTime  time.Time `json:"sampleTime"`
	SubgroupNo  int       `json:"subgroupNo"`
	Values      []float64 `json:"values" binding:"required,min=1"`
}

// CollectDataResponse 数据采集响应
type CollectDataResponse struct {
	SampleID    uint                      `json:"sampleId"`
	OocFlag     bool                      `json:"oocFlag"`
	OosFlag     bool                      `json:"oosFlag"`
	Violations  []engine.RuleViolation    `json:"violations,omitempty"`
	Alarms      []uint                    `json:"alarms,omitempty"`
	Message     string                    `json:"message"`
}

// CollectData 采集数据并进行SPC分析
func (s *CollectService) CollectData(req *CollectDataRequest) (*CollectDataResponse, error) {
	chartService := &ChartService{}
	
	// 1. 获取控制图配置
	chart, err := chartService.GetSpcChartByCode(req.ChartCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("控制图不存在: " + req.ChartCode)
		}
		return nil, err
	}

	if chart.Status != 1 {
		return nil, errors.New("控制图未启用")
	}

	// 2. 获取规格
	var spec spc.SpcSpec
	err = global.GVA_DB.Where("id = ?", chart.SpecID).First(&spec).Error
	if err != nil {
		return nil, errors.New("规格配置不存在")
	}

	// 3. 获取当前控制限
	controlLimit, err := chartService.GetCurrentControlLimit(chart.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 4. 计算样本统计量
	mean := engine.CalculateMean(req.Values)
	rangeVal := engine.CalculateRange(req.Values)
	stdVal := engine.CalculateStdDev(req.Values)

	// 5. 检查OOS
	oosFlag := false
	for _, v := range req.Values {
		oos, _ := engine.CheckOOS(v, spec.USL, spec.LSL)
		if oos {
			oosFlag = true
			break
		}
	}

	// 6. 创建样本记录
	sample := &spc.SpcSample{
		ChartID:     chart.ID,
		SampleTime:  &req.SampleTime,
		SubgroupNo:  req.SubgroupNo,
		N:           len(req.Values),
		MeanVal:     &mean,
		RangeVal:    &rangeVal,
		StdVal:      &stdVal,
		OosFlag:     oosFlag,
		OocFlag:     false,
		EquipmentID: req.EquipmentID,
		ChamberID:   req.ChamberID,
		RecipeID:    req.RecipeID,
	}

	// 处理LotID和WaferID
	if req.LotID != nil && *req.LotID != "" {
		var lot spc.SpcLot
		err = global.GVA_DB.Where("lot_id = ?", *req.LotID).First(&lot).Error
		if err == nil {
			sample.LotID = &lot.ID
		}
	}
	
	if req.WaferID != nil && *req.WaferID != "" {
		var wafer spc.SpcWafer
		err = global.GVA_DB.Where("wafer_id = ?", *req.WaferID).First(&wafer).Error
		if err == nil {
			sample.WaferID = &wafer.ID
		}
	}

	err = global.GVA_DB.Create(sample).Error
	if err != nil {
		return nil, errors.New("保存样本失败: " + err.Error())
	}

	// 7. 保存测量值
	for i, v := range req.Values {
		measurement := &spc.SpcMeasurement{
			SampleID: sample.ID,
			SeqNo:    i + 1,
			Value:    &v,
		}
		err = global.GVA_DB.Create(measurement).Error
		if err != nil {
			return nil, errors.New("保存测量值失败")
		}
	}

	// 8. OOC检测
	violations := []engine.RuleViolation{}
	alarmIDs := []uint{}

	if !errors.Is(err, gorm.ErrRecordNotFound) && controlLimit.ID > 0 {
		// 获取最近的样本用于规则检测
		var recentSamples []spc.SpcSample
		err = global.GVA_DB.Where("chart_id = ?", chart.ID).
			Order("sample_time DESC, subgroup_no DESC").
			Limit(30).Find(&recentSamples).Error

		if err == nil && len(recentSamples) > 0 {
			// 反转顺序（从旧到新）
			for i, j := 0, len(recentSamples)-1; i < j; i, j = i+1, j-1 {
				recentSamples[i], recentSamples[j] = recentSamples[j], recentSamples[i]
			}

			// 提取均值序列
			values := make([]float64, len(recentSamples))
			for i, s := range recentSamples {
				if s.MeanVal != nil {
					values[i] = *s.MeanVal
				}
			}

			// 获取启用的规则
			rules, _ := chartService.GetActiveRules(chart.ID)
			enabledRules := []string{}
			for _, r := range rules {
				enabledRules = append(enabledRules, r.RuleCode)
			}

			// 如果没有配置规则，默认使用WE1
			if len(enabledRules) == 0 {
				enabledRules = []string{"WE1"}
			}

			// 检测OOC
			if controlLimit.UCL != nil && controlLimit.LCL != nil && controlLimit.CL != nil {
				violations = engine.CheckOOC(values, *controlLimit.UCL, *controlLimit.CL, *controlLimit.LCL, enabledRules)
			}
		}
	}

	// 9. 创建告警
	if len(violations) > 0 {
		sample.OocFlag = true
		global.GVA_DB.Save(sample)

		for _, v := range violations {
			alarm := &spc.SpcAlarm{
				SampleID:  sample.ID,
				ChartID:   chart.ID,
				AlarmType: "OOC",
				RuleCode:  v.RuleCode,
				Severity:  v.Severity,
				Status:    "OPEN",
				HoldLot:   false,
				Remark:    v.Message,
			}
			err = global.GVA_DB.Create(alarm).Error
			if err == nil {
				alarmIDs = append(alarmIDs, alarm.ID)
			}
		}
	}

	if oosFlag {
		alarm := &spc.SpcAlarm{
			SampleID:  sample.ID,
			ChartID:   chart.ID,
			AlarmType: "OOS",
			RuleCode:  "",
			Severity:  "CRIT",
			Status:    "OPEN",
			HoldLot:   false,
			Remark:    "超出规格限",
		}
		err = global.GVA_DB.Create(alarm).Error
		if err == nil {
			alarmIDs = append(alarmIDs, alarm.ID)
		}
	}

	// 10. 构建响应
	response := &CollectDataResponse{
		SampleID:   sample.ID,
		OocFlag:    sample.OocFlag,
		OosFlag:    sample.OosFlag,
		Violations: violations,
		Alarms:     alarmIDs,
		Message:    "数据采集成功",
	}

	if sample.OocFlag || sample.OosFlag {
		msgs := []string{}
		if sample.OocFlag {
			msgs = append(msgs, "失控")
		}
		if sample.OosFlag {
			msgs = append(msgs, "超规格")
		}
		response.Message = "数据采集成功，检测到异常: " + strings.Join(msgs, "、")
	}

	return response, nil
}
