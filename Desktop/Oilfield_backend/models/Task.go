package models

import (
    "Oilfield_backend/utils"
    "github.com/jinzhu/gorm"
    "time"
)

// 根任务
type RootTask struct {
    Id                int       `gorm:"primary_key;auto_increment"` // 自增主键
    TaskName          string    // 任务名称
    TaskId            string    // 任务id
    CameraName        string    // 摄像头名称
    CameraIp          string    // 摄像头ip地址
    CameraNumber      string    // 摄像头序列号
    StartTime         string    // 开始时间
    EndTime           string    // 结束时间
    AlgorithmReturn   string    // 算法返回类型
    AlgorithmCallback string    // 算法回调地址
    IsTourCheck       int       // 是否为巡检任务(0不是 1是)
    IsStart           int       // 任务是否开始(0未开始 1开始)
    CreateTime        string    // 创建时间
    UpdateTime        string    // 更新时间
    Son               []SonTask // 子任务切片
}

// 子任务
type SonTask struct {
    Id          int    `gorm:"primary_key;auto_increment"` // 自增主键
    TaskName    string // 任务名称
    TaskId      string // 任务id
    AlgorithmId string // 应用的算法id
    ParentId    string // 父任务id
    Area        string // 作用区域
    CatchPhoto  int    // 抓图次数
    CatchTime   int    // 持续时间
    Ptz         string // 云台参数
}
type UpdateTask struct {
    TaskId string                 // 需修改的任务id
    Update map[string]interface{} // 需修改字段
}
type DeleteTask struct {
    TaskId string // 需删除的任务id
}

// PostCreateRootTask 创建并添加根任务
func PostCreateRootTask(check RootTask) error {
    db := utils.DBOpen()
    err := db.Create(&check).Error
    db.Model(&Overview{}).Update("all_monitor", gorm.Expr("all_monitor + ?", 1))
    return err
}

// PostCreateSonTask 创建并添加子任务
func PostCreateSonTask(task SonTask) error {
    db := utils.DBOpen()
    err := db.Create(&task).Error
    return err
}

// GetAllTask 查看所有任务
func GetAllTask() ([]RootTask, error) {
    db := utils.DBOpen()
    var task []RootTask
    err := db.Find(&task).Error
    for i, _ := range task {
        task[i].Son, err = GetSonTask(task[i].TaskId)
    }
    return task, err
}

// GetSomeTask 条件查看任务
func GetSomeTask(TaskType string) ([]RootTask, error) {
    db := utils.DBOpen()
    var task []RootTask
    var err error
    if TaskType == "TourCheck" {
        err = db.Where("is_tour_check = ?", 1).Find(&task).Error
    } else {
        err = db.Where("is_tour_check = ?", 0).Find(&task).Error
    }
    for i, _ := range task {
        task[i].Son, err = GetSonTask(task[i].TaskId)
    }
    return task, err
}

// GetSonTask 根据父id查询任务
func GetSonTask(id string) ([]SonTask, error) {
    db := utils.DBOpen()
    var result []SonTask
    err := db.Where("parent_id = ?", id).Find(&result).Error
    return result, err
}

// PutUpdateRootTask 更新根任务
func PutUpdateRootTask(updateId string, update map[string]interface{}) error {
    db := utils.DBOpen()
    var check RootTask
    updateTime := time.Now().Format("2006-01-02 15:04:05")
    db.Model(&check).Where("task_id = ?", updateId).Update("update_time", updateTime)
    err := db.Model(&check).Where("task_id = ?", updateId).Updates(update).Error
    return err
}

// PutUpdateSonTask 更新子任务
func PutUpdateSonTask(updateId string, update map[string]interface{}) error {
    db := utils.DBOpen()
    var check SonTask
    err := db.Model(&check).Where("task_id = ?", updateId).Updates(update).Error
    return err
}

// DeleteRootTask 删除根任务
func DeleteRootTask(id string) error {
    db := utils.DBOpen()
    var check RootTask
    err := db.Where("task_id = ?", id).Delete(&check).Error
    db.Model(&Overview{}).Update("all_monitor", gorm.Expr("all_monitor - ?", 1))
    return err
}

// DeleteSonTask 根据父id或子id删除子任务
func DeleteSonTask(id string) error {
    db := utils.DBOpen()
    var check SonTask
    err := db.Where("task_id = ?", id).Delete(&check).Error
    err = db.Where("parent_id = ?", id).Delete(&check).Error
    return err
}
