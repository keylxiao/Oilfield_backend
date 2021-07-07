package controllers

import (
	"C"
	"Oilfield_backend/models"
	"github.com/kataras/iris/v12"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"gitlab.com/tingshuo/go-diskstate/diskstate"
	"net/http"
	"time"
)
// GetHostMonitor 获取主机监控信息
func GetHostMonitor(c iris.Context){
	state := diskstate.DiskUsage("/")
	percent, _:= cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	host := models.HostMonitor{
		AllDisk:  state.All/diskstate.MB,
		UsedDisk: state.Used/diskstate.MB,
		UsageCPU: uint64(percent[0]),
		UsageGPU: uint64(memInfo.UsedPercent),
	}
	c.StatusCode(http.StatusOK)
	c.JSON(host)
}