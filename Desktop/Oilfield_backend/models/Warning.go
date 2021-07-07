package models

import (
    "Oilfield_backend/utils"
    "github.com/jinzhu/gorm"
    "os"
    "time"
)

type Warning struct {
    Id            int    `gorm:"primary_key;auto_increment"` // 自增主键
    WarningId     string // 告警id
    CameraNumber  string // 摄像头序列号
    WarningReason string // 告警原因
    WarningTime   int64  // 告警时间戳
    WarningDegree int    // 告警等级
}
type WarningPhotoReceive struct {
    WarningId string // 告警id
}
type WarningPhotoSend struct {
    WarningId string // 告警id
    Photo     string // 告警图片base64
}

//type WarningSetting struct {
//    WarningPhotoPath   string // 告警图片储存地址
//    WarningVideoPath   string // 告警视频储存地址
//    WarningVideoType   string // 告警视频格式
//    WarningPhotoType   string // 告警图片格式
//    WarningStorageTime string // 文件储存时间
//}

// GetAllWarning 查看所有告警
func GetAllWarning() ([]Warning, error) {
    db := utils.DBOpen()
    var warn []Warning
    err := db.Find(&warn).Error
    return warn, err
}

// PostCreateWarning // 添加告警
func PostCreateWarning(warn Warning) error {
    db := utils.DBOpen()
    err := db.Create(&warn).Error
    return err
}

// PostSomeWarning 条件查看告警
func PostSomeWarning(condition map[string]interface{}) ([]Warning, error) {
    db := utils.DBOpen()
    var result []Warning
    err := db.Where(condition).Find(&result).Error
    return result, err
}

// AddHomeWarning 添加预览告警
func AddHomeWarning() {
    db := utils.DBOpen()
    db.Model(&Overview{}).Update("all_monitor", gorm.Expr("all_monitor + ?", 1))
}

// CleanWarning 清理告警
func CleanWarning() {
    db := utils.DBOpen()
    setting, _ := utils.ReadWarning()
    var warn []Warning
    var warnId []Warning
    overTime := time.Now().Unix() - setting["warningstoragetime"].(int64)*86400
    // 收集清理的告警id
    db.Where("warning_time < ?", overTime).Find(&warnId)
    // 清理告警日志
    db.Where("warning_time < ?", overTime).Delete(&warn)
    // 清理告警文件
    for _, k := range warnId {
        photo := setting["warningphotopath"].(string) + k.WarningId + setting["warningphototype"].(string)
        video := setting["warningvideopath"].(string) + k.WarningId + setting["warningvideotype"].(string)
        os.Remove(photo)
        os.Remove(video)
    }
}
