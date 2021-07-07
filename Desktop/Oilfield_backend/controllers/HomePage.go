package controllers

import (
    "Oilfield_backend/models"
    "github.com/kataras/iris/v12"
    "net/http"
)

// GetOverview 首页总览信息
func GetOverview(c iris.Context) {
    result, err := models.GetOverview()
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}
