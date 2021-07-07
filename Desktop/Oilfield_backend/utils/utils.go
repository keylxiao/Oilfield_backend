package utils

import (
    "Oilfield_backend/config"
    "encoding/base64"
    "fmt"
    "github.com/go-redis/redis"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gofrs/uuid"
    "github.com/jinzhu/gorm"
    "net/http"
    "os"
    "time"
)

// BackDBInstance 数据库配置
func BackDBInstance() (string, interface{}) {
    command := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%s&loc=%s", config.DBUserName, config.DBPassword, config.DBRemote, config.DBName, config.DBCharset, config.DBParseTime, config.DBLoc)
    return config.DB, command
}

// DBOpen 打开数据库
func DBOpen() *gorm.DB {
    db, _ := gorm.Open(BackDBInstance())
    return db
}

// RedisOpen 打开Redis
func RedisOpen() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     config.RedisRemote,
        Password: config.RedisPassword,
        DB:       config.RedisName,
    })
    return rdb
}

// PasswordSha1 密码sha1加密
//func PasswordSha1(s string) string {
//    o := sha1.New()
//    o.Write([]byte(s))
//    return hex.EncodeToString(o.Sum(nil))
//}

// 生成随机数 32位
func GetUUID() string {
    uid := uuid.NewV5(uuid.Must(uuid.NewV4()), "oilfield").String()
    uid = uid[:8] + uid[9:13] + uid[14:18] + uid[19:23] + uid[24:]
    return uid
}

// WriteLoginLog 记录登录日志
func WriteLoginLog(name string) {
    message := map[string]interface{}{
        "Time":    time.Now().Unix(),
        "Account": name,
    }
    Write(config.LoginLogPath, config.LoginLogName, message)
}

// ReadLoginLog 读取登录日志
func
ReadLoginLog() ([]map[string]interface{}, error) {
    return Read(config.LoginLogPath, config.LoginLogName)
}

// WriteWarningLog 记录告警日志
//func WriteWarningLog(id,number,reason string,degree int) error {
//    message := map[string]interface{}{
//        "WarningId":     id,
//        "CameraNumber":  number,
//        "WarningReason": reason,
//        "WarningTime":   time.Now().Format("2006-01-02 15:04:05"),
//        "WarningDegree": degree,
//	}
//    err := Write(config.WarningLogPath, config.WarningLogName, message)
//    return err
//}
// ReadWarningLog 读取登录日志
//func ReadWarningLog() ([]map[string]interface{}, error) {
//	return Read(config.WarningLogPath, config.WarningLogName)
//}

// PhotoToBase64 图片转码base64
func PhotoToBase64(path, name string) string {
    setting, _ := ReadWarning()
    ff, _ := os.Open(path + name + setting["warningphototype"].(string))
    defer ff.Close()
    fi, _ := ff.Stat()
    sourceBuffer := make([]byte, fi.Size())
    n, _ := ff.Read(sourceBuffer)
    sourceString := base64.StdEncoding.EncodeToString(sourceBuffer[:n])
    return sourceString
}
func GetIP(r *http.Request) string {
    forwarded := r.Header.Get("X-FORWARDED-FOR")
    if forwarded != "" {
        return forwarded
    }
    return r.RemoteAddr
}