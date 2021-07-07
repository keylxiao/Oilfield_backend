package controllers

import (
    "Oilfield_backend/models"
    "github.com/kataras/iris/v12"
    "net/http"
)

// PostCreateEquipIpc 新增设备Ipc
func PostCreateEquipIpc(c iris.Context) {
    var result models.Ipc
    c.ReadJSON(&result)
    err := models.PostCreateEquip("ipc", result)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("ok")
    }
}

// PostCreateEquipNvr 新增设备Nvr
func PostCreateEquipNvr(c iris.Context) {
    var result models.Nvr
    c.ReadJSON(&result)
    err := models.PostCreateEquip("Nvr", result)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("ok")
    }
}

// GetAllEquip 查看所有设备
func GetAllEquip(c iris.Context) {
    equip := c.URLParam("type")
    result, err := models.GetAllEquip(equip)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// PostSomeEquip 条件查看设备
func PostSomeEquip(c iris.Context) {
    var condition models.Equip
    c.ReadJSON(&condition)
    result, err := models.PostSomeEquip(condition.EquipType, condition.Condition)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// PutUpdateEquip 更新设备信息
func PutUpdateEquip(c iris.Context) {
    var condition []models.UpdateEquip
    c.ReadJSON(&condition)
    for _, k := range condition {
       err := models.PutUpdateEquip(k.Type, k.Equip.EquipType, k.Equip.Condition)
       if err != nil {
           c.StatusCode(http.StatusInternalServerError)
           return
       }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("更新成功")
}

// DeleteEquip 删除设备
func DeleteEquip(c iris.Context) {
    var dele []models.DeleteSomeEquip
    c.ReadJSON(&dele)
    for _, k := range dele {
        err := models.DeleteEquip(k.EquipType, k.Number)
        if err != nil {
            c.StatusCode(http.StatusInternalServerError)
            return
        }
    }
    c.StatusCode(http.StatusOK)
    c.JSON("删除成功")
}
