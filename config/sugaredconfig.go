package config

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config 将配置文件的参数解析,比如解析时间为 time.Ticker
type Config struct {
	*ConfigBase
	AuthExpireTime time.Duration
}

// 讀取 yml 檔案
func NewYmlConfig(filePath string) *Config {
	// 初始化配置文件
	pflag.StringP("config", "c", filePath, "config file")
	pflag.Parse()
	viper.SetConfigType("yaml")
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
	conf := viper.GetString("config")
	viper.SetConfigFile(conf)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("load config %s fail: %v", conf, err))
	}

	// 解析初始配置
	baseConf := &ConfigBase{}
	if err := viper.Unmarshal(baseConf); err != nil {
		if err != nil {
			panic(err)
		}
	}

	// AuthExpireTime 解析为 time.Duration
	authExpireTime, err := time.ParseDuration(baseConf.Auth.ExpireTime)
	if err != nil {
		panic(err)
	}

	// 构造 Config
	pConfig := &Config{
		ConfigBase:     baseConf,
		AuthExpireTime: authExpireTime,
	}

	return pConfig
}
