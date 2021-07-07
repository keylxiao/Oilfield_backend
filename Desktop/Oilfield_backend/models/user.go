package models

import (
    "Oilfield_backend/utils"
    "time"
)

type User struct {
    Id         int    `gorm:"primary_key;auto_increment"` // 自增主键
    Account    string // 登录账号
    Password   string // 密码(sha1加密)
    UserName   string // 用户名
    TelePhone  string // 电话号码
    Status     int    // 用户权限(0系统管理员 1项目管理员 2普通人员)
    CreateTime string // 创建时间
    UpdateTime string // 修改时间
    IsDelete   int    // 逻辑删除(0未删 1删除)
}
type UserInfo struct {
    Account string                 // 需修改的用户账号
    Update  map[string]interface{} // 需修改字段
}
type UserState struct {
    Account string // 账号
    Status  int    // 预期权限
}

// PostUserRegister 用户注册
func PostUserRegister(person User) error {
    db := utils.DBOpen()
    err := db.Create(&person).Error
    return err
}

// CheckOnlyOne 检查用户名或登录账号是否已被注册
func CheckOnlyOne(account, name string) string {
    db := utils.DBOpen()
    var user User
    db.Where("user_name = ?", name).Where("is_delete = ?", 0).Find(&user)
    if user.UserName != "" {
        return "用户名已被注册"
    }
    db.Where("account = ?", account).Where("is_delete = ?", 0).Find(&user)
    if user.UserName != "" {
        return "账号已被注册"
    }
    return ""
}

// PostUserLogin 用户登录
func PostUserLogin(account, password string) (bool, error) {
    db := utils.DBOpen()
    var user User
    err := db.Where("account = ?", account).Find(&user).Error
    if err != nil {
        return false, err
    }
    if user.IsDelete == 1 {
        return false, err
    }
    if user.Password == password {
        return true, err
    } else {
        return false, err
    }
}

// GetUserInfo 查看用户信息
func GetUserInfo(name string) (User, error) {
    db := utils.DBOpen()
    var user User
    err := db.Where("user_name = ?", name).Find(&user).Error
    return user, err
}

// GetAllUserInfo 查看所有用户信息
func GetAllUserInfo() ([]User, error) {
    db := utils.DBOpen()
    var user []User
    err := db.Find(&user).Error
    return user, err
}

// PutUpdateUserStatus 修改用户权限
func PutUpdateUserStatus(name string, status int) error {
    db := utils.DBOpen()
    var user User
    updateTime := time.Now().Format("2006-01-02 15:04:05")
    db.Model(&user).Where("account = ?", name).Update("update_time", updateTime)
    err := db.Model(&user).Where("account = ?", name).Update("status", status).Error
    return err
}

// DeleteUser 删除用户
func DeleteUser(name string) error {
    db := utils.DBOpen()
    var user User
    updateTime := time.Now().Format("2006-01-02 15:04:05")
    db.Model(&user).Where("user_name = ?", name).Update("update_time", updateTime)
    err := db.Model(&user).Where("user_name = ?", name).Update("is_delete", 1).Error
    return err
}

// PutUpdateUserInfo 用户信息修改
func PutUpdateUserInfo(account string, info map[string]interface{}) error {
    db := utils.DBOpen()
    var user User
    updateTime := time.Now().Format("2006-01-02 15:04:05")
    db.Model(&user).Where("account = ?", account).Update("update_time", updateTime)
    err := db.Model(&user).Where("account = ?", account).Updates(info).Error
    return err
}

// AddOnline 添加上线缓存
func AddOnline(account, ip string) error {
    db := utils.RedisOpen()
    err := db.Set(account, ip, 0).Err()
    return err
}

// CheckOnline 轮询检查
func CheckOnline(account, ip string) (bool, error) {
    db := utils.RedisOpen()
    lastIp, err := db.Get(account).Result()
    if lastIp == ip {
        return true, err
    } else {
        return false, err
    }
}

// AddOffline 删除上线缓存
func AddOffline(account string) error {
    db := utils.RedisOpen()
    return db.Set(account, "", 0).Err()
}
