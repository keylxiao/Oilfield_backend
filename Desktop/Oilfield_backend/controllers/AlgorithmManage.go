package controllers

import (
    "Oilfield_backend/config"
    "Oilfield_backend/models"
    "Oilfield_backend/utils"
    "github.com/kataras/iris/v12"
    "io"
    "net/http"
    "os"
    "time"
)

// PostCreateAlgorithm 创建一条算法
func PostCreateAlgorithm(c iris.Context) {
    var alg models.Algorithm
    c.ReadJSON(&alg)
    alg.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    alg.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
    alg.AlgorithmId = utils.GetUUID()
    alg.Address = config.AlgorithmAddress
    err := models.PostCreateAlgorithm(alg)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(alg.AlgorithmId)
    }
}

// PostCreateAlgorithmFile 上传算法文件
func PostCreateAlgorithmFile(c iris.Context) {
    file, _, err := c.FormFile("file")
    if err != nil {
        c.StatusCode(http.StatusBadRequest)
        c.JSON("上传错误")
        return
    }
    dest, err := os.Create(config.AlgorithmAddress + c.FormValue("name"))
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        return
    }
    _, err = io.Copy(dest, file)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    }else{
        c.StatusCode(http.StatusOK)
        c.JSON("上传成功")
    }
}

// GetAllAlgorithm 查看所有算法
func GetAllAlgorithm(c iris.Context){
    alg,err := models.GetAllAlgorithm()
    if err != nil{
        c.StatusCode(http.StatusInternalServerError)
        return
    }else{
        c.StatusCode(http.StatusOK)
        c.JSON(alg)
    }
}

// PutUpdateAlgorithm 更新算法信息
func PutUpdateAlgorithm(c iris.Context){
    var update []models.UpdateAlgorithm
    c.ReadJSON(&update)
    for _,result := range update{
        err := models.PutUpdateAlgorithm(result.AlgorithmId,result.Update)
        if err != nil {
            c.StatusCode(http.StatusInternalServerError)
            return
        }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("更新成功")
}

// DeleteAlgorithm 删除算法
func DeleteAlgorithm(c iris.Context){
    var result []models.DeleteSomeAlgorithm
    c.ReadJSON(&result)
    for _,id := range result{
        err := models.DeleteAlgorithm(id.AlgorithmId)
        if err != nil {
            c.StatusCode(http.StatusInternalServerError)
            return
        }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("删除成功")
}