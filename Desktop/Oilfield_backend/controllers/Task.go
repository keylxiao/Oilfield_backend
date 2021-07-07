package controllers

import (
    "Oilfield_backend/Timer"
    "Oilfield_backend/algorithm"
    "Oilfield_backend/config"
    "Oilfield_backend/models"
    "Oilfield_backend/utils"
    "fmt"
    "github.com/kataras/iris/v12"
    "net/http"
    "time"
)

var channel = make(chan string, 1)

// PostCreateRootTask 创建并添加根任务
func PostCreateRootTask(c iris.Context) {
    var tour models.RootTask
    c.ReadJSON(&tour)
    tour.TaskId = utils.GetUUID()
    tour.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    tour.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
    err := models.PostCreateRootTask(tour)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(tour.TaskId)
    }
}

// PostCreateSonTask 创建并添加子任务
func PostCreateSonTask(c iris.Context) {
    var task []models.SonTask
    c.ReadJSON(&task)
    for _, k := range task {
        // 先加任务, 再加记录
        k.TaskId = utils.GetUUID()
        state, err := algorithm.AddTask(k)
        if err != nil || state == "false" {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("添加任务失败")
            return
        }
        err = models.PostCreateSonTask(k)
        if err != nil {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("添加任务记录失败")
            return
        }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("ok")
}

// GetAllTask 查看所有任务
func GetAllTask(c iris.Context) {
    result, err := models.GetAllTask()
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// GetSomeTask 条件查看任务
func GetSomeTask(c iris.Context) {
    taskType := c.URLParam("type")
    result, err := models.GetSomeTask(taskType)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// PutUpdateRootTask 更新根任务
func PutUpdateRootTask(c iris.Context) {
    var update []models.UpdateTask
    c.ReadJSON(&update)
    for _, result := range update {
        // 先更新任务, 再更新记录
        state, err := algorithm.CheckRootUpdate(result.TaskId, result.Update)
        if err != nil || state == "false" {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("更新任务失败")
            return
        }
        err = models.PutUpdateRootTask(result.TaskId, result.Update)
        if err != nil {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("更新任务记录失败")
            return
        }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("更新成功")
}

// PutUpdateSonTask 更新子任务
func PutUpdateSonTask(c iris.Context) {
    var update []models.UpdateTask
    c.ReadJSON(&update)
    for _, result := range update {
        // 先更新任务, 再更新记录
        state, err := algorithm.UpdateTask(result.TaskId, result.Update)
        if err != nil || state == "false" {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("更新任务失败")
            return
        }
        err = models.PutUpdateSonTask(result.TaskId, result.Update)
        if err != nil {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("更新任务记录失败")
            return
        }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("更新成功")
}

// DeleteRootTask 删除根任务
func DeleteRootTask(c iris.Context) {
    var result []models.DeleteTask
    c.ReadJSON(&result)
    for _, id := range result {
        // 先删任务, 再删记录
        state, err := algorithm.DeleteTask(id.TaskId)
        if err != nil || state == "false" {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("删除任务失败")
            return
        }
        err = models.DeleteRootTask(id.TaskId)
        err = models.DeleteSonTask(id.TaskId)
        if err != nil {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("删除任务记录失败")
            return
        }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("删除成功")
}

// DeleteSonTask 删除子任务
func DeleteSonTask(c iris.Context) {
    var result []models.DeleteTask
    c.ReadJSON(&result)
    for _, id := range result {
        // 先删任务, 再删记录
        state, err := algorithm.DeleteTask(id.TaskId)
        if err != nil || state == "false" {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("删除任务失败")
            return
        }
        err = models.DeleteSonTask(id.TaskId)
        if err != nil {
            c.StatusCode(http.StatusInternalServerError)
            c.JSON("删除任务记录失败")
            return
        }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("删除成功")
}

// GetRtspToRtmp RTSP转流
func GetRtspToRtmp(c iris.Context) {
    rtsp := c.URLParam("rtsp")
    ip := c.URLParam("ip")
    rtmp := fmt.Sprintf(config.RtmpPath + ip)
    red := utils.RedisOpen()
    end := time.NewTimer(3 * time.Second)

    ch1 := make(chan int)
    count, _ := red.Get(ip).Int()
    // 如果存在这个协程, 直接返回它的rtmp地址
    if count != 0 {
        red.Set(ip, count+1, 0)
        c.StatusCode(http.StatusOK)
        c.JSON(rtmp)
        return
    } else {
        go Timer.RtspToRtmp(ip, rtsp, rtmp, ch1, channel)
    }

    // 没有这个协程时, 尝试转流
    select {
    // 转流成功
    case <-end.C:
        a, _ := red.Get(ip).Int()
        red.Set(ip, a+1, 0)
        c.StatusCode(http.StatusOK)
        c.JSON(rtmp)
        return
    // 转流失败
    case <-ch1:
        c.StatusCode(http.StatusOK)
        c.JSON("转流失败")
        return
    }
}

// GetStopRtmp 某用户停止一个rtmp流
func GetStopRtmp(c iris.Context){
    rtmp := c.URLParam("rtmp")
    channel <- rtmp
    c.StatusCode(http.StatusOK)
}