package algorithm

import (
    "Oilfield_backend/config"
    "Oilfield_backend/models"
    "Oilfield_backend/utils"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strings"
)

type AlgTask struct {
    TaskId            string // 任务id
    StartTime         string // 开始时间
    EndTime           string // 结束时间
    AlgorithmId       string // 算法id
    AlgorithmReturn   string // 算法返回类型
    AlgorithmCallback string // 算法回调地址
    Ptz               string // 云台参数
    Area              string // 作用区域
    CatchPhoto        int    // 抓图次数
    CatchTime         int    // 持续时间
    CameraIp          string // 摄像头ip
}
type Update struct {
    TaskId            string // 任务id
    StartTime         string // 开始时间
    EndTime           string // 结束时间
    AlgorithmId       string // 算法id
    AlgorithmReturn   string // 算法返回类型
    AlgorithmCallback string // 算法回调地址
    Ptz               string // 云台参数
    Area              string // 作用区域
    CatchPhoto        int    // 抓图次数
    CatchTime         int    // 持续时间
}

// AddTask 向算法端添加一个子任务
func AddTask(task models.SonTask) (string, error) {
    apiUrl := config.AlgorithmPath + "/add_tasks"
    contentType := "application/json"

    var father models.RootTask
    db := utils.DBOpen()
    db.Where("task_id = ?", task.ParentId).Find(&father)

    alg := AlgTask{
        TaskId:            task.TaskId,
        StartTime:         father.StartTime,
        EndTime:           father.EndTime,
        AlgorithmId:       task.AlgorithmId,
        AlgorithmReturn:   father.AlgorithmReturn,
        AlgorithmCallback: father.AlgorithmCallback,
        Ptz:               task.Ptz,
        Area:              task.Area,
        CatchPhoto:        task.CatchPhoto,
        CatchTime:         task.CatchTime,
        CameraIp:          father.CameraIp,
    }
    data, _ := json.Marshal(&alg)

    resp, err := http.Post(apiUrl, contentType, strings.NewReader(string(data)))
    if err != nil {
        return "false", err
    }
    defer resp.Body.Close()
    b, err := ioutil.ReadAll(resp.Body)

    return string(b), err
}

// CheckRootUpdate 检查根任务中是否有需修改字段, 并进行修改
func CheckRootUpdate(TaskId string, update map[string]interface{}) (string, error) {
    // 检查是否有需修改字段
    var flag = 0
    for k, _ := range update {
        switch k {
        case "start_time":
            flag = 1
        case "end_time":
            flag = 1
        case "algorithm_return":
            flag = 1
        case "algorithm_callback":
            flag = 1
        }
    }
    // 有修改字段
    if flag == 1 {
        var task []models.SonTask
        db := utils.DBOpen()
        db.Where("parent_id = ?", TaskId).Find(&task)
        for _, k := range task {
            state, err := UpdateTask(k.TaskId, update)
            if err != nil || state == "false" {
                return state, err
            }
        }
    }
    return "true", nil
}

// UpdateTask 向算法端更新一个子任务
func UpdateTask(TaskId string, update map[string]interface{}) (string, error) {
    apiUrl := config.AlgorithmPath + "/update_tasks"
    contentType := "application/json"

    updateThing := MapToUpdate(TaskId,update)
    data, _ := json.Marshal(&updateThing)

    client := &http.Client{}
    resp, err := http.NewRequest(http.MethodPut, apiUrl, strings.NewReader(string(data)))
    if err != nil {
        return "false", err
    }
    resp.Header.Add("Content-Type", contentType)

    res, err := client.Do(resp)
    if err != nil {
        return "false", err
    }
    defer res.Body.Close()
    b, err := ioutil.ReadAll(res.Body)

    return string(b), err
}

// DeleteTask 向算法端删除一个子任务
func DeleteTask(TaskId string) (string, error) {
    apiUrl := config.AlgorithmPath + "/delete_tasks"
    contentType := "application/json"

    data, _ := json.Marshal(&TaskId)

    client := &http.Client{}
    resp, err := http.NewRequest(http.MethodDelete, apiUrl, strings.NewReader(string(data)))
    if err != nil {
        return "false", err
    }
    resp.Header.Add("Content-Type", contentType)

    res, err := client.Do(resp)
    if err != nil {
        return "false", err
    }
    defer res.Body.Close()
    b, err := ioutil.ReadAll(res.Body)

    return string(b), err
}

// MapToUpdate 数组映射
func MapToUpdate(id string, update map[string]interface{}) Update {
    updateThing := NewUpdate()
    updateThing.TaskId = id
    for i, k := range update {
        switch i {
        case "start_time":
            updateThing.StartTime = k.(string)
        case "end_time":
            updateThing.EndTime = k.(string)
        case "algorithm_id":
            updateThing.AlgorithmId = k.(string)
        case "algorithm_callback":
            updateThing.AlgorithmCallback = k.(string)
        case "algorithm_return":
            updateThing.AlgorithmReturn = k.(string)
        case "ptz":
            updateThing.Ptz = k.(string)
        case "area":
            updateThing.Area = k.(string)
        case "catch_photo":
            updateThing.CatchPhoto = k.(int)
        case "catch_time":
            updateThing.CatchTime = k.(int)
        }
    }
    return updateThing
}

// NewUpdate 新建一个更新
func NewUpdate() Update {
    return Update{
        TaskId:            "None",
        StartTime:         "None",
        EndTime:           "None",
        AlgorithmId:       "None",
        AlgorithmReturn:   "None",
        AlgorithmCallback: "None",
        Ptz:               "None",
        Area:              "None",
        CatchPhoto:        0,
        CatchTime:         0,
    }
}
