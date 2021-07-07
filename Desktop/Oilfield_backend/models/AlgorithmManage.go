package models

import (
    "Oilfield_backend/utils"
    "github.com/jinzhu/gorm"
)

// 算法组件
type Algorithm struct {
    Id          int    `gorm:"primary_key;auto_increment"` // 自增主键
    AlgorithmId string // 项目id
    EnglishName string // 算法英文名
    Version     string // 版本号
    Type        string // 类型
    Writer      string // 作者
    Description string // 算法描述
    Address     string // 算法路径
    CreateTime  string // 创建时间
    UpdateTime  string // 修改时间
}
type UpdateAlgorithm struct {
    AlgorithmId string                 // 需修改的巡检id
    Update      map[string]interface{} // 需修改字段
}
type DeleteSomeAlgorithm struct {
    AlgorithmId string // 需删除的巡检id
}

// 创建一条算法
func PostCreateAlgorithm(alg Algorithm) error {
    db := utils.DBOpen()
    err := db.Create(&alg).Error
    db.Model(&Overview{}).Update("all_algorithm", gorm.Expr("all_algorithm + ?", 1))
    return err
}

// 查找所有算法
func GetAllAlgorithm() ([]Algorithm, error) {
    db := utils.DBOpen()
    var alg []Algorithm
    err := db.Find(&alg).Error
    return alg, err
}

// 更新算法信息
func PutUpdateAlgorithm(updateId string, update map[string]interface{}) error {
    db := utils.DBOpen()
    var alg Algorithm
    err := db.Model(&alg).Where("algorithm_id = ?", updateId).Updates(update).Error
    return err
}

// 删除算法
func DeleteAlgorithm(id string) error {
    db := utils.DBOpen()
    var check Algorithm
    err := db.Where("algorithm_id = ?", id).Delete(&check).Error
    db.Model(&Overview{}).Update("all_algorithm", gorm.Expr("all_algorithm - ?", 1))
    return err
}
