package utils

import (
    "Oilfield_backend/config"
    "fmt"
    "github.com/spf13/viper"
    "os"
    "strconv"
)

// 配置项文件
type Options struct {
    Path string
    Name string
}

var DefaultOption = &Options{
    Path: "./log/",
    Name: "login",
}

// 文件处理操作

// New 创建配置文件
func NewOption(path, name string) error {
    // 创建目录机制
    CreateDic(path)
    // 创建文件机制
    path = GetAddress(path, name)
    if _, err := os.Stat(path); err == nil {
        viper.SetConfigFile(path)
        return nil
    }
    viper.SetConfigFile(path)
    // 新建表头
    viper.Set("name", name)
    viper.Set("count", 0)
    err := viper.WriteConfig()
    if err != nil {
        return err
    }
    return nil
}

// Read 读取文件
func Read(path, name string) ([]map[string]interface{}, error) {
    path = GetAddress(path, name)
    viper.SetConfigFile(path)
    err := viper.ReadInConfig()
    if err != nil {
        return nil, err
    }
    count := viper.GetInt("count")
    content := make([]map[string]interface{}, count)
    for i := 1; i <= count; i++ {
        content[i-1] = viper.GetStringMap(strconv.Itoa(i))
    }
    return content, nil
}

// Write 写入文件
func Write(path, name string, message interface{}) error {
    path = GetAddress(path, name)
    // 打开文件
    viper.SetConfigFile(path)
    err := viper.ReadInConfig()
    if err != nil {
        return err
    }
    // 更新消息
    count := viper.GetInt("count")
    viper.Set("count", count+1)
    viper.Set(strconv.Itoa(count+1), message)
    err = viper.WriteConfig()
    if err != nil {
        return err
    }
    return nil
}
// ReadWarning 读取配置
func ReadWarning() (map[string]interface{}, error) {
    path := GetAddress(config.WarningLogPath,config.WarningLogName)
    viper.SetConfigFile(path)
    err := viper.ReadInConfig()
    if err != nil {
        return nil, err
    }
    return viper.GetStringMap("setting"), err
}

// WriteWarning 写入配置
func WriteWarning(message map[string]interface{}) error {
    path := GetAddress(config.WarningLogPath,config.WarningLogName)
    // 打开文件
    viper.SetConfigFile(path)
    err := viper.ReadInConfig()
    if err != nil {
        return err
    }
    // 更新配置
    viper.Set("setting",message)
    err = viper.WriteConfig()
    if err != nil {
        return err
    }
    return nil
}
// GetAddress 获取使用的完整路径
func GetAddress(path, name string) string {
    return fmt.Sprintf(path + name + ".toml")
}

// CreateDic 会在目录不存在时自动创建相应的目录
func CreateDic(path string) {
    _, err := os.Stat(path)
    if err != nil {
        _ = os.Mkdir(path, 0777)
    }
}
