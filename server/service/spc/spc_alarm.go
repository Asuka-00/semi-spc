package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
)

type AlarmService struct{}

// GetSpcAlarmList 分页获取告警列表
func (s *AlarmService) GetSpcAlarmList(info request.PageInfo, status string, alarmType string) (list []spc.SpcAlarm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&spc.SpcAlarm{}).Preload("Sample").Preload("Chart")

	if status != "" {
		db = db.Where("status = ?", status)
	}
	if alarmType != "" {
		db = db.Where("alarm_type = ?", alarmType)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&list).Error
	return
}

// AcknowledgeAlarm 确认告警
func (s *AlarmService) AcknowledgeAlarm(id uint, remark string) error {
	return global.GVA_DB.Model(&spc.SpcAlarm{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": "ACK",
			"remark": remark,
		}).Error
}

// CloseAlarm 关闭告警
func (s *AlarmService) CloseAlarm(id uint, remark string) error {
	return global.GVA_DB.Model(&spc.SpcAlarm{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": "CLOSED",
			"remark": remark,
		}).Error
}

// GetOpenAlarmCount 获取未关闭告警数量
func (s *AlarmService) GetOpenAlarmCount() (count int64, err error) {
	err = global.GVA_DB.Model(&spc.SpcAlarm{}).
		Where("status IN ?", []string{"OPEN", "ACK"}).Count(&count).Error
	return
}

// GetAlarmStatistics 获取告警统计
func (s *AlarmService) GetAlarmStatistics(days int) (stats map[string]interface{}, err error) {
	stats = make(map[string]interface{})

	// 按类型统计
	var typeStats []struct {
		AlarmType string `json:"alarmType"`
		Count     int64  `json:"count"`
	}
	err = global.GVA_DB.Model(&spc.SpcAlarm{}).
		Select("alarm_type, COUNT(*) as count").
		Where("created_at >= DATE_SUB(NOW(), INTERVAL ? DAY)", days).
		Group("alarm_type").Find(&typeStats).Error
	if err != nil {
		return nil, err
	}
	stats["byType"] = typeStats

	// 按状态统计
	var statusStats []struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}
	err = global.GVA_DB.Model(&spc.SpcAlarm{}).
		Select("status, COUNT(*) as count").
		Where("created_at >= DATE_SUB(NOW(), INTERVAL ? DAY)", days).
		Group("status").Find(&statusStats).Error
	if err != nil {
		return nil, err
	}
	stats["byStatus"] = statusStats

	// 按严重度统计
	var severityStats []struct {
		Severity string `json:"severity"`
		Count    int64  `json:"count"`
	}
	err = global.GVA_DB.Model(&spc.SpcAlarm{}).
		Select("severity, COUNT(*) as count").
		Where("created_at >= DATE_SUB(NOW(), INTERVAL ? DAY)", days).
		Group("severity").Find(&severityStats).Error
	if err != nil {
		return nil, err
	}
	stats["bySeverity"] = severityStats

	return stats, nil
}
