package spc

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type ControlLimitService struct{}

func (s *ControlLimitService) CreateSpcControlLimit(limit *spc.SpcControlLimit) error {
	// 验证控制限顺序: UCL > CL > LCL
	if limit.UCL != nil && limit.CL != nil && *limit.UCL <= *limit.CL {
		return errors.New("上控制限必须大于中心线")
	}
	if limit.CL != nil && limit.LCL != nil && *limit.CL <= *limit.LCL {
		return errors.New("中心线必须大于下控制限")
	}
	if limit.UCL != nil && limit.LCL != nil && *limit.UCL <= *limit.LCL {
		return errors.New("上控制限必须大于下控制限")
	}
	
	// S图的控制限也需要验证
	if limit.UCLS != nil && limit.CLS != nil && *limit.UCLS <= *limit.CLS {
		return errors.New("S图上控制限必须大于中心线")
	}
	if limit.CLS != nil && limit.LCLS != nil && *limit.CLS <= *limit.LCLS {
		return errors.New("S图中心线必须大于下控制限")
	}
	
	return global.GVA_DB.Create(limit).Error
}

func (s *ControlLimitService) DeleteSpcControlLimit(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcControlLimit{}, id).Error
}

func (s *ControlLimitService) UpdateSpcControlLimit(limit *spc.SpcControlLimit) error {
	return global.GVA_DB.Save(limit).Error
}

func (s *ControlLimitService) GetSpcControlLimit(id uint) (limit spc.SpcControlLimit, err error) {
	err = global.GVA_DB.Preload("Chart").Where("id = ?", id).First(&limit).Error
	return
}

func (s *ControlLimitService) GetSpcControlLimitList(info request.PageInfo, chartID uint) (list []spc.SpcControlLimit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcControlLimit{}).Preload("Chart")
	
	if chartID > 0 {
		db = db.Where("chart_id = ?", chartID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	return
}

type RuleService struct{}

func (s *RuleService) CreateSpcRule(rule *spc.SpcRule) error {
	return global.GVA_DB.Create(rule).Error
}

func (s *RuleService) DeleteSpcRule(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcRule{}, id).Error
}

func (s *RuleService) UpdateSpcRule(rule *spc.SpcRule) error {
	return global.GVA_DB.Save(rule).Error
}

func (s *RuleService) GetSpcRule(id uint) (rule spc.SpcRule, err error) {
	err = global.GVA_DB.Preload("Chart").Where("id = ?", id).First(&rule).Error
	return
}

func (s *RuleService) GetSpcRuleList(info request.PageInfo, chartID uint) (list []spc.SpcRule, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcRule{}).Preload("Chart")
	
	if chartID > 0 {
		db = db.Where("chart_id = ?", chartID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type SampleService struct{}

func (s *SampleService) GetSpcSample(id uint) (sample spc.SpcSample, err error) {
	err = global.GVA_DB.Preload("Chart").Preload("Lot").Preload("Wafer").
		Preload("Equipment").Preload("Chamber").Preload("Recipe").
		Where("id = ?", id).First(&sample).Error
	return
}

func (s *SampleService) GetSpcSampleList(info request.PageInfo, chartID uint) (list []spc.SpcSample, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcSample{}).
		Preload("Chart").Preload("Lot").Preload("Equipment")
	
	if chartID > 0 {
		db = db.Where("chart_id = ?", chartID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Order("sample_time DESC").Limit(limit).Offset(offset).Find(&list).Error
	return
}

type MeasurementService struct{}

func (s *MeasurementService) GetSpcMeasurementList(sampleID uint) (list []spc.SpcMeasurement, err error) {
	err = global.GVA_DB.Where("sample_id = ?", sampleID).Order("seq_no ASC").Find(&list).Error
	return
}

type OcapService struct{}

func (s *OcapService) CreateSpcOcap(ocap *spc.SpcOcap) error {
	return global.GVA_DB.Create(ocap).Error
}

func (s *OcapService) DeleteSpcOcap(id uint) error {
	return global.GVA_DB.Delete(&spc.SpcOcap{}, id).Error
}

func (s *OcapService) UpdateSpcOcap(ocap *spc.SpcOcap) error {
	return global.GVA_DB.Save(ocap).Error
}

func (s *OcapService) GetSpcOcap(id uint) (ocap spc.SpcOcap, err error) {
	err = global.GVA_DB.Preload("Chart").Where("id = ?", id).First(&ocap).Error
	return
}

func (s *OcapService) GetSpcOcapList(info request.PageInfo, chartID uint) (list []spc.SpcOcap, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcOcap{}).Preload("Chart")
	
	if chartID > 0 {
		db = db.Where("chart_id = ?", chartID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return
}

type OcapExecutionService struct{}

func (s *OcapExecutionService) CreateSpcOcapExecution(exec *spc.SpcOcapExecution) error {
	return global.GVA_DB.Create(exec).Error
}

func (s *OcapExecutionService) UpdateSpcOcapExecution(exec *spc.SpcOcapExecution) error {
	return global.GVA_DB.Save(exec).Error
}

func (s *OcapExecutionService) GetSpcOcapExecution(id uint) (exec spc.SpcOcapExecution, err error) {
	err = global.GVA_DB.Preload("Alarm").Preload("Ocap").Where("id = ?", id).First(&exec).Error
	return
}

func (s *OcapExecutionService) GetSpcOcapExecutionList(info request.PageInfo, alarmID uint) (list []spc.SpcOcapExecution, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcOcapExecution{}).Preload("Alarm").Preload("Ocap")
	
	if alarmID > 0 {
		db = db.Where("alarm_id = ?", alarmID)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	return
}
