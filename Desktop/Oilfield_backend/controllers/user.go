package controllers

import (
    "Oilfield_backend/models"
    "Oilfield_backend/utils"
    "github.com/kataras/iris/v12"
    "net/http"
    "time"
)

// GetUserOnlineCheck 用户登录轮询
func GetUserOnlineCheck(c iris.Context) {
    account := c.URLParam("account")
    check, err := models.CheckOnline(account, utils.GetIP(c.Request()))
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        return
    }
    // ip匹配
    if check {
        c.StatusCode(http.StatusOK)
        c.JSON("true")
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("false")
    }
}

// GetUserOffline 用户下线, 清空缓存
func GetUserOffline(c iris.Context) {
    account := c.URLParam("account")
    err := models.AddOffline(account)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("ok")
    }
}

// PostUserRegister 用户注册
func PostUserRegister(c iris.Context) {
    var user models.User
    c.ReadJSON(&user)
    // 查重
    only := models.CheckOnlyOne(user.Account, user.UserName)
    if only != "" {
        c.StatusCode(http.StatusOK)
        c.JSON(only)
        return
    }
    user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    user.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
    user.IsDelete = 0
    err := models.PostUserRegister(user)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("ok")
    }
}

// PostUserLogin 用户登录
func PostUserLogin(c iris.Context) {
    var user models.User
    c.ReadJSON(&user)
    // 检查用户是否存在
    only := models.CheckOnlyOne(user.Account, "")
    if only != "账号已被注册" {
        c.StatusCode(http.StatusOK)
        c.JSON("not found")
        return
    }
    isReal, err := models.PostUserLogin(user.Account, user.Password)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        return
    }
    if isReal {
        models.AddOnline(user.Account, utils.GetIP(c.Request()))
        c.StatusCode(http.StatusOK)
        c.JSON("true")
        // 写入日志
        utils.WriteLoginLog(user.Account)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("false")
    }
}

// GetUserLoginLog 查看登录情况
func GetUserLoginLog(c iris.Context) {
    log, err := utils.ReadLoginLog()
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(log)
    }
}

// GetUserInfo 查看用户信息
func GetUserInfo(c iris.Context) {
    name := c.URLParam("name")
    result, err := models.GetUserInfo(name)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }

}

// GetAllUserInfo 查看所有用户信息
func GetAllUserInfo(c iris.Context) {
    result, err := models.GetAllUserInfo()
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// PutUpdateUserStatus 修改用户权限
func PutUpdateUserStatus(c iris.Context) {
    var target models.UserState
    c.ReadJSON(&target)
    err := models.PutUpdateUserStatus(target.Account, target.Status)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("修改成功")
    }
}

// PutUpdateUserInfo 用户信息修改
func PutUpdateUserInfo(c iris.Context) {
    var update models.UserInfo
    c.ReadJSON(&update)
    err := models.PutUpdateUserInfo(update.Account, update.Update)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("修改成功")
    }
}

// DeleteUser 删除用户
func DeleteUser(c iris.Context) {
    name := c.URLParam("name")
    err := models.DeleteUser(name)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("删除成功")
    }
}
