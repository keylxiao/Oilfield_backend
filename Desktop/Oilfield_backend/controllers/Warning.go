package controllers

import (
    "Oilfield_backend/models"
    "Oilfield_backend/utils"
    "github.com/kataras/iris/v12"
    "io"
    "net/http"
    "os"
    "time"
)

// GetAllWarning 查看所有告警
func GetAllWarning(c iris.Context) {
    warn, err := models.GetAllWarning()
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(warn)
    }
}

// GetDownloadFile 告警数据下载
func GetDownloadFile(c iris.Context) {
    name := c.URLParam("name")
    c.StatusCode(http.StatusOK)
    setting, _ := utils.ReadWarning()
    c.SendFile(setting["warningvideopath"].(string)+name+setting["warningvideotype"].(string), name+setting["warningvideotype"].(string))
}

// GetFileStream 告警视频流传输
func GetFileStream(c iris.Context) {
    id := c.URLParam("id")
    c.StatusCode(http.StatusOK)
    setting, _ := utils.ReadWarning()
    http.ServeFile(c.ResponseWriter(), c.Request(), setting["warningvideopath"].(string)+id+setting["warningvideotype"].(string))
}
// GetSettings 获取当前告警配置
func GetSettings(c iris.Context){
    setting ,_ :=utils.ReadWarning()
    c.JSON(setting)
}
// PostCreateWarning 创建告警
func PostCreateWarning(c iris.Context) {
    var warn models.Warning
    c.ReadJSON(&warn)
    warn.WarningTime = time.Now().Unix()
    warn.WarningId = utils.GetUUID()
    err := models.PostCreateWarning(warn)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        models.AddHomeWarning()
        c.StatusCode(http.StatusOK)
        c.JSON(warn.WarningId)
    }
    models.CleanWarning()
}

// PostSomeWarning 条件查看告警
func PostSomeWarning(c iris.Context) {
    var condition map[string]interface{}
    c.ReadJSON(&condition)
    result, err := models.PostSomeWarning(condition)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// PostWarningPhotos 获取告警图片
func PostWarningPhotos(c iris.Context) {
    setting, _ := utils.ReadWarning()
    var receive []models.WarningPhotoReceive
    c.ReadJSON(&receive)
    var send []models.WarningPhotoSend
    send = make([]models.WarningPhotoSend, len(receive))
    for i, k := range receive {
        send[i].WarningId = k.WarningId
        send[i].Photo = utils.PhotoToBase64(setting["warningphotopath"].(string), k.WarningId)
    }
    c.StatusCode(http.StatusOK)
    c.JSON(send)
}

// PostCreateWarningFile 告警数据缓存
func PostCreateWarningFile(c iris.Context) {
    setting, _ := utils.ReadWarning()
    file, _, err := c.FormFile("file")
    if err != nil {
        c.StatusCode(http.StatusBadRequest)
        c.JSON("上传错误")
        return
    }
    var dest *os.File
    if c.FormValue("type") == "photo" {
        dest, err = os.Create(setting["warningphotopath"].(string) + c.FormValue("name"))
    } else {
        dest, err = os.Create(setting["warningvideopath"].(string) + c.FormValue("name"))
    }
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        return
    }
    _, err = io.Copy(dest, file)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("上传成功")
    }
}

// PutUpdateSettings 告警设置更新
func PutUpdateSettings(c iris.Context) {
    var condition map[string]interface{}
    c.ReadJSON(&condition)
    condition["warningstoragetime"] = int(condition["warningstoragetime"].(float64))
    err := utils.WriteWarning(condition)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("ok")
    }
}
