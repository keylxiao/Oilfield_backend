package models

type HostMonitor struct {
	AllDisk    uint64 // 磁盘总量
	UsedDisk   uint64 // 磁盘已使用量
	UsageCPU   uint64 // CPU使用率
	UsageGPU   uint64 // GPU使用率
}