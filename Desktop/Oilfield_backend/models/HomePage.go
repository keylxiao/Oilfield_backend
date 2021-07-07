package models

import "Oilfield_backend/utils"

// 项目总览
type Overview struct {
    AllProject   int // 项目总数
    AllMonitor   int // 监控总数
    AllAlgorithm int // 算法总数
    AllWarning   int // 告警总数
}

// GetOverview 首页总览信息
func GetOverview() (Overview, error) {
    db := utils.DBOpen()
    var result Overview
    err := db.Find(&result).Error
    return result, err
}
