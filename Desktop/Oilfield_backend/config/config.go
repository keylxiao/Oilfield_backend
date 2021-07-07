package config

import "time"

// 基本参数配置
const (
    // 后端地址
    BackendAddress string = "localhost"
    // 端口
    Port string = ":8000"
    // 算法储存地址
    AlgorithmAddress string = "./scripts/"
    // 算法文件最大限制
    FileSizeLimit int64 = 1024
    // 登录日志地址
    LoginLogPath string = "./log/"
    // 登录日志名称
    LoginLogName string = "Login"
    // 告警日志地址
    WarningLogPath string = "./log/"
    // 告警日志名称
    WarningLogName string = "Warning"
    // 算法端地址
    AlgorithmPath string = "http://localhost:8001"
    // 访问摄像头超时时间
    GetEquipTime = 3 * time.Second
    // 循环检测设备地址的时间间隔
    CycleEquipTimeDuration = 168 * time.Hour
    // Rtmp地址首部
    RtmpPath string = "http://10.0.120.127:8002/flv?port=1985&app=live&stream="
    // Rtmp视频流停留时间
    RtmpTime = 300 * time.Second
)

// MySQL参数配置
const (
    // 数据库类型
    DB string = "mysql"
    // 数据库用户名称
    DBUserName string = "oilfield"
    // 数据库密码
    DBPassword string = "1984051718"
    // 数据库名称
    DBName string = "oilfield"
    // 数据库占用端口
    DBRemote string = "127.0.0.1:3306"
    // 数据库编码
    DBCharset string = "utf8"
    // 解析时间
    DBParseTime string = "True"
    // DBLoc
    DBLoc string = "Local"
)

// Redis参数配置
const (
    // 数据库占用端口
    RedisRemote string = "localhost:6379"
    // 数据库密码
    RedisPassword string = "1984051718"
    // 数据库名称
    RedisName int = 0
)
