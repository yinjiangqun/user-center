package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
    BACKEND_SERVER_PORT = 8081
)

// Provider 定义配置提供者接口
type Provider interface {
    GetString(key string) string
    GetInt(key string) int
    GetBool(key string) bool
    GetDuration(key string) time.Duration
    // 添加其他必要的方法...
}

type viperProvider struct {
    v *viper.Viper
}

var defaultConfig *viperProvider

func init() {
    v := viper.New()
    v.SetEnvPrefix("USER-CENTER")
    v.AutomaticEnv()
    
    // 设置默认值
    v.SetDefault("server.port", BACKEND_SERVER_PORT)
    
    // 加载配置文件
    v.SetConfigFile("config.json")
    if err := v.ReadInConfig(); err != nil {
        // 使用日志记录错误，而不是直接panic
        fmt.Printf("Error reading config file: %v\n", err)
    }
    
    defaultConfig = &viperProvider{v: v}
}

// Config 返回默认配置提供者
func GetConfig() Provider {
    return defaultConfig
}

// 实现Provider接口的方法
func (p *viperProvider) GetString(key string) string {
    return p.v.GetString(key)
}

func (p *viperProvider) GetInt(key string) int {
    return p.v.GetInt(key)
}

func (p *viperProvider) GetBool(key string) bool {
    return p.v.GetBool(key)
}

func (p *viperProvider) GetDuration(key string) time.Duration {
    return p.v.GetDuration(key)
}