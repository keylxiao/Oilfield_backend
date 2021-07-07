package models

import (
    "Oilfield_backend/utils"
)

type Ipc struct {
    Id          int    `gorm:"primary_key;auto_increment"` // 自增主键
    Number      string // 设备序列号
    Name        string // 设备名称
    Account     string // 设备账号
    Password    string // 设备密码
    Ip          string // 设备ip
    Nvr         string // 关联的NVR序列号
    Factory     string // 厂商
    Type        string // 设备型号
    Area        string // 安装区域
    Place       string // 安装位置
    IsOnline    int    // 是否在线(0在线 1不在线)
    ISTourCheck int    //是否能进行巡检(0不能 1能)
    IsStart     int    // 是否激活(0激活 1未激活)
}

type Nvr struct {
    Id       int    `gorm:"primary_key;auto_increment"` // 自增主键
    Number   string // 设备序列号
    Name     string // 设备名称
    Account  string // 设备账号
    Password string // 设备密码
    Ip       string // 设备ip
    Factory  string // 厂商
    Type     string // 设备型号
    IsOnline int    // 是否在线(0在线 1不在线)
}
type UpdateEquip struct {
    Type  string // 修改设备类型
    Equip        // 具体修改
}
type Equip struct {
    EquipType string                 // 查看类型/设备序列号
    Condition map[string]interface{} // 筛选条件/修改条件
}
type DeleteSomeEquip struct {
    EquipType string // 删除设备类型
    Number    string // 删除设备序列号
}

// GetAllEquip 获取所有设备信息
func GetAllEquip(equip string) (interface{}, error) {
    db := utils.DBOpen()
    var ipc []Ipc
    var nvr []Nvr
    if equip == "ipc" {
        err := db.Find(&ipc).Error
        return ipc, err
    } else {
        err := db.Find(&nvr).Error
        return nvr, err
    }
}

// PostCreateEquip 添加设备
func PostCreateEquip(equip string, value interface{}) error {
    db := utils.DBOpen()
    if equip == "ipc" {
        result := value.(Ipc)
        err := db.Create(&result).Error
        return err
    } else {
        result := value.(Nvr)
        err := db.Create(&result).Error
        return err
    }
}

// PostSomeEquip 条件查看设备
func PostSomeEquip(equip string, condition map[string]interface{}) (interface{}, error) {
    db := utils.DBOpen()
    var ipc []Ipc
    var nvr []Nvr
    if equip == "ipc" {
        err := db.Where(condition).Find(&ipc).Error
        return ipc, err
    } else {
        err := db.Where(condition).Find(&nvr).Error
        return nvr, err
    }
}

// PutUpdateEquip 修改设备属性
func PutUpdateEquip(equip, number string, condition map[string]interface{}) error {
    db := utils.DBOpen()
    var ipc Ipc
    var nvr Nvr
    if equip == "ipc" {
        err := db.Model(&ipc).Where("number = ?", number).Update(condition).Error
        return err
    } else {
        err := db.Model(&nvr).Where("number = ?", number).Update(condition).Error
        return err
    }
}

// DeleteEquip 删除设备
func DeleteEquip(equip string, number string) error {
    db := utils.DBOpen()
    var ipc Ipc
    var nvr Nvr
    if equip == "ipc" {
        err := db.Where("number = ?", number).Update(&ipc).Error
        return err
    } else {
        err := db.Where("number = ?", number).Delete(&nvr).Error
        return err
    }
}

// UpdateOnline 根据ip修改设备在线状态
func UpdateOnline(equip, ip string,condition int) {
    db := utils.DBOpen()
    tx := db.Begin()
    var ipc Ipc
    var nvr Nvr
    var err error
    if equip == "ipc" {
        err = tx.Model(&ipc).Where("ip = ?", ip).Update("is_online",condition).Error
    } else {
        err = tx.Model(&nvr).Where("ip = ?", ip).Update("is_online",condition).Error
    }
    if err != nil{
        tx.Rollback()
    }
    tx.Commit()
}
