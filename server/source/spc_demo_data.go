package source

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"gorm.io/gorm"
)

type initSpcDemoData struct{}

// auto initialize data
func init() {
	RegisterInit(InitOrderExternal, &initSpcDemoData{})
}

func (i *initSpcDemoData) InitializerName() string {
	return "spc_demo_data"
}

func (i *initSpcDemoData) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (i *initSpcDemoData) TableCreated(ctx context.Context) bool {
	return true
}

func (i *initSpcDemoData) InitializeData(ctx context.Context) (next context.Context, err error) {
	db := global.GVA_DB

	// 检查是否已有数据
	var count int64
	db.Model(&spc.SpcSite{}).Count(&count)
	if count > 0 {
		return ctx, nil
	}

	// 1. 创建厂区
	site := &spc.SpcSite{
		Code:     "FAB1",
		Name:     "Demo Fab 1",
		Timezone: "Asia/Shanghai",
		Status:   1,
		Remark:   "演示厂区",
	}
	if err = db.Create(site).Error; err != nil {
		return ctx, err
	}

	// 2. 创建区域
	areas := []spc.SpcArea{
		{SiteID: site.ID, Code: "PHOTO", Name: "光刻区", Status: 1},
		{SiteID: site.ID, Code: "ETCH", Name: "刻蚀区", Status: 1},
	}
	if err = db.Create(&areas).Error; err != nil {
		return ctx, err
	}

	// 3. 创建设备
	equipments := []spc.SpcEquipment{
		{SiteID: site.ID, AreaID: areas[0].ID, Code: "LITHO-01", Name: "光刻机01", EqpType: "LITHO", Vendor: "ASML", Status: 1},
		{SiteID: site.ID, AreaID: areas[1].ID, Code: "ETCH-01", Name: "刻蚀机01", EqpType: "ETCH", Vendor: "LAM", Status: 1},
		{SiteID: site.ID, AreaID: areas[1].ID, Code: "CD-SEM-01", Name: "CD测量01", EqpType: "METROLOGY", Vendor: "Hitachi", Status: 1},
	}
	if err = db.Create(&equipments).Error; err != nil {
		return ctx, err
	}

	// 4. 创建技术节点和产品
	tech := &spc.SpcTechnology{
		Code:   "28N",
		Name:   "28nm Technology",
		NodeNm: 28.0,
		Status: 1,
	}
	if err = db.Create(tech).Error; err != nil {
		return ctx, err
	}

	product := &spc.SpcProduct{
		TechnologyID: tech.ID,
		Code:         "DEV-28N",
		Name:         "开发产品 28nm",
		Status:       1,
	}
	if err = db.Create(product).Error; err != nil {
		return ctx, err
	}

	// 5. 创建工艺步骤
	steps := []spc.SpcProcessStep{
		{Code: "GATE_LITHO", Name: "栅极光刻", StepType: "LITHO", Status: 1},
		{Code: "GATE_ETCH", Name: "栅极刻蚀", StepType: "ETCH", Status: 1},
	}
	if err = db.Create(&steps).Error; err != nil {
		return ctx, err
	}

	// 6. 创建参数
	gateCD := &spc.SpcParameter{
		Code:          "GATE_CD",
		Name:          "栅极关键尺寸 Gate CD",
		DataType:      "VARIABLE",
		Unit:          "nm",
		DecimalPlaces: 2,
		SampleLevel:   "WAFER",
		Status:        1,
	}
	if err = db.Create(gateCD).Error; err != nil {
		return ctx, err
	}

	oxideThk := &spc.SpcParameter{
		Code:          "OXIDE_THK",
		Name:          "氧化层厚度 Oxide Thickness",
		DataType:      "VARIABLE",
		Unit:          "Å",
		DecimalPlaces: 1,
		SampleLevel:   "WAFER",
		Status:        1,
	}
	if err = db.Create(oxideThk).Error; err != nil {
		return ctx, err
	}

	// 7. 创建规格
	uslCD := 48.0
	targetCD := 45.0
	lslCD := 42.0
	specCD := &spc.SpcSpec{
		ParameterID:   gateCD.ID,
		ProductID:     product.ID,
		ProcessStepID: steps[1].ID,
		Version:       "V1.0",
		USL:           &uslCD,
		Target:        &targetCD,
		LSL:           &lslCD,
		Status:        1,
	}
	if err = db.Create(specCD).Error; err != nil {
		return ctx, err
	}

	uslThk := 1550.0
	targetThk := 1500.0
	lslThk := 1450.0
	specThk := &spc.SpcSpec{
		ParameterID:   oxideThk.ID,
		ProductID:     product.ID,
		ProcessStepID: steps[0].ID,
		Version:       "V1.0",
		USL:           &uslThk,
		Target:        &targetThk,
		LSL:           &lslThk,
		Status:        1,
	}
	if err = db.Create(specThk).Error; err != nil {
		return ctx, err
	}

	// 8. 创建控制图
	chartCD := &spc.SpcChart{
		Code:         "CHART_GATE_CD",
		Name:         "栅极CD控制图",
		ParameterID:  gateCD.ID,
		SpecID:       specCD.ID,
		ChartType:    "XBAR_R",
		SubgroupSize: 5,
		Ruleset:      "WE1,WE2,WE4",
		LimitMethod:  "CALC",
		Status:       1,
	}
	if err = db.Create(chartCD).Error; err != nil {
		return ctx, err
	}

	chartThk := &spc.SpcChart{
		Code:         "CHART_OXIDE_THK",
		Name:         "氧化层厚度控制图",
		ParameterID:  oxideThk.ID,
		SpecID:       specThk.ID,
		ChartType:    "I_MR",
		SubgroupSize: 1,
		Ruleset:      "WE1,WE4,NELSON5",
		LimitMethod:  "CALC",
		Status:       1,
	}
	if err = db.Create(chartThk).Error; err != nil {
		return ctx, err
	}

	// 9. 创建控制限
	uclCD := 47.0
	clCD := 45.0
	lclCD := 43.0
	uclR := 6.0
	clR := 3.0
	limitCD := &spc.SpcControlLimit{
		ChartID: chartCD.ID,
		UCL:     &uclCD,
		CL:      &clCD,
		LCL:     &lclCD,
		UCLS:    &uclR,
		CLS:     &clR,
		CalcN:   30,
		Source:  "CALC",
	}
	if err = db.Create(limitCD).Error; err != nil {
		return ctx, err
	}

	uclThk := 1520.0
	clThk := 1500.0
	lclThk := 1480.0
	limitThk := &spc.SpcControlLimit{
		ChartID: chartThk.ID,
		UCL:     &uclThk,
		CL:      &clThk,
		LCL:     &lclThk,
		CalcN:   30,
		Source:  "CALC",
	}
	if err = db.Create(limitThk).Error; err != nil {
		return ctx, err
	}

	// 10. 创建规则
	rules := []spc.SpcRule{
		{ChartID: chartCD.ID, RuleCode: "WE1", Enabled: true, N: 1, K: 3.0, Remark: "点超出控制限"},
		{ChartID: chartCD.ID, RuleCode: "WE2", Enabled: true, N: 3, K: 2.0, Remark: "3点中2点超出2σ"},
		{ChartID: chartCD.ID, RuleCode: "WE4", Enabled: true, N: 8, K: 0.0, Remark: "连续8点同侧"},
	}
	if err = db.Create(&rules).Error; err != nil {
		return ctx, err
	}

	// 11. 创建演示批次和晶圆
	lot := &spc.SpcLot{
		SiteID:    site.ID,
		ProductID: product.ID,
		LotID:     "LOT001",
		LotType:   "PROD",
		Qty:       25,
		Status:    1,
	}
	if err = db.Create(lot).Error; err != nil {
		return ctx, err
	}

	// 12. 创建演示样本数据（含正常和异常数据）
	baseTime := time.Now().Add(-48 * time.Hour)
	samples := make([]*spc.SpcSample, 0, 50)

	// 正常数据 (前30个点)
	for i := 0; i < 30; i++ {
		mean := 45.0 + float64(i%5-2)*0.3
		rangeVal := 2.5 + float64(i%3)*0.5
		stdVal := 1.0 + float64(i%3)*0.2
		sampleTime := baseTime.Add(time.Duration(i) * time.Hour)

		sample := &spc.SpcSample{
			ChartID:     chartCD.ID,
			LotID:       &lot.ID,
			EquipmentID: &equipments[1].ID,
			SampleTime:  &sampleTime,
			SubgroupNo:  i + 1,
			N:           5,
			MeanVal:     &mean,
			RangeVal:    &rangeVal,
			StdVal:      &stdVal,
			OocFlag:     false,
			OosFlag:     false,
		}
		samples = append(samples, sample)
	}

	// 异常数据 - OOC (点31-33: 连续3点偏高)
	for i := 30; i < 33; i++ {
		mean := 46.5 + float64(i-30)*0.2
		rangeVal := 3.0
		stdVal := 1.2
		sampleTime := baseTime.Add(time.Duration(i) * time.Hour)

		sample := &spc.SpcSample{
			ChartID:     chartCD.ID,
			LotID:       &lot.ID,
			EquipmentID: &equipments[1].ID,
			SampleTime:  &sampleTime,
			SubgroupNo:  i + 1,
			N:           5,
			MeanVal:     &mean,
			RangeVal:    &rangeVal,
			StdVal:      &stdVal,
			OocFlag:     false,
			OosFlag:     false,
		}
		samples = append(samples, sample)
	}

	// 异常数据 - OOS (点34: 超规格)
	meanOOS := 48.5
	rangeVal := 2.8
	stdVal := 1.1
	sampleTimeOOS := baseTime.Add(34 * time.Hour)
	sampleOOS := &spc.SpcSample{
		ChartID:     chartCD.ID,
		LotID:       &lot.ID,
		EquipmentID: &equipments[1].ID,
		SampleTime:  &sampleTimeOOS,
		SubgroupNo:  34,
		N:           5,
		MeanVal:     &meanOOS,
		RangeVal:    &rangeVal,
		StdVal:      &stdVal,
		OocFlag:     false,
		OosFlag:     true,
	}
	samples = append(samples, sampleOOS)

	// 恢复正常 (点35-50)
	for i := 35; i < 50; i++ {
		mean := 45.0 + float64(i%4-2)*0.25
		rangeVal := 2.7 + float64(i%3)*0.4
		stdVal := 1.05 + float64(i%3)*0.15
		sampleTime := baseTime.Add(time.Duration(i) * time.Hour)

		sample := &spc.SpcSample{
			ChartID:     chartCD.ID,
			LotID:       &lot.ID,
			EquipmentID: &equipments[1].ID,
			SampleTime:  &sampleTime,
			SubgroupNo:  i + 1,
			N:           5,
			MeanVal:     &mean,
			RangeVal:    &rangeVal,
			StdVal:      &stdVal,
			OocFlag:     false,
			OosFlag:     false,
		}
		samples = append(samples, sample)
	}

	if err = db.Create(&samples).Error; err != nil {
		return ctx, err
	}

	// 13. 为OOS样本创建告警
	alarmOOS := &spc.SpcAlarm{
		SampleID:  sampleOOS.ID,
		ChartID:   chartCD.ID,
		AlarmType: "OOS",
		Severity:  "CRIT",
		Status:    "OPEN",
		HoldLot:   false,
		Remark:    "超出规格上限",
	}
	if err = db.Create(alarmOOS).Error; err != nil {
		return ctx, err
	}

	// 14. 创建OCAP模板
	ocap := &spc.SpcOcap{
		ChartID:     chartCD.ID,
		Name:        "GATE_CD异常处理",
		TriggerType: "BOTH",
		StepsJson:   `[{"step":1,"action":"检查设备状态"},{"step":2,"action":"检查工艺参数"},{"step":3,"action":"重新校准设备"},{"step":4,"action":"通知工程师"}]`,
		Status:      1,
	}
	if err = db.Create(ocap).Error; err != nil {
		return ctx, err
	}

	global.GVA_LOG.Info("SPC演示数据初始化完成")
	return ctx, nil
}

func (i *initSpcDemoData) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var count int64
	db.Model(&spc.SpcSite{}).Where("code = ?", "FAB1").Count(&count)
	return count > 0
}
